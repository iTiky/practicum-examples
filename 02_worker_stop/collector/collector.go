package collector

import (
	"fmt"
	"time"

	"github.com/itiky/practicum-examples/02_worker_stop/market"
)

const (
	DefaultWindowSize = 5
)

type (
	// MarketEventsCollector collects market events and builds price statistics per asset pair.
	MarketEventsCollector struct {
		windowSize int
		state      marketsState
		eventsCh   chan marketEvent
		statsReqCh chan statisticsRequest
		stopCh     chan struct{}
	}

	// MarketEventsCollectorOption defines option used by the NewMarketEventsCollector constructor.
	MarketEventsCollectorOption func(*MarketEventsCollector)
)

// WithMovingAvgWindowSize sets moving average window size for price and receivedAt values.
func WithMovingAvgWindowSize(size int) MarketEventsCollectorOption {
	return func(c *MarketEventsCollector) {
		c.windowSize = size
	}
}

// NewMarketEventsCollector creates a new MarketEventsCollector instance.
func NewMarketEventsCollector(opts ...MarketEventsCollectorOption) *MarketEventsCollector {
	c := &MarketEventsCollector{
		windowSize: DefaultWindowSize,
		state:      newMarketsState(),
		eventsCh:   make(chan marketEvent, 10),
		statsReqCh: make(chan statisticsRequest),
		stopCh:     make(chan struct{}),
	}

	for _, opt := range opts {
		opt(c)
	}

	go c.worker()

	return c
}

// MarketEventHandler returns market.StockEventHandler callback.
func (c *MarketEventsCollector) MarketEventHandler() market.StockEventHandler {
	return func(baseAsset, quoteAsset string, askPrice, bidPrice float64, receivedAt time.Time) {
		c.eventsCh <- marketEvent{
			baseAsset:  baseAsset,
			quoteAsset: quoteAsset,
			askPrice:   askPrice,
			bidPrice:   bidPrice,
			receivedAt: receivedAt,
		}
	}
}

// BuildStatistics requests statistics build and returns the result.
func (c *MarketEventsCollector) BuildStatistics() Statistics {
	req := newStatisticsRequest()
	c.statsReqCh <- req

	return <-req.retCh
}

// Close stops the collector.
func (c *MarketEventsCollector) Close() error {
	if c.stopCh == nil {
		return fmt.Errorf("already stopped")
	}
	close(c.stopCh)

	return nil
}

// worker is a main collector loop.
func (c *MarketEventsCollector) worker() {
	for {
		select {
		case <-c.stopCh:
			break
		case event := <-c.eventsCh:
			c.handleEvent(event)
		case req := <-c.statsReqCh:
			c.handleStatisticsRequest(req)
		}
	}
}

// handleEvent handles market event.
func (c *MarketEventsCollector) handleEvent(event marketEvent) {
	// Build stats key
	pairID := fmt.Sprintf("%s-%s", event.baseAsset, event.quoteAsset)

	// Create / get market stats
	marketState := c.state[pairID]
	if marketState == nil {
		marketState = newMarketState(c.windowSize)
		c.state[pairID] = marketState
	}

	// Update price stats
	marketState.askPriceAvg.Add(event.askPrice)
	marketState.bidPriceAvg.Add(event.bidPrice)

	// Update receivedAt diff stats
	var updatedAtDiff time.Duration
	if marketState.receivedAtPrev.IsZero() {
		updatedAtDiff = 0
	} else {
		updatedAtDiff = event.receivedAt.Sub(marketState.receivedAtPrev)
	}
	marketState.updateDurAvg.Add(float64(updatedAtDiff))
	marketState.receivedAtPrev = event.receivedAt
}

// handleStatisticsRequest handles statistics request event.
func (c *MarketEventsCollector) handleStatisticsRequest(req statisticsRequest) {
	req.retCh <- newStatistics(c.state)
}

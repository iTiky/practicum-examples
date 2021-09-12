package collector

import (
	"time"

	movingaverage "github.com/RobinUS2/golang-moving-average"
)

// marketEvent defines market price change event meta.
type marketEvent struct {
	baseAsset  string
	quoteAsset string
	askPrice   float64
	bidPrice   float64
	receivedAt time.Time
}

type statisticsRequest struct {
	retCh chan Statistics
}

// newStatisticsRequest creates a new statisticsRequest.
func newStatisticsRequest() statisticsRequest {
	return statisticsRequest{
		retCh: make(chan Statistics),
	}
}

type (
	// marketsState defines per market statistics state (key: base-quote asset pair).
	marketsState map[string]*marketState

	// marketState defines average price change and price update timestamp diff state for a market asset pair.
	marketState struct {
		askPriceAvg    *movingaverage.MovingAverage
		bidPriceAvg    *movingaverage.MovingAverage
		updateDurAvg   *movingaverage.MovingAverage
		receivedAtPrev time.Time
	}
)

// newMarketsState creates a new marketsState object.
func newMarketsState() marketsState {
	return make(marketsState)
}

// newMarketState creates a new marketState object.
func newMarketState(windowSize int) *marketState {
	return &marketState{
		askPriceAvg:  movingaverage.New(windowSize),
		bidPriceAvg:  movingaverage.New(windowSize),
		updateDurAvg: movingaverage.New(windowSize),
	}
}

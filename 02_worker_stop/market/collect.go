package market

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/rs/zerolog/log"
)

// StockEventHandler defines CollectBinanceSpotPrices callback.
type StockEventHandler func(baseAsset, quoteAsset string, askPrice, bidPrice float64, receivedAt time.Time)

// CollectBinanceSpotPrices collects stock prices for the defines market symbol.
func CollectBinanceSpotPrices(ctx context.Context, baseAsset, quoteAsset string, handler StockEventHandler) error {
	if handler == nil {
		return fmt.Errorf("handler: nil")
	}

	symbol := strings.ToTitle(baseAsset) + strings.ToTitle(quoteAsset)

	logger := log.With().
		Str("service", "stock-collector").
		Str("symbol", symbol).
		Logger()
	logger.Info().Msg("Started")

	// Define web socket callbacks
	wsDepthHandler := func(event *binance.WsDepthEvent) {
		if len(event.Asks) == 0 || len(event.Bids) == 0 {
			logger.Warn().Msg("Depth event has not bid / ask values (skip)")
			return
		}

		askPrice, _, err := event.Asks[0].Parse()
		if err != nil {
			logger.Error().Err(err).Msg("Parsing ASK price")
			return
		}

		bidPrice, _, err := event.Bids[0].Parse()
		if err != nil {
			logger.Error().Err(err).Msg("Parsing BID price")
			return
		}

		receivedAt := time.Unix(0, event.Time*1e6)

		go handler(baseAsset, quoteAsset, askPrice, bidPrice, receivedAt)
		logger.Debug().Msgf("Event collected: %s", receivedAt)
	}

	wsErrHandler := func(err error) {
		logger.Error().Err(err).Msg("WebSocket connection error")
	}

	// Create web socket connection
	_, wsStopCh, err := binance.WsDepthServe(symbol, wsDepthHandler, wsErrHandler)
	if err != nil {
		return fmt.Errorf("subscribing to Binance web socket: %w", err)
	}

	// Worker
	go func() {
		// Wait for cancel
		<-ctx.Done()

		// Close web socket connection
		wsStopCh <- struct{}{}

		logger.Info().Err(ctx.Err()).Msg("Stopped")
	}()

	return nil
}

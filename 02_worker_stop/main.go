package main

import (
	"context"
	"os"
	"time"

	"github.com/itiky/practicum-examples/02_worker_stop/collector"
	"github.com/itiky/practicum-examples/02_worker_stop/market"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	statsCollector := collector.NewMarketEventsCollector()
	defer statsCollector.Close()

	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	if err := market.CollectBinanceSpotPrices(ctx, "btc", "usdt", statsCollector.MarketEventHandler()); err != nil {
		log.Fatal().Err(err).Msg("Starting the btc-usdt collector")
		os.Exit(1)
	}
	if err := market.CollectBinanceSpotPrices(ctx, "eth", "usdt", statsCollector.MarketEventHandler()); err != nil {
		log.Fatal().Err(err).Msg("Starting the eth-usdt collector")
		os.Exit(1)
	}

	for i := 0; i < 3; i++ {
		time.Sleep(3 * time.Second)
		result := statsCollector.BuildStatistics()
		log.Info().Msgf("Current results: %s", result.String())
	}
}

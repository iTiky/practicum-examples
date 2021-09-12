package collector

import (
	"fmt"
	"strings"
	"time"
)

type (
	// Statistics keeps market price statistics.
	Statistics []StatisticsItem

	StatisticsItem struct {
		AssetPair         string        `json:"asset_pair"`
		AvgAskPrice       float64       `json:"avg_ask_price"`
		AvgBidPrice       float64       `json:"avg_bid_price"`
		AvgPriceUpdateDur time.Duration `json:"avg_price_update_dur,omitempty"`
	}
)

// newStatistics converts marketsState to Statistics.
func newStatistics(state marketsState) Statistics {
	stats := make(Statistics, 0, len(state))
	for pairID, marketState := range state {
		stats = append(
			stats,
			StatisticsItem{
				AssetPair:         pairID,
				AvgAskPrice:       marketState.askPriceAvg.Avg(),
				AvgBidPrice:       marketState.bidPriceAvg.Avg(),
				AvgPriceUpdateDur: time.Duration(marketState.updateDurAvg.Avg()),
			},
		)
	}

	return stats
}

// String implements the fmt.Stringer interface.
func (s Statistics) String() string {
	str := strings.Builder{}

	for _, item := range s {
		str.WriteString("\n")
		str.WriteString(fmt.Sprintf("- %s:\n", item.AssetPair))
		str.WriteString(fmt.Sprintf("  Avg ASK price:        %f\n", item.AvgAskPrice))
		str.WriteString(fmt.Sprintf("  Avg BID price:        %f\n", item.AvgBidPrice))
		str.WriteString(fmt.Sprintf("  Avg price update dur: %v", item.AvgPriceUpdateDur))
	}

	return str.String()
}

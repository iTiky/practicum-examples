//go:generate mockgen -source=interface.go -destination=./mock/service.go -package=orderservicemock
package order

import (
	"context"

	"github.com/itiky/practicum-examples/01_log_and_trace_context/model"
)

// Processor defines model.Order operations.
type Processor interface {
	// ProcessOrder handles order product items enriching them with library service.
	ProcessOrder(ctx context.Context, order model.Order) error
}

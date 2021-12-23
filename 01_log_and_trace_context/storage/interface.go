//go:generate mockgen -source=interface.go -destination=./mock/storage.go -package=storagemock
package storage

import (
	"context"

	"github.com/itiky/practicum-examples/01_log_and_trace_context/model"
)

// ProductStorer defines test's product item operations.
type ProductStorer interface {
	// SaveOrderProductItem stores test's product item.
	SaveOrderProductItem(ctx context.Context, orderID string, item model.Product) error
}

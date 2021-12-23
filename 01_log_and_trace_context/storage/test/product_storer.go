package test

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/itiky/practicum-examples/01_log_and_trace_context/model"
	"github.com/itiky/practicum-examples/01_log_and_trace_context/pkg/tracing"
	"github.com/itiky/practicum-examples/01_log_and_trace_context/storage"
)

var _ storage.ProductStorer = (*Storage)(nil)

// Storage keeps storage repository dependencies.
type Storage struct{}

// NewStorage creates a new Storage instance.
func NewStorage() *Storage {
	return &Storage{}
}

// SaveOrderProductItem implements the ProductStorer interface.
func (s Storage) SaveOrderProductItem(ctx context.Context, orderID string, item model.Product) (retErr error) {
	ctx, span := tracing.StartSpanFromCtx(ctx, "Saving test's product item")
	defer func() { tracing.FinishSpan(span, retErr) }()

	time.Sleep(time.Duration(rand.Int31n(100)) * time.Millisecond)

	if item.Quantity < 0 {
		return fmt.Errorf("quantity: must be GT 0")
	}

	return nil
}

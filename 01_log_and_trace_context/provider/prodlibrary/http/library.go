package http

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/itiky/practicum-examples/01_log_and_trace_context/pkg/tracing"
	"github.com/itiky/practicum-examples/01_log_and_trace_context/provider/prodlibrary"
)

var _ prodlibrary.ProductNameProvider = (*Provider)(nil)

// Provider keeps products library service configuration.
type Provider struct{}

// NewProvider returns a new Provider instance.
func NewProvider() *Provider {
	return &Provider{}
}

// GetProductName implements the ProductNameProvider interface.
func (p Provider) GetProductName(ctx context.Context, barcode string) (retName string, retErr error) {
	ctx, span := tracing.StartSpanFromCtx(ctx, "Getting product name from library")
	defer func() { tracing.FinishSpan(span, retErr) }()

	time.Sleep(time.Duration(rand.Int31n(100)) * time.Millisecond)

	if barcode == "" {
		return "", fmt.Errorf("barcode: empty")
	}

	return fmt.Sprintf("Product_%s", barcode), nil
}

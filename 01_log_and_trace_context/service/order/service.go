package order

import (
	"context"
	"fmt"

	"github.com/itiky/practicum-examples/01_log_and_trace_context/pkg/logging"
	"github.com/itiky/practicum-examples/01_log_and_trace_context/provider/product_library"
	"github.com/itiky/practicum-examples/01_log_and_trace_context/storage/order"
	"github.com/rs/zerolog"
)

var _ Processor = (*Service)(nil)

type (
	// Service keeps service dependencies.
	Service struct {
		productLibrarySvc product_library.ProductNameProvider
		productSt         order.ProductStorer
	}

	// Option defines functional argument for Service constructor.
	Option func(*Service)
)

// WithProductNameProvider sets product_library.ProductNameProvider for the Service service.
func WithProductNameProvider(provider product_library.ProductNameProvider) Option {
	return func(p *Service) {
		p.productLibrarySvc = provider
	}
}

// WithProductStorage sets order.ProductStorer for the Service service.
func WithProductStorage(storage order.ProductStorer) Option {
	return func(p *Service) {
		p.productSt = storage
	}
}

// NewProcessor creates a new Service instance checking dependencies and configuration.
func NewProcessor(opts ...Option) (*Service, error) {
	svc := &Service{}
	for _, opt := range opts {
		opt(svc)
	}

	if svc.productLibrarySvc == nil {
		return nil, fmt.Errorf("productLibrarySvc: nil")
	}
	if svc.productSt == nil {
		return nil, fmt.Errorf("productSt: nil")
	}

	return svc, nil
}

// Log returns logger with service field set.
func (svc Service) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "order processor").Logger()

	return &logger
}

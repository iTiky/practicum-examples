package v1

import (
	"context"
	"fmt"

	"github.com/itiky/practicum-examples/01_log_and_trace_context/pkg/logging"
	"github.com/itiky/practicum-examples/01_log_and_trace_context/provider/prodlibrary"
	"github.com/itiky/practicum-examples/01_log_and_trace_context/service/order"
	"github.com/itiky/practicum-examples/01_log_and_trace_context/storage"
	"github.com/rs/zerolog"
)

var _ order.Processor = (*Service)(nil)

type (
	// Service keeps service dependencies.
	Service struct {
		productLibrarySvc prodlibrary.ProductNameProvider
		productSt         storage.ProductStorer
	}

	// Option defines functional argument for Service constructor.
	Option func(*Service) error
)

// WithProductNameProvider sets prodlibrary.ProductNameProvider for the Service service.
func WithProductNameProvider(provider prodlibrary.ProductNameProvider) Option {
	return func(p *Service) error {
		if provider == nil {
			return fmt.Errorf("productNameProvider: nil")
		}
		p.productLibrarySvc = provider

		return nil
	}
}

// WithProductStorage sets test.ProductStorer for the Service service.
func WithProductStorage(storage storage.ProductStorer) Option {
	return func(p *Service) error {
		if storage == nil {
			return fmt.Errorf("productStorer: nil")
		}
		p.productSt = storage

		return nil
	}
}

// NewProcessor creates a new Service instance checking dependencies and configuration.
func NewProcessor(opts ...Option) (*Service, error) {
	svc := &Service{}
	for _, opt := range opts {
		if err := opt(svc); err != nil {
			return nil, fmt.Errorf("initialising dependencies: %w", err)
		}
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
	logger = logger.With().Str(logging.ServiceKey, "Order processor").Logger()

	return &logger
}

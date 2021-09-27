package order

import (
	"fmt"

	"github.com/itiky/practicum-examples/04_pgsql/pkg/logging"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	serviceName = "order-processor"
)

type (
	Processor struct {
		orderStorage StorageExpected
	}

	ProcessorOption func(svc *Processor)
)

// WithOrderStorage sets StorageExpected.
func WithOrderStorage(st StorageExpected) ProcessorOption {
	return func(svc *Processor) {
		svc.orderStorage = st
	}
}

// New creates a new Processor service.
func New(opts ...ProcessorOption) (*Processor, error) {
	svc := &Processor{}
	for _, opt := range opts {
		opt(svc)
	}

	if svc.orderStorage == nil {
		return nil, fmt.Errorf("orderStorage: nil")
	}

	return svc, nil
}

// Close closes all service dependencies.
func (svc *Processor) Close() error {
	if svc.orderStorage == nil {
		return nil
	}

	if err := svc.orderStorage.Close(); err != nil {
		return fmt.Errorf("closing orderStorage: %w", err)
	}

	return nil
}

// Logger returns logger with service context.
func (svc *Processor) Logger() zerolog.Logger {
	logCtx := log.With().Str(logging.ServiceKey, serviceName)

	return logCtx.Logger()
}

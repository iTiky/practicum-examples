package user

import (
	"fmt"

	"github.com/itiky/practicum-examples/04_pgsql/pkg/logging"
	"github.com/itiky/practicum-examples/04_pgsql/storage"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	serviceName = "user-processor"
)

type (
	Processor struct {
		config             Config
		allowedRegionCodes map[string]struct{}

		userStorage storage.UserWriter
	}

	ProcessorOption func(svc *Processor)
)

// WithUserWriter sets storage.UserWriter.
func WithUserWriter(st storage.UserWriter) ProcessorOption {
	return func(svc *Processor) {
		svc.userStorage = st
	}
}

// WithConfig sets Config.
func WithConfig(config Config) ProcessorOption {
	return func(svc *Processor) {
		svc.config = config
	}
}

// New creates a new Processor service.
func New(opts ...ProcessorOption) (*Processor, error) {
	svc := &Processor{
		allowedRegionCodes: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(svc)
	}

	if err := svc.config.Validate(); err != nil {
		return nil, fmt.Errorf("config validation: %w", err)
	}

	if svc.userStorage == nil {
		return nil, fmt.Errorf("userStorage: nil")
	}

	for _, regionCode := range svc.config.AllowedRegionCodes {
		svc.allowedRegionCodes[regionCode] = struct{}{}
	}

	return svc, nil
}

// Close closes all service dependencies.
func (svc *Processor) Close() error {
	if svc.userStorage == nil {
		return nil
	}

	if err := svc.userStorage.Close(); err != nil {
		return fmt.Errorf("closing userStorage: %w", err)
	}

	return nil
}

// Logger returns logger with service context.
func (svc *Processor) Logger() zerolog.Logger {
	logCtx := log.With().Str(logging.ServiceKey, serviceName)

	return logCtx.Logger()
}

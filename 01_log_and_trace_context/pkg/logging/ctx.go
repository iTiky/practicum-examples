package logging

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/itiky/practicum-examples/01_log_and_trace_context/pkg"

	"github.com/rs/zerolog"
)

const (
	// contextKeyLogger defines the key for the logger to be stored within request context.
	contextKeyLogger = pkg.ContextKey("Logger")

	// contextKeyCorrelationID defines the key for the correlation ID to be stored within request context.
	contextKeyCorrelationID = pkg.ContextKey("CorrelationID")
)

// GetCtxLogger returns a logger stored within the context.
// Creates a new eventservice with correlation ID field and stores it in the context.
func GetCtxLogger(ctx context.Context) (context.Context, zerolog.Logger) {
	if ctx == nil {
		ctx = context.Background()
	}

	if ctxValue := ctx.Value(contextKeyLogger); ctxValue != nil {
		if ctxLogger, ok := ctxValue.(zerolog.Logger); ok {
			return ctx, ctxLogger
		}
	}

	correlationID, _ := uuid.NewUUID()
	logger := NewLogger().With().Str(CorrelationIDKey, correlationID.String()).Logger()

	ctx = context.WithValue(ctx, contextKeyCorrelationID, correlationID.String())

	return SetCtxLogger(ctx, logger), logger
}

// SetCtxLogger adds the logger to the context overwriting and existing eventservice.
func SetCtxLogger(ctx context.Context, logger zerolog.Logger) context.Context {
	return context.WithValue(ctx, contextKeyLogger, logger)
}

// GetCorrelationID returns the correlation ID contained within the context.
func GetCorrelationID(ctx context.Context) (string, error) {
	id, ok := ctx.Value(contextKeyCorrelationID).(string)
	if !ok {
		return "", fmt.Errorf("correlation ID: not found")
	}

	return id, nil
}

// SetCorrelationID sets the correlation ID field for a logger stored within the context (overwrites an existing eventservice).
func SetCorrelationID(ctx context.Context, id string) (context.Context, zerolog.Logger) {
	ctx, logger := GetCtxLogger(ctx)
	logger = logger.With().Str(CorrelationIDKey, id).Logger()

	ctx = context.WithValue(ctx, contextKeyCorrelationID, id)

	return SetCtxLogger(ctx, logger), logger
}

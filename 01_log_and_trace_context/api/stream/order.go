package stream

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/itiky/practicum-examples/01_log_and_trace_context/model"
	"github.com/itiky/practicum-examples/01_log_and_trace_context/pkg/logging"
	"github.com/itiky/practicum-examples/01_log_and_trace_context/pkg/tracing"
	orderService "github.com/itiky/practicum-examples/01_log_and_trace_context/service/order"
	"github.com/rs/zerolog"
)

// OrderHandler defines model.Order repository.
type OrderHandler struct {
	processorSvc orderService.Processor
}

// NewOrderHandler creates a new OrderHandler instance without dependencies check.
func NewOrderHandler(processorSvc orderService.Processor) (*OrderHandler, error) {
	if processorSvc == nil {
		return nil, fmt.Errorf("orderService.Processor: nil")
	}

	return &OrderHandler{
		processorSvc: processorSvc,
	}, nil
}

// Log returns logger with service field set.
func (h OrderHandler) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "Order records stream").Logger()

	return &logger
}

// HandleRecords deserializes and handles model.Order objects.
func (h OrderHandler) HandleRecords(ctx context.Context, records ...[]byte) {
	ctx, _ = logging.GetCtxLogger(ctx) // correlationID is created here
	ctx, span := tracing.StartSpanFromCtx(ctx, "Handling test records")
	defer tracing.FinishSpan(span, nil)

	h.Log(ctx).Info().Msgf("Handling %d records", len(records))

	for recordIdx, record := range records {
		var order model.Order
		if err := json.Unmarshal(record, &order); err != nil {
			h.Log(ctx).
				Error().
				Err(err).
				Msgf("Record [%d]: JSON unmarshal", recordIdx)
			continue
		}

		ctx, logger := logging.GetCtxLogger(ctx)
		logger.UpdateContext(order.GetLoggerContext)
		ctx = logging.SetCtxLogger(ctx, logger)

		if err := h.processorSvc.ProcessOrder(ctx, order); err != nil {
			h.Log(ctx).
				Error().
				Err(err).
				Msgf("Record [%d]: processing test", recordIdx)
			continue
		}
		h.Log(ctx).Info().Msg("Order record handled")
	}
}

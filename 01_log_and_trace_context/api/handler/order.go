package handler

import (
	"context"
	"encoding/json"

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
func NewOrderHandler(processorSvc orderService.Processor) *OrderHandler {
	return &OrderHandler{
		processorSvc: processorSvc,
	}
}

// Log returns logger with service field set.
func (h OrderHandler) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "order records handler").Logger()

	return &logger
}

// HandleOrderRecords deserializes and handles model.Order objects.
func (h OrderHandler) HandleOrderRecords(records ...[]byte) {
	ctx, _ := logging.GetCtxLogger(context.Background()) // correlationID is created here
	ctx, span := tracing.StartSpanFromCtx(ctx, "Handling order records")
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
				Msgf("Record [%d]: processing order", recordIdx)
			continue
		}
		h.Log(ctx).Info().Msg("Order record handled")
	}
}

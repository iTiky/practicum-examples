package order

import (
	"context"
	"fmt"

	"github.com/itiky/practicum-examples/01_log_and_trace_context/model"
	"github.com/itiky/practicum-examples/01_log_and_trace_context/pkg/logging"
	"github.com/itiky/practicum-examples/01_log_and_trace_context/pkg/tracing"
)

// ProcessOrder implements the Processor interface.
func (svc Service) ProcessOrder(ctx context.Context, order model.Order) (retErr error) {
	ctx, span := tracing.StartSpanFromCtx(ctx, "Processing order")
	defer func() { tracing.FinishSpan(span, retErr) }()

	for itemIdx, item := range order.Items {
		ctx, logger := logging.GetCtxLogger(ctx)
		logger.UpdateContext(item.GetLoggerContext)
		ctx = logging.SetCtxLogger(ctx, logger)

		if err := svc.processItem(ctx, order.ID, item); err != nil {
			svc.Log(ctx).
				Warn().
				Err(err).
				Msgf("Product item [%d]: enriching and storing", itemIdx)
			return fmt.Errorf("processing item [%d]: %w", itemIdx, err)
		}
	}

	return
}

func (svc Service) processItem(ctx context.Context, orderID string, item model.Product) error {
	itemName, err := svc.productLibrarySvc.GetProductName(ctx, item.Barcode)
	if err != nil {
		return fmt.Errorf("getting product name from the library service: %w", err)
	}
	item.Name = itemName

	if err := svc.productSt.SaveOrderProductItem(ctx, orderID, item); err != nil {
		return fmt.Errorf("saving product item: %w", err)
	}

	svc.Log(ctx).
		Info().
		Str("product_name", item.Name).
		Int("product_quantity", item.Quantity).
		Msg("Product item saved")

	return nil
}

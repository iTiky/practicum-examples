package order

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/itiky/practicum-examples/04_pgsql/model"
	"github.com/itiky/practicum-examples/04_pgsql/pkg"
	"github.com/itiky/practicum-examples/04_pgsql/pkg/input"
)

func (svc *Processor) CreateOrder(ctx context.Context, userID uuid.UUID, status model.OrderStatus) (model.Order, error) {
	logger := svc.Logger()

	// Input checks
	if userID == uuid.Nil {
		return model.Order{}, fmt.Errorf("%w: userID: nil", pkg.ErrInvalidInput)
	}

	if err := status.Validate(); err != nil {
		return model.Order{}, fmt.Errorf("%w: status: %v", pkg.ErrInvalidInput, err)
	}

	// Build input
	input := model.Order{
		UserID: userID,
		Status: status,
	}

	logger.UpdateContext(input.GetLoggerContext)
	logger.Info().Msg("Creating order")

	// Create
	order, err := svc.orderStorage.CreateOrder(ctx, input)
	if err != nil {
		logger.Warn().Err(err).Msg("Creating order")
		return model.Order{}, fmt.Errorf("creating order: %w", err)
	}

	logger.UpdateContext(order.GetLoggerContext)
	logger.Info().Msg("Order created")

	return order, nil
}

func (svc *Processor) GetOrdersForUser(ctx context.Context,
	userID uuid.UUID,
	timeRangeStart, timeRangeEnd *time.Time,
	pageParams input.PageParams,
) ([]model.Order, error) {

	// Input checks
	if userID == uuid.Nil {
		return nil, fmt.Errorf("%w: userID is empty", pkg.ErrInvalidInput)
	}

	if timeRangeStart != nil && timeRangeEnd != nil {
		if timeRangeStart.After(*timeRangeEnd) {
			return nil, fmt.Errorf("%w: range start must be LT end", pkg.ErrInvalidInput)
		}
	}

	if err := pageParams.Validate(); err != nil {
		return nil, fmt.Errorf("%w: pagination params validation: %v", pkg.ErrInvalidInput, err)
	}

	// Request
	return svc.orderStorage.GetOrdersForUser(ctx, userID, timeRangeStart, timeRangeEnd, pageParams)
}

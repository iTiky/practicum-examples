package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/itiky/practicum-examples/04_pgsql/model"
	"github.com/itiky/practicum-examples/04_pgsql/pkg"
	"github.com/itiky/practicum-examples/04_pgsql/pkg/input"
	"github.com/itiky/practicum-examples/04_pgsql/storage/psql/schema"
	"github.com/uptrace/bun/driver/pgdriver"
)

// CreateOrder implements the storage.OrderWriter interface.
func (st Storage) CreateOrder(ctx context.Context, order model.Order) (model.Order, error) {
	dbObj := schema.NewOrderFromCanonical(order)
	dbObj.ID, dbObj.CreatedAt, dbObj.DeletedAt = uuid.Nil, time.Time{}, time.Time{}

	_, err := st.db.NewInsert().
		Model(&dbObj).
		Returning("*").
		Exec(ctx)
	if err != nil {
		pgErr := &pgdriver.Error{}
		if errors.As(err, pgErr) {
			if pgErr.IntegrityViolation() {
				return model.Order{}, fmt.Errorf("%w: user_id not found", pkg.ErrNotExists)
			}
		}
		return model.Order{}, err
	}

	obj, err := dbObj.ToCanonical()
	if err != nil {
		return model.Order{}, fmt.Errorf("conveting to canonical model: %w", err)
	}

	return obj, nil
}

// GetOrderByID implements the storage.OrderReader interface.
func (st Storage) GetOrderByID(ctx context.Context, id uuid.UUID) (*model.Order, error) {
	dbObj := &schema.Order{}

	err := st.db.NewSelect().
		Model(dbObj).
		Where("id = ?", id).
		Limit(1).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	obj, err := dbObj.ToCanonical()
	if err != nil {
		return nil, fmt.Errorf("conveting to canonical model: %w", err)
	}

	return &obj, nil
}

// GetOrdersForUser implements the storage.OrderReader interface.
func (st Storage) GetOrdersForUser(
	ctx context.Context,
	userID uuid.UUID,
	createdAtRangeStart, createdAtRangeEnd *time.Time,
	pageParams input.PageParams,
) ([]model.Order, error) {

	var dbObjs schema.Orders

	q := st.db.NewSelect().
		Model(&dbObjs).
		Where("user_id = ?", userID).
		Order("created_at ASC")

	if createdAtRangeStart != nil {
		q.Where("created_at >= ?", *createdAtRangeStart)
	}
	if createdAtRangeEnd != nil {
		q.Where("created_at <= ?", *createdAtRangeEnd)
	}
	paginateQuery(q, pageParams)

	if err := q.Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	objs, err := dbObjs.ToCanonical()
	if err != nil {
		return nil, fmt.Errorf("conveting to canonical models: %w", err)
	}

	return objs, nil
}

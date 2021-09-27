package schema

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/itiky/practicum-examples/04_pgsql/model"
	"github.com/uptrace/bun"
)

type (
	// Order is a DB representation of model.Order canonical model.
	Order struct {
		bun.BaseModel `bun:"orders,alias:o"`
		ID            uuid.UUID `bun:"id"`
		UserID        uuid.UUID `bun:"user_id"`
		Status        int       `bun:"status"`
		CreatedAt     time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
		DeletedAt     time.Time `bun:"deleted_at,soft_delete"`
	}

	Orders []Order
)

// NewOrderFromCanonical creates a new Order DB object from canonical model.
func NewOrderFromCanonical(obj model.Order) Order {
	return Order{
		ID:        obj.ID,
		UserID:    obj.UserID,
		Status:    obj.Status.Int(),
		CreatedAt: obj.CreatedAt,
		DeletedAt: obj.DeletedAt,
	}
}

// ToCanonical converts a DB object to canonical model.
func (o Order) ToCanonical() (model.Order, error) {
	obj := model.Order{
		ID:        o.ID,
		UserID:    o.UserID,
		Status:    model.NewOrderStatusFromInt(o.Status),
		CreatedAt: o.CreatedAt,
		DeletedAt: o.DeletedAt,
	}

	if err := obj.Status.Validate(); err != nil {
		return model.Order{}, fmt.Errorf("status: %w", err)
	}

	return obj, nil
}

// ToCanonical converts a DB object to canonical model.
func (o Orders) ToCanonical() ([]model.Order, error) {
	objs := make([]model.Order, 0, len(o))
	for dbObjIdx, dbObj := range o {
		obj, err := dbObj.ToCanonical()
		if err != nil {
			return nil, fmt.Errorf("dbObj [%d]: %w", dbObjIdx, err)
		}
		objs = append(objs, obj)
	}

	return objs, nil
}

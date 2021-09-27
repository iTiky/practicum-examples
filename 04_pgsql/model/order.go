package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/itiky/practicum-examples/04_pgsql/pkg/logging"
	"github.com/rs/zerolog"
)

// Order keeps order by user data.
type Order struct {
	ID        uuid.UUID   `json:"id" yaml:"id"`
	UserID    uuid.UUID   `json:"user_id" yaml:"user_id"`
	Status    OrderStatus `json:"status" yaml:"status"`
	CreatedAt time.Time   `json:"created_at" yaml:"created_at"`
	DeletedAt time.Time   `json:"deleted_at" yaml:"deleted_at"`
}

type OrderStatus string

const (
	OrderStatusCreated   OrderStatus = "CREATED"
	OrderStatusShipped   OrderStatus = "SHIPPED"
	OrderStatusDelivered OrderStatus = "DELIVERED"
	OrderStatusCompleted OrderStatus = "COMPLETED"
)

var (
	// orderStatusMap maps OrderStatus value to its int representation.
	orderStatusToIntMap = map[OrderStatus]int{
		OrderStatusCreated:   1,
		OrderStatusShipped:   2,
		OrderStatusDelivered: 3,
		OrderStatusCompleted: 4,
	}

	// orderStatusToStrMap maps OrderStatus value to its string representation.
	orderStatusToStrMap = map[int]OrderStatus{
		1: OrderStatusCreated,
		2: OrderStatusShipped,
		3: OrderStatusDelivered,
		4: OrderStatusCompleted,
	}
)

// NewOrderStatusFromInt returns OrderStatus by its int representation (might be invalid).
func NewOrderStatusFromInt(v int) OrderStatus {
	return orderStatusToStrMap[v]
}

// String implements the fmt.Stringer interface.
func (s OrderStatus) String() string {
	return string(s)
}

// Int returns enum value int representation.
func (s OrderStatus) Int() int {
	return orderStatusToIntMap[s]
}

// Validate validates enum value.
func (s OrderStatus) Validate() error {
	_, found := orderStatusToIntMap[s]
	if !found {
		return fmt.Errorf("unknown value: %v", s)
	}

	return nil
}

// GetLoggerContext enriches logger context with essential Order fields.
func (o Order) GetLoggerContext(logCtx zerolog.Context) zerolog.Context {
	logCtx = logCtx.Str("user_id", o.UserID.String())
	if o.ID != uuid.Nil {
		logCtx = logCtx.Str(logging.IDKey, o.ID.String())
	}

	return logCtx
}

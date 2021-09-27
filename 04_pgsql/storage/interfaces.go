//go:generate mockgen -source=interfaces.go -destination=./mock/storage.go -package=storagemock
package storage

import (
	"context"
	"io"
	"time"

	"github.com/google/uuid"
	"github.com/itiky/practicum-examples/04_pgsql/model"
	"github.com/itiky/practicum-examples/04_pgsql/pkg/input"
)

// UserWriter defines model.User create/update operations.
type UserWriter interface {
	io.Closer

	// CreateUser creates a new model.User.
	// Returns ErrAlreadyExists if user exists.
	CreateUser(ctx context.Context, user model.User) (model.User, error)
}

// UserReader defines model.User read operations.
type UserReader interface {
	io.Closer

	// GetUserByEmail returns model.User by its email if exists.
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
}

// OrderWriter defines model.Order create/update operations.
type OrderWriter interface {
	io.Closer

	// CreateOrder creates a new model.Order.
	CreateOrder(ctx context.Context, order model.Order) (model.Order, error)
}

// OrderReader defines model.Order read operations.
type OrderReader interface {
	io.Closer

	// GetOrderByID returns model.Order by its unique ID if exists.
	GetOrderByID(ctx context.Context, id uuid.UUID) (*model.Order, error)
	// GetOrdersForUser returns model.Order objects for specific user with optional time range filter.
	GetOrdersForUser(ctx context.Context, userID uuid.UUID, createdAtRangeStart, createdAtRangeEnd *time.Time, pageParams input.PageParams) ([]model.Order, error)
}

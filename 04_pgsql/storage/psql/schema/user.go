package schema

import (
	"time"

	"github.com/google/uuid"
	"github.com/itiky/practicum-examples/04_pgsql/model"
	"github.com/uptrace/bun"
)

// User is a DB representation of model.User canonical model.
type User struct {
	bun.BaseModel `bun:"users,alias:u"`
	ID            uuid.UUID `bun:"id"`
	Name          string    `bun:"name"`
	Email         string    `bun:"email"`
	Phone         string    `bun:"phone"`
	Region        string    `bun:"region"`
	CreatedAt     time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
	DeletedAt     time.Time `bun:"deleted_at,soft_delete"`
}

// NewUserFromCanonical creates a new User DB object from canonical model.
func NewUserFromCanonical(obj model.User) User {
	return User{
		ID:        obj.ID,
		Name:      obj.Name,
		Email:     obj.Email,
		Phone:     obj.PhoneNumber,
		Region:    obj.RegionCode,
		CreatedAt: obj.CreatedAt,
		UpdatedAt: obj.UpdatedAt,
		DeletedAt: obj.DeletedAt,
	}
}

// ToCanonical converts a DB object to canonical model.
func (u User) ToCanonical() (model.User, error) {
	return model.User{
		ID:          u.ID,
		Name:        u.Name,
		Email:       u.Email,
		PhoneNumber: u.Phone,
		RegionCode:  u.Region,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
		DeletedAt:   u.DeletedAt,
	}, nil
}

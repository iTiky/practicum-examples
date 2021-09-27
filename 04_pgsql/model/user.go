package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/itiky/practicum-examples/04_pgsql/pkg/logging"
	"github.com/rs/zerolog"
)

// User keeps user data.
type User struct {
	ID          uuid.UUID `json:"id" yaml:"id"`
	Name        string    `json:"name" yaml:"name"`
	Email       string    `json:"email" yaml:"email"`
	PhoneNumber string    `json:"phone_number" yaml:"phone_number"`
	RegionCode  string    `json:"region_code" yaml:"region_code"`
	CreatedAt   time.Time `json:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" yaml:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at" yaml:"deleted_at"`
}

// GetLoggerContext enriches logger context with essential User fields.
func (u User) GetLoggerContext(logCtx zerolog.Context) zerolog.Context {
	logCtx = logCtx.Str("email", u.Email)
	if u.ID != uuid.Nil {
		logCtx = logCtx.Str(logging.IDKey, u.ID.String())
	}

	return logCtx
}

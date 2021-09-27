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
	"github.com/itiky/practicum-examples/04_pgsql/storage/psql/schema"
	"github.com/uptrace/bun/driver/pgdriver"
)

// CreateUser implements the storage.UserWriter interface.
func (st Storage) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	dbObj := schema.NewUserFromCanonical(user)
	dbObj.ID, dbObj.CreatedAt, dbObj.UpdatedAt, dbObj.DeletedAt = uuid.Nil, time.Time{}, time.Time{}, time.Time{}

	_, err := st.db.NewInsert().
		Model(&dbObj).
		Returning("*").
		Exec(ctx)
	if err != nil {
		pgErr := &pgdriver.Error{}
		if errors.As(err, pgErr) {
			if pgErr.IntegrityViolation() {
				return model.User{}, pkg.ErrAlreadyExists
			}
		}
		return model.User{}, err
	}

	obj, err := dbObj.ToCanonical()
	if err != nil {
		return model.User{}, fmt.Errorf("conveting to canonical model: %w", err)
	}

	return obj, nil
}

// GetUserByEmail implements the storage.UserReader interface.
func (st Storage) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	dbObj := &schema.User{}

	err := st.db.NewSelect().
		Model(dbObj).
		Where("email = ?", email).
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

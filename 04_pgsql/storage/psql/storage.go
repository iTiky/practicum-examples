package psql

import (
	"context"
	"database/sql"
	"fmt"
	"runtime"

	"github.com/itiky/practicum-examples/04_pgsql/storage/psql/migrations"
	"github.com/itiky/practicum-examples/04_pgsql/storage/psql/schema"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/migrate"
)

const (
	serviceName = "psql"

	dbTableLoggingKey     = "db-table"
	dbOperationLoggingKey = "db-operation"
)

type (
	Storage struct {
		config Config
		db     *bun.DB
	}

	StorageOption func(st *Storage) error
)

// WithConfig overrides default Storage config.
func WithConfig(config Config) StorageOption {
	return func(st *Storage) error {
		st.config = config

		return nil
	}
}

// New creates a new Storage with custom options.
func New(opts ...StorageOption) (*Storage, error) {
	st := &Storage{
		config: NewDefaultConfig(),
	}
	for optIdx, opt := range opts {
		if err := opt(st); err != nil {
			return nil, fmt.Errorf("applying option [%d]: %w", optIdx, err)
		}
	}

	if err := st.config.Validate(); err != nil {
		return nil, fmt.Errorf("config validation: %w", err)
	}

	sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(st.config.DSN)))
	maxOpenConnections := 4 * runtime.GOMAXPROCS(0)

	st.db = bun.NewDB(sqlDb, pgdialect.New())
	st.db.AddQueryHook(newQueryHook(st))
	st.db.SetMaxOpenConns(maxOpenConnections)
	st.db.SetMaxIdleConns(maxOpenConnections)
	st.db.RegisterModel(
		(*schema.User)(nil),
		(*schema.Order)(nil),
	)

	if err := st.db.Ping(); err != nil {
		return nil, fmt.Errorf("ping for DSN (%s) failed: %w", st.config.DSN, err)
	}

	return st, nil
}

// Close closes DB connection.
func (st Storage) Close() error {
	if st.db == nil {
		return nil
	}

	return st.db.Close()
}

// Migrate performs DB migrations.
func (st Storage) Migrate(ctx context.Context) error {
	logger := st.Logger(withOperation("migration"))

	migrations, err := migrations.GetMigrations()
	if err != nil {
		return err
	}

	migration := migrate.NewMigrator(st.db, migrations)

	if err := migration.Init(ctx); err != nil {
		return fmt.Errorf("initialising migration: %w", err)
	}

	res, err := migration.Migrate(ctx)
	if err != nil {
		return fmt.Errorf("performing migration: %w", err)
	}

	logger.Info().Msgf("Migration applied: %s", res.Migrations.LastGroup().String())

	return nil
}

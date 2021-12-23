package common

import (
	"fmt"

	"github.com/itiky/practicum-examples/04_pgsql/service/order"
	"github.com/itiky/practicum-examples/04_pgsql/service/user"
	"github.com/itiky/practicum-examples/04_pgsql/storage/psql"
)

// Config combines sub-configs for all services, storages and providers.
type Config struct {
	UserService user.Config `mapstructure:"user_service"`
	PSQLStorage psql.Config `mapstructure:"psql_storage"`
}

// BuildPsqlStorage builds psql.Storage dependency.
func (c Config) BuildPsqlStorage() (*psql.Storage, error) {
	st, err := psql.New(
		psql.WithConfig(c.PSQLStorage),
	)
	if err != nil {
		return nil, fmt.Errorf("building psql storage: %w", err)
	}

	return st, nil
}

// BuildUserService builds user.Processor dependency.
func (c Config) BuildUserService() (*user.Processor, error) {
	st, err := c.BuildPsqlStorage()
	if err != nil {
		return nil, err
	}

	svc, err := user.New(
		user.WithConfig(c.UserService),
		user.WithUserWriter(st),
	)
	if err != nil {
		return nil, fmt.Errorf("building user service: %w", err)
	}

	return svc, nil
}

// BuildOrderService builds test.Processor dependency.
func (c Config) BuildOrderService() (*order.Processor, error) {
	st, err := c.BuildPsqlStorage()
	if err != nil {
		return nil, err
	}

	svc, err := order.New(
		order.WithOrderStorage(st),
	)
	if err != nil {
		return nil, fmt.Errorf("building test service: %w", err)
	}

	return svc, nil
}

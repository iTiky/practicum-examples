package fixtures

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"text/template"
	"time"

	"github.com/itiky/practicum-examples/04_pgsql/storage/psql/schema"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dbfixture"
)

// Fixtures keeps all fixture objects.
type Fixtures struct {
	Users  []schema.User
	Orders []schema.Order
}

func (f *Fixtures) appendUser(obj interface{}) error {
	dbObj, ok := obj.(*schema.User)
	if !ok {
		return fmt.Errorf("user: type assert failed: %T", obj)
	}
	if dbObj == nil {
		return fmt.Errorf("user: type assert failed: nil")
	}
	f.Users = append(f.Users, *dbObj)

	return nil
}

func (f *Fixtures) appendOrder(obj interface{}) error {
	dbObj, ok := obj.(*schema.Order)
	if !ok {
		return fmt.Errorf("test: type assert failed: %T", obj)
	}
	if dbObj == nil {
		return fmt.Errorf("test: type assert failed: nil")
	}
	f.Orders = append(f.Orders, *dbObj)

	return nil
}

// LoadFixtures load fixtures to DB and returns DB objects aggregate.
func LoadFixtures(ctx context.Context, db *bun.DB) (Fixtures, error) {
	type fixturesAppender struct {
		id     string
		append func(obj interface{}) error
	}

	fixtureManager := dbfixture.New(
		db,
		dbfixture.WithTemplateFuncs(template.FuncMap{
			"now": func() string {
				return time.Now().UTC().Format(time.RFC3339Nano)
			},
		}),
	)

	err := fixtureManager.Load(
		ctx,
		os.DirFS(getFixturesDir()),
		"users.yaml", "orders.yaml",
	)
	if err != nil {
		return Fixtures{}, fmt.Errorf("loading fixtures: %w", err)
	}

	fixtures := Fixtures{}
	appenders := []fixturesAppender{
		{id: "User.bob", append: fixtures.appendUser},
		{id: "User.alice", append: fixtures.appendUser},
		{id: "Order.bobOrder_1", append: fixtures.appendOrder},
		{id: "Order.bobOrder_2", append: fixtures.appendOrder},
		{id: "Order.bobOrder_3", append: fixtures.appendOrder},
	}
	for _, appender := range appenders {
		obj, err := fixtureManager.Row(appender.id)
		if err != nil {
			return Fixtures{}, fmt.Errorf("reading fixtures row (%s): %w", appender.id, err)
		}
		if obj == nil {
			return Fixtures{}, fmt.Errorf("reading fixtures row (%s): nil", appender.id)
		}
		if err := appender.append(obj); err != nil {
			return Fixtures{}, fmt.Errorf("appending fixtures row (%s): %w", appender.id, err)
		}
	}

	return fixtures, nil
}

// getFixturesDir returns current file directory.
func getFixturesDir() string {
	_, filePath, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}

	return filepath.Dir(filePath)
}

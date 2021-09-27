package psql

import (
	"context"
	"testing"
	"time"

	"github.com/itiky/practicum-examples/04_pgsql/storage/psql/fixtures"
	"github.com/itiky/practicum-examples/04_pgsql/testutils"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite

	container *testutils.PostgreSQLContainer
	storage   *Storage
	fixtures  fixtures.Fixtures

	ctx context.Context
}

func (s *TestSuite) SetupSuite() {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer ctxCancel()

	c, err := testutils.NewPostgreSQLContainer(ctx)
	s.Require().NoError(err)

	stCfg := NewDefaultConfig()
	stCfg.DSN = c.GetDSN()

	st, err := New(WithConfig(stCfg))
	s.Require().NoError(err)

	s.Require().NoError(st.Migrate(ctx))

	fixtures, err := fixtures.LoadFixtures(ctx, st.db)
	s.Require().NoError(err)

	s.ctx = context.TODO()
	s.container = c
	s.storage = st
	s.fixtures = fixtures
}

func (s *TestSuite) TearDownSuite() {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer ctxCancel()

	s.Require().NoError(s.storage.Close())
	s.Require().NoError(s.container.Terminate(ctx))
}

func TestSuite_PostgreSQLStorage(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

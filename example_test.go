package postgres_test_suite

import (
	"context"
	"fmt"
	"testing"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/jackc/pgx/v5"
)

type ExampleTestSuite struct {
	*PostgresTestSuite
	port uint32
}

func TestExampleTestSuite(t *testing.T) {
	port := uint32(46462)
	testSuite := &ExampleTestSuite{
		PostgresTestSuite: &PostgresTestSuite{PostgresConf: embeddedpostgres.DefaultConfig().Port(port)},
		port:              port,
	}
	suite.Run(t, testSuite)
}

func (testSuite *ExampleTestSuite) TestEmbeddedPostgresIsCreated() {
	ctx := context.Background()
	connString := fmt.Sprintf("postgresql://postgres:postgres@localhost:%d/postgres", testSuite.port)
	conn, err := pgx.Connect(ctx, connString)
	if assert.NoError(testSuite.T(), err) {
		assert.NoError(testSuite.T(), conn.Ping(ctx))
	}

	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			testSuite.T().Error(err)
		}
	}(conn, ctx)
}

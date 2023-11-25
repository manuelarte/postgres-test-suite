package postgres_test_suite

import (
	"context"
	"fmt"
	"testing"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EPClosedAfter struct {
	*PostgresTestSuite
	port uint32
}

func TestEPClosedAfter(t *testing.T) {
	port := uint32(46355)
	testSuite := &EPClosedAfter{
		PostgresTestSuite: &PostgresTestSuite{PostgresConf: embeddedpostgres.DefaultConfig().Port(port)},
		port:              port,
	}
	suite.Run(t, testSuite)
	ctx := context.Background()
	connString := fmt.Sprintf("postgresql://postgres:postgres@localhost:%d/postgres", testSuite.port)
	_, err := pgx.Connect(ctx, connString)
	assert.Error(testSuite.T(), err, "Embedded postgres should be close after a test panic")
}

func (suite *EPClosedAfter) TestEmpty() {

}

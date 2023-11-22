package postgres_test_suite

import (
	"context"
	"fmt"
	"net"
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
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Error(err)
	}
	port := uint32(l.Addr().(*net.TCPAddr).Port)
	testSuite := &EPClosedAfter{
		PostgresTestSuite: &PostgresTestSuite{PostgresConf: embeddedpostgres.DefaultConfig().Port(port)},
		port:              port,
	}
	suite.Run(t, testSuite)
	ctx := context.Background()
	connString := fmt.Sprintf("postgresql://postgres:postgres@localhost:%d/postgres", testSuite.port)
	_, err = pgx.Connect(ctx, connString)
	assert.Error(testSuite.T(), err, "Embedded postgres should be close after a test panic")
}

func (suite *EPClosedAfter) TestEmpty() {

}

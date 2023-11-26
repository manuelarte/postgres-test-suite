package postgres_test_suite

import (
	"os"
	"os/signal"
	"syscall"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/stretchr/testify/suite"
)

type PostgresTestSuite struct {
	suite.Suite
	PostgresConf embeddedpostgres.Config
	ep           *embeddedpostgres.EmbeddedPostgres
	// add hooks
}

func (psuite *PostgresTestSuite) SetupSuite() {
	psuite.ep = psuite.getEmbeddedPostgres()
	err := psuite.ep.Start()
	if err != nil {
		psuite.T().Fatal(err)
	}
	go closeWithSignal(psuite, psuite.ep)
}

func (psuite *PostgresTestSuite) TearDownSuite() {
	closeEmbeddedPostgres(psuite, psuite.ep)
}

func (psuite *PostgresTestSuite) getEmbeddedPostgres() *embeddedpostgres.EmbeddedPostgres {
	return embeddedpostgres.NewDatabase(psuite.PostgresConf)
}

func closeWithSignal(psuite *PostgresTestSuite, embeddedPostgres *embeddedpostgres.EmbeddedPostgres) {
	c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-c
	closeEmbeddedPostgres(psuite, embeddedPostgres)
	close(c)
}

func closeEmbeddedPostgres(psuite *PostgresTestSuite, ep *embeddedpostgres.EmbeddedPostgres) {
	err := ep.Stop()
	if err != nil {
		psuite.Error(err, "Can't stop embedded postgres, please kill the process manually, e.g. kill $(lsof -t -i:$PORT)")
	}
}

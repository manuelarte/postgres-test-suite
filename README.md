[![Go](https://github.com/manuelarte/postgres-test-suite/actions/workflows/go.yml/badge.svg)](https://github.com/manuelarte/postgres-test-suite/actions/workflows/go.yml)

# PostgreSQL Test suite #

When doing a test that involves reading/writing to a database, it's good to test as close as possible to the production code.
This library allows you to instantiate a PostgreSQL database for your test suite, and it will automatically start and stop at the beginning and enf of your test.
It works using the test [suite](https://pkg.go.dev/github.com/stretchr/testify/suite) functionality in GO, in which the database is starting at the beginning of the test suite, then all the tests are run, and it will stop at the end.

To use this library:

```
go get github.com/manuelarte/postgres-test-suite
````

## Example ##

Check the example in the file [example_test.go](example_test.go):

```go
func TestExampleTestSuite(t *testing.T) {
	port := // port where you want to run your postgres for testing
	testSuite := &ExampleTestSuite{
		PostgresTestSuite: &PostgresTestSuite{PostgresConf: embeddedpostgres.DefaultConfig().Port(port)},
	}
	suite.Run(t, testSuite)
}

func (testSuite *ExampleTestSuite) Test...() {
    ctx := context.Background()
    connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", username, password, host, port, dbname)
    conn, err := pgx.Connect(ctx, connString)
    require.NoError(err)
    defer conn.Close(ctx)
	// do your test
}
```

There is no cleaning of the database between tests inside the same test suite, so it's part of your test suite to clean it to make the isolated of the rest of the tests. For that you can use the [BeforeTest](https://pkg.go.dev/github.com/stretchr/testify/suite#BeforeTest) or [AfterTests](https://pkg.go.dev/github.com/stretchr/testify/suite#AfterTest) functions provided in test suite.

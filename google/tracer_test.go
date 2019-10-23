package google

import (
	"context"
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"

	"github.com/sour-is/instrumentedsql"
)

// WrapDriverGoogle demonstrates how to call wrapDriver and register a new driver.
// This example uses MySQL and google tracing to illustrate this
func ExampleWrapDriver_google() {
	logger := instrumentedsql.LoggerFunc(func(ctx context.Context, msg string, keyvals ...interface{}) {
		log.Printf("%s %v", msg, keyvals)
	})

	sql.Register("instrumented-mysql", instrumentedsql.WrapDriver(mysql.MySQLDriver{}, instrumentedsql.WithTracer(NewTracer()), instrumentedsql.WithLogger(logger)))
	db, err := sql.Open("instrumented-mysql", "connString")

	// Proceed to handle connection errors and use the database as usual
	_, _ = db, err
}

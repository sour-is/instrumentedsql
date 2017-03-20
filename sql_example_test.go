package instrumentedsql_test

import (
	"database/sql"
	"context"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/mattn/go-sqlite3"
	"github.com/ExpansiveWorlds/instrumentedsql"
	"github.com/ExpansiveWorlds/instrumentedsql/google"
	"github.com/ExpansiveWorlds/instrumentedsql/opentracing"
)

// WrapDriverGoogle demonstrates how to call wrapDriver and register a new driver.
// This example uses MySQL and google tracing to illustrate this
func ExampleWrapDriver_google() {
	logger := func(ctx context.Context, msg string, keyvals ...interface{}) {
		log.Printf("%s %v", msg, keyvals)
	}

	sql.Register("instrumented-mysql", instrumentedsql.WrapDriver(mysql.MySQLDriver{}, instrumentedsql.WithTracer(google.NewTracer()), instrumentedsql.WithLogger(instrumentedsql.NewFuncLogger(logger))))
	db, err := sql.Open("instrumented-mysql", "connString")

	// Proceed to handle connection errors and use the database as usual
	_, _ = db, err
}


// WrapDriverOpentracing demonstrates how to call wrapDriver and register a new driver.
// This example uses MySQL and opentracing to illustrate this
func ExampleWrapDriver_opentracing() {
	logger := func(ctx context.Context, msg string, keyvals ...interface{}) {
		log.Printf("%s %v", msg, keyvals)
	}

	sql.Register("instrumented-mysql", instrumentedsql.WrapDriver(mysql.MySQLDriver{}, instrumentedsql.WithTracer(opentracing.NewTracer()), instrumentedsql.WithLogger(instrumentedsql.NewFuncLogger(logger))))
	db, err := sql.Open("instrumented-mysql", "connString")

	// Proceed to handle connection errors and use the database as usual
	_, _ = db, err
}

// WrapDriverJustLogging demonstrates how to call wrapDriver and register a new driver.
// This example uses sqlite, but does not trace, but merely logs all calls
func ExampleWrapDriver_justLogging() {
	logger := func(ctx context.Context, msg string, keyvals ...interface{}) {
		log.Printf("%s %v", msg, keyvals)
	}

	sql.Register("instrumented-mysql", instrumentedsql.WrapDriver(&sqlite3.SQLiteDriver{}, instrumentedsql.WithLogger(instrumentedsql.NewFuncLogger(logger))))
	db, err := sql.Open("instrumented-mysql", "connString")

	// Proceed to handle connection errors and use the database as usual
	_, _ = db, err
}
package instrumentedsql

import (
	"context"
	"database/sql"
	"log"

	sqlite3 "github.com/mattn/go-sqlite3"
)

// WrapDriverJustLogging demonstrates how to call wrapDriver and register a new driver.
// This example uses sqlite, but does not trace, but merely logs all calls
func ExampleWrapDriver_justLogging() {
	logger := LoggerFunc(func(ctx context.Context, msg string, keyvals ...interface{}) {
		log.Printf("%s %v", msg, keyvals)
	})

	sql.Register("instrumented-sqlite", WrapDriver(&sqlite3.SQLiteDriver{}, WithLogger(logger)))
	db, err := sql.Open("instrumented-sqlite", "connString")

	// Proceed to handle connection errors and use the database as usual
	_, _ = db, err
}

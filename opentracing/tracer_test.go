package opentracing

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/sour-is/instrumentedsql"

	ot "github.com/opentracing/opentracing-go"
)

func TestSpanWithParent(t *testing.T) {
	ctx := ot.ContextWithSpan(
		context.Background(),
		ot.GlobalTracer().StartSpan("some_span"),
	)

	tr := NewTracer()
	span := tr.GetSpan(ctx)
	span.SetLabel("key", "value")

	child := span.NewChild("child")
	child.SetLabel("child_key", "child_value")
	child.Finish()

	span.Finish()
}

func TestSpanWithoutParent(t *testing.T) {
	ctx := context.Background() // Background has no span
	tr := NewTracer()
	span := tr.GetSpan(ctx)
	span.SetLabel("key", "value")

	child := span.NewChild("child")
	child.SetLabel("child_key", "child_value")
	child.Finish()

	span.Finish()
}

// WrapDriverOpentracing demonstrates how to call wrapDriver and register a new driver.
// This example uses MySQL and opentracing to illustrate this
func ExampleWrapDriver_opentracing() {
	logger := instrumentedsql.LoggerFunc(func(ctx context.Context, msg string, keyvals ...interface{}) {
		log.Printf("%s %v", msg, keyvals)
	})

	sql.Register(
		"instrumented-mysql",
		instrumentedsql.WrapDriver(
			mysql.MySQLDriver{},
			instrumentedsql.WithTracer(NewTracer()),
			instrumentedsql.WithLogger(logger),
		),
	)

	db, err := sql.Open("instrumented-mysql", "connString")

	// Proceed to handle connection errors and use the database as usual
	_, _ = db, err
}

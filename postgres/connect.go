package postgres

import (
	"context"
	"database/sql"

	"github.com/XSAM/otelsql"
	_ "github.com/jackc/pgx/v5/stdlib"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

	func InitPostgresWithTracing(ctx context.Context, dsn string) *sql.DB {
		db, err := otelsql.Open("pgx", dsn, otelsql.WithAttributes(
			semconv.DBSystemPostgreSQL,
		))
		if err != nil {
			panic(err)
		}

		if err = db.PingContext(ctx); err != nil {
			panic(err)
		}
		return db
	}

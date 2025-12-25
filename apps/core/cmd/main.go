package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/LeCarteloo/ecommerce-tanstack-start-go/internal/env"
	"github.com/jackc/pgx/v5"
)

func main() {
	dbConfig := dbConfig{
		dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=postgres password=postgres dbname=ecommerce sslmode=disable"),
	}

	ctx := context.Background()
	dbConn, err := pgx.Connect(ctx, dbConfig.dsn)
	if err != nil {
		slog.Error("unable to connect to database", "error", err)
		os.Exit(1)
	}

	app := application{
		config: config{
			addr:     ":8080",
			dbConfig: dbConfig,
		},
		dbConn: dbConn,
	}

	app.run(app.mount())
}

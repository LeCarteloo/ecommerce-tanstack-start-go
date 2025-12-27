package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/LeCarteloo/ecommerce-tanstack-start-go/internal/env"
	"github.com/jackc/pgx/v5"
)

func main() {
	setupLogger()

	dbConfig := dbConfig{
		dsn: env.GetString("GOOSE_DBSTRING"),
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

func setupLogger() {
	var handler slog.Handler

	if env.GetString("ENV") == "prod" {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}

	slog.SetDefault(slog.New(handler))
}

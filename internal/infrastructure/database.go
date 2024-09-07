package infrastructure

import (
	"context"
	"log/slog"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

func DatabaseProvider(lc fx.Lifecycle) *sqlx.DB {
	db, err := sqlx.Connect("postgres", os.Getenv("DB_URL"))

	if err != nil {
		slog.Error("Database error", "error", err)
		os.Exit(1)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			slog.Info("Database connection closed")
			return db.Close()
		},
	})

	return db
}

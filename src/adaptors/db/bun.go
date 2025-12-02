package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var DB *bun.DB

func NewBun() (*bun.DB, error) {
	if DB != nil {
		return DB, nil
	}

	dsn := os.Getenv("postgressUrl")
	if dsn == "" {
		return nil, fmt.Errorf("postgressUrl environment variable not set")
	}

	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN(dsn),
	))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sqldb.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	DB = bun.NewDB(sqldb, pgdialect.New())
	return DB, nil
}

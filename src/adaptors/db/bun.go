package db

import (
	"context"
	"database/sql"
	"errors"
	"go-api/ports/db"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var dbInstance *bunDB

type bunDB struct {
	bun         *bun.DB
	notifyFuncs []func(ctx context.Context) error
}

func (b *bunDB) GetDB(ctx context.Context) (*bun.DB, error) {
	if b.bun != nil {
		return b.bun, nil
	}

	dsn := os.Getenv("POSTGRES_URL")
	if dsn == "" {
		return nil, errors.New("POSTGRES_URL environment variable not set")
	}

	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN(dsn),
	))

	if err := sqldb.PingContext(ctx); err != nil {
		return nil, errors.New("failed to connect to postgres" + err.Error())
	}

	b.bun = bun.NewDB(sqldb, pgdialect.New())

	for _, f := range b.notifyFuncs {
		f(ctx)
	}

	return b.bun, nil
}

func (b *bunDB) NotifyFirstAvailable(f func(ctx context.Context) error) {
	b.notifyFuncs = append(b.notifyFuncs, f)
}

func NewBun() db.DB {
	if dbInstance != nil {
		return dbInstance
	}

	dbInstance = &bunDB{}
	return dbInstance
}

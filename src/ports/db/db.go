package db

import (
	"context"

	"github.com/uptrace/bun"
)

type DB interface {
	GetDB(ctx context.Context) (*bun.DB, error)
	NotifyFirstAvailable(f func(ctx context.Context) error)
}

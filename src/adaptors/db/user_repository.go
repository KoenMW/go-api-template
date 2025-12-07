package db

import (
	"context"
	"go-api/domain/model"

	"github.com/uptrace/bun"
)

func NewUserRepository(db *bun.DB, ctx context.Context) *Repository[model.User] {
	return NewRepository[model.User](db, ctx)
}

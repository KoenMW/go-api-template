package db

import (
	"go-api/domain/model"

	"github.com/uptrace/bun"
)

func NewUserRepository(db *bun.DB) *Repository[model.User] {
	return NewRepository[model.User](db)
}

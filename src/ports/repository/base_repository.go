package repository

import (
	"context"
	"go-api/domain/model"

	"github.com/google/uuid"
)

type BaseRepository[T model.BaseEntity] interface {
	Create(ctx context.Context, question T) (T, error)
	GetByID(ctx context.Context, id uuid.UUID) (T, error)
	List(ctx context.Context, perPage int, page int) ([]T, error)
	Update(ctx context.Context, question T) (T, error)
	Delete(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
}

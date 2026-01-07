package service

import (
	"context"
	"go-api/domain/model"
	"go-api/ports/repository"

	"github.com/google/uuid"
)

type BaseService[T model.BaseEntity, DTO model.BaseDTO[T], CreateDTO model.BaseCreateDTO[T], Repository repository.BaseRepository[T]] interface {
	Create(ctx context.Context, dto CreateDTO) (DTO, error)
	GetByID(ctx context.Context, id uuid.UUID) (DTO, error)
	Update(ctx context.Context, id uuid.UUID, dto DTO) (DTO, error)
	Delete(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
	List(ctx context.Context, perPage int, page int) ([]DTO, error)
}

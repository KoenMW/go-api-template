package service

import (
	"context"
	"go-api/domain/helper"
	"go-api/domain/model"
	"go-api/ports/repository"

	"github.com/google/uuid"
)

type BaseService[T model.BaseEntity, DTO model.BaseDTO[T], CreateDTO model.BaseCreateDTO[T], Repository repository.BaseRepository[T]] struct {
	repo Repository
}

func NewBaseService[T model.BaseEntity, DTO model.BaseDTO[T], CreateDTO model.BaseCreateDTO[T], Repository repository.BaseRepository[T]](
	repo Repository,
) *BaseService[T, DTO, CreateDTO, Repository] {
	return &BaseService[T, DTO, CreateDTO, Repository]{repo: repo}
}

func (s *BaseService[T, DTO, CreateDTO, Repository]) Create(ctx context.Context, dto CreateDTO) (DTO, error) {
	var zero DTO
	if err := dto.Validate(); err != nil {
		return zero, err
	}

	entity := model.NewEntity[T]()

	dto.ApplyToEntity(entity)

	entity.SetID(uuid.New())
	entity.SetCreatedAt()
	entity.SetUpdatedAt()

	saved, err := s.repo.Create(ctx, entity)
	if err != nil {
		return zero, err
	}

	return helper.EntityToDTO[T, DTO](saved), nil
}

func (s *BaseService[T, DTO, CreateDTO, Repository]) GetByID(ctx context.Context, id uuid.UUID) (DTO, error) {
	var zero DTO
	e, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return zero, err
	}
	var dto DTO = model.NewEntity[DTO]()
	dto.RecieveEntity(e)

	return dto, nil
}

func (s *BaseService[T, DTO, CreateDTO, Repository]) Update(ctx context.Context, id uuid.UUID, dto DTO) (DTO, error) {
	var zero DTO
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return zero, err
	}

	var entity T = model.NewEntity[T]()
	dto.ApplyToEntity(entity)
	entity.SetID(id)
	entity.SetUpdatedAt()
	e, err := s.repo.Update(ctx, entity)
	if err != nil {
		return zero, err
	}

	return helper.EntityToDTO[T, DTO](e), nil
}

func (s *BaseService[T, DTO, CreateDTO, Repository]) List(ctx context.Context, perPage int, page int) ([]DTO, error) {
	entities, err := s.repo.List(ctx, perPage, page)
	if err != nil {
		return nil, err
	}
	var results []DTO
	for _, entity := range entities {
		results = append(results, helper.EntityToDTO[T, DTO](entity))
	}
	return results, nil
}

func (s *BaseService[T, DTO, CreateDTO, Repository]) Delete(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	return s.repo.Delete(ctx, id)
}

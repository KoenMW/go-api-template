package db

import (
	"context"
	"errors"
	"go-api/domain/core"
	"go-api/domain/model"

	dbPort "go-api/ports/db"

	"github.com/google/uuid"
)

type Repository[T model.BaseEntity] struct {
	db   dbPort.DB
	zero T
}

func NewRepository[T model.BaseEntity](db dbPort.DB) *Repository[T] {
	repo := &Repository[T]{db: db, zero: model.NewEntity[T]()}

	return repo
}

func (r *Repository[T]) Create(ctx context.Context, entity T) (T, error) {
	db, err := r.db.GetDB(ctx)
	if err != nil {
		return r.zero, err
	}

	_, err = db.NewInsert().Model(entity).Exec(ctx)
	return entity, err
}

func (r *Repository[T]) GetByID(ctx context.Context, id uuid.UUID) (T, error) {
	db, err := r.db.GetDB(ctx)
	if err != nil {
		return r.zero, err
	}

	entity := model.NewEntity[T]()
	err = db.NewSelect().
		Model(entity).
		Where(core.WhereId, id).
		Scan(ctx)

	if err != nil {
		return entity, err
	}

	return entity, nil
}

func (r *Repository[T]) Update(ctx context.Context, entity T) (T, error) {
	db, err := r.db.GetDB(ctx)
	if err != nil {
		return r.zero, err
	}

	_, err = db.NewUpdate().Model(entity).Where(
		core.WhereId, entity.GetId(),
	).Exec(ctx)
	return entity, err
}

func (r *Repository[T]) Delete(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	db, err := r.db.GetDB(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	res, err := db.NewDelete().
		Model(r.zero).
		Where(core.WhereId, id).
		Exec(ctx)
	if err != nil {
		return id, err
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		return id, errors.New("no rows deleted")
	}

	return id, nil
}

func (r *Repository[T]) List(ctx context.Context, perPage, page int) ([]T, error) {
	db, err := r.db.GetDB(ctx)
	if err != nil {
		return nil, err
	}

	var entities []T

	err = db.NewSelect().
		Model(&entities).
		Limit(perPage).
		Offset((page - 1) * perPage).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return entities, nil
}

func (r *Repository[T]) Count(ctx context.Context) (int, error) {
	db, err := r.db.GetDB(ctx)
	if err != nil {
		return 0, err
	}

	var count int
	err = db.NewSelect().Model(new(T)).ColumnExpr("count(*)").Scan(ctx, &count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

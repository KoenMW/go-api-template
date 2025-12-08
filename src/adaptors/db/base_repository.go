package db

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Repository[T any] struct {
	db   *bun.DB
	ctx  context.Context
	zero *T
}

func NewRepository[T any](db *bun.DB, ctx context.Context) *Repository[T] {
	var zero T
	repo := &Repository[T]{db: db, ctx: ctx, zero: &zero}

	_, err := db.NewCreateTable().
		Model(repo.zero).
		IfNotExists().
		Exec(ctx)

	if err != nil {
		panic(fmt.Errorf("failed creating users table: %w", err))
	}

	return repo
}

func (r *Repository[T]) Create(entity *T) (*T, error) {
	_, err := r.db.NewInsert().Model(entity).Exec(r.ctx)
	return entity, err
}

func (r *Repository[T]) GetByID(id uuid.UUID) (*T, error) {
	var entity T
	err := r.db.NewSelect().
		Model(&entity).
		Where("id = ?", id).
		Scan(r.ctx)

	if err != nil {
		return &entity, err
	}

	return &entity, nil
}

func (r *Repository[T]) Update(entity *T) (*T, error) {
	_, err := r.db.NewUpdate().Model(entity).Exec(r.ctx)
	return entity, err
}

func (r *Repository[T]) Delete(id uuid.UUID) (uuid.UUID, error) {
	res, err := r.db.NewDelete().
		Model(r.zero).
		Where("id = ?", id).
		Exec(r.ctx)
	if err != nil {
		return id, err
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		return id, fmt.Errorf("not found")
	}

	return id, nil
}

func (r *Repository[T]) List(perPage, page int) ([]T, error) {
	var entities []T

	err := r.db.NewSelect().
		Model(&entities).
		Limit(perPage).
		Offset((page - 1) * perPage).
		Scan(r.ctx)

	if err != nil {
		return nil, err
	}

	return entities, nil
}

func (r *Repository[T]) Count() (int, error) {
	var count int
	err := r.db.NewSelect().Model(new(T)).ColumnExpr("count(*)").Scan(r.ctx, &count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

package db

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Repository[T any] struct {
	db   *bun.DB
	zero *T
}

func NewRepository[T any](db *bun.DB) *Repository[T] {
	var zero T
	repo := &Repository[T]{db: db, zero: &zero}

	_, err := db.NewCreateTable().
		Model(repo.zero).
		IfNotExists().
		Exec(context.Background())

	if err != nil {
		panic(fmt.Errorf("failed creating users table: %w", err))
	}

	return repo
}

func (r *Repository[T]) Create(ctx context.Context, entity *T) (*T, error) {
	_, err := r.db.NewInsert().Model(entity).Exec(ctx)
	return entity, err
}

func (r *Repository[T]) GetByID(ctx context.Context, id uuid.UUID) (*T, error) {
	var entity T
	err := r.db.NewSelect().
		Model(&entity).
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		return &entity, err
	}

	return &entity, nil
}

func (r *Repository[T]) Update(ctx context.Context, entity *T) (*T, error) {
	_, err := r.db.NewUpdate().Model(entity).Exec(ctx)
	return entity, err
}

func (r *Repository[T]) Delete(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	res, err := r.db.NewDelete().
		Model(r.zero).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return id, err
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		return id, fmt.Errorf("not found")
	}

	return id, nil
}

func (r *Repository[T]) List(ctx context.Context, perPage, page int) ([]T, error) {
	var entities []T

	err := r.db.NewSelect().
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
	var count int
	err := r.db.NewSelect().Model(new(T)).ColumnExpr("count(*)").Scan(ctx, &count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

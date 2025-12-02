package db

import (
	"context"
	"fmt"
	"go-api/domain/model"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type UserRepository struct {
	db  *bun.DB
	ctx context.Context
}

func NewUserRepository(db *bun.DB, ctx context.Context) *UserRepository {
	repo := &UserRepository{db: db, ctx: ctx}

	_, err := db.NewCreateTable().
		Model((*model.User)(nil)).
		IfNotExists().
		Exec(ctx)

	if err != nil {
		panic(fmt.Errorf("failed creating users table: %w", err))
	}

	return repo
}

func (r *UserRepository) CreateUser(u *model.CreateUserDTO) (*model.UserDTO, error) {

	user := &model.User{
		Name:  u.Name,
		Email: u.Email,
	}

	user.ID = uuid.New()

	_, err := r.db.NewInsert().Model(user).Exec(r.ctx)
	if err != nil {
		return nil, err
	}
	userDTO := &model.UserDTO{
		ID:   user.ID,
		Name: user.Name,
	}
	return userDTO, nil
}

func (r *UserRepository) GetUserByID(id int64) (*model.UserDTO, error) {
	var user model.User
	err := r.db.NewSelect().Model(&user).Where("id = ?", id).Scan(r.ctx)
	if err != nil {
		return nil, err
	}

	userDTO := model.UserDTO{
		ID:   user.ID,
		Name: user.Name,
	}
	return &userDTO, nil
}

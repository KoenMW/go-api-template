package repository

import (
	"context"
	userModel "go-api/domain/model"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user *userModel.User) (*userModel.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*userModel.User, error)
	List(ctx context.Context, perPage int, page int) ([]userModel.User, error)
	Update(ctx context.Context, user *userModel.User) (*userModel.User, error)
	Delete(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
}

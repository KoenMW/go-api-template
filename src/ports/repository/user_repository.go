package repository

import (
	userModel "go-api/domain/model"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user *userModel.User) (*userModel.User, error)
	GetByID(id uuid.UUID) (*userModel.User, error)
	List(perPage int, page int) ([]userModel.User, error)
	Update(user *userModel.User) (*userModel.User, error)
	Delete(id uuid.UUID) (uuid.UUID, error)
}

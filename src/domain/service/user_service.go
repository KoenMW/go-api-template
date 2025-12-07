package service

import (
	"go-api/domain/model"
	"go-api/ports/repository"

	"github.com/google/uuid"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *model.CreateUserDTO) (*model.UserDTO, error) {
	newUser := &model.User{
		Name:  user.Name,
		Email: user.Email,
	}
	createdUser, err := s.repo.Create(newUser)
	return &model.UserDTO{
		ID:    createdUser.ID,
		Name:  createdUser.Name,
		Email: createdUser.Email,
	}, err
}

func (s *UserService) GetUserByID(id uuid.UUID) (*model.UserDTO, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &model.UserDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

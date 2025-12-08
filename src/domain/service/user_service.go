package service

import (
	"context"
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

func (s *UserService) CreateUser(ctx context.Context, user *model.CreateUserDTO) (*model.UserDTO, error) {
	newUser := &model.User{
		ID:    uuid.New(),
		Name:  user.Name,
		Email: user.Email,
	}
	createdUser, err := s.repo.Create(ctx, newUser)
	return &model.UserDTO{
		ID:    createdUser.ID,
		Name:  createdUser.Name,
		Email: createdUser.Email,
	}, err
}

func (s *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (*model.UserDTO, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &model.UserDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *UserService) ListUsers(ctx context.Context, perPage int, page int) ([]model.UserDTO, error) {
	users, err := s.repo.List(ctx, perPage, page)
	if err != nil {
		return nil, err
	}
	var userDTOs []model.UserDTO
	for _, user := range users {
		userDTOs = append(userDTOs, model.UserDTO{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return userDTOs, nil
}

func (s *UserService) UpdateUser(ctx context.Context, userDTO *model.UserDTO) (*model.UserDTO, error) {
	user := &model.User{
		ID:    userDTO.ID,
		Name:  userDTO.Name,
		Email: userDTO.Email,
	}
	updatedUser, err := s.repo.Update(ctx, user)
	if err != nil {
		return nil, err
	}
	return &model.UserDTO{
		ID:    updatedUser.ID,
		Name:  updatedUser.Name,
		Email: updatedUser.Email,
	}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	return s.repo.Delete(ctx, id)
}

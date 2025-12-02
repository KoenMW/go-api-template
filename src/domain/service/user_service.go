package service

import (
	userModel "go-api/domain/model"
	userRepository "go-api/ports/repository"
)

type UserService struct {
	repo userRepository.UserRepository
}

func NewUserService(repo userRepository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *userModel.CreateUserDTO) (*userModel.UserDTO, error) {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetUserByID(id int64) (*userModel.UserDTO, error) {
	return s.repo.GetUserByID(id)
}

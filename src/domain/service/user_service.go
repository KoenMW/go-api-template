package service

import (
	"go-api/domain/model"
	"go-api/ports/repository"
	"go-api/ports/service"
)

type UserService struct {
	service.BaseService[*model.User, *model.UserDTO, *model.CreateUserDTO, repository.UserRepository]
}

func NewUserService(repo repository.UserRepository) service.UserService {
	return &UserService{
		BaseService: NewBaseService[*model.User, *model.UserDTO, *model.CreateUserDTO](repo),
	}
}

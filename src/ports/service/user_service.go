package service

import (
	"go-api/domain/model"
	"go-api/ports/repository"
)

type UserService = interface {
	BaseService[*model.User, *model.UserDTO, *model.CreateUserDTO, repository.UserRepository]
}

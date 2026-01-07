package rest

import (
	"go-api/domain/model"
	"go-api/ports/handler"
	"go-api/ports/repository"
	"go-api/ports/service"
)

type UserHandler struct {
	Basehandler[*model.User, *model.UserDTO, *model.CreateUserDTO, repository.UserRepository, service.UserService]
	UserService service.UserService
}

func NewUserHandler(s service.UserService) handler.UserHandler {
	return &UserHandler{
		UserService: s,
	}
}

package repository

import userModel "go-api/domain/model"

type UserRepository interface {
	CreateUser(user *userModel.CreateUserDTO) (*userModel.UserDTO, error)
	GetUserByID(id int64) (*userModel.UserDTO, error)
}

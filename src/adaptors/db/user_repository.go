package db

import (
	"go-api/domain/model"
	dbPort "go-api/ports/db"
	"go-api/ports/repository"
)

type UserRepository struct {
	repository.BaseRepository[*model.User]
}

func NewUserRepository(db dbPort.DB) repository.UserRepository {

	return &UserRepository{BaseRepository: NewRepository[*model.User](db)}
}

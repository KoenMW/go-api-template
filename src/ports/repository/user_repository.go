package repository

import (
	"go-api/domain/model"
)

type UserRepository interface {
	BaseRepository[*model.User]
}

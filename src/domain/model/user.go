package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	_         struct{}  `bun:"table:users"`
	ID        uuid.UUID `bun:"id,pk,type:uuid"`
	Name      string    `bun:"name,notnull"`
	Email     string    `bun:"email,unique,notnull"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}

type UserDTO struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type CreateUserDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

package model

import "github.com/google/uuid"

type User struct {
	ID    uuid.UUID `bun:"id,pk,type:uuid"`
	Name  string    `bun:"name,notnull"`
	Email string    `bun:"email,unique,notnull"`
}

type UserDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type CreateUserDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

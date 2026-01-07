package model

import (
	"errors"
	"go-api/domain/core"
	"net/mail"
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

func (u *User) SetID(id uuid.UUID) {
	u.ID = id
}

func (u *User) GetId() uuid.UUID {
	return u.ID
}

func (u *User) SetCreatedAt() {
	u.CreatedAt = time.Now().UTC()
}

func (u *User) SetUpdatedAt() {
	u.UpdatedAt = time.Now().UTC()
}

type UserDTO struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

func (dto *UserDTO) ApplyToEntity(entity *User) {
	entity.Name = dto.Name
	entity.Email = dto.Email
}

func (dto *UserDTO) RecieveEntity(entity *User) {
	dto.ID = entity.ID
	dto.Name = entity.Name
	dto.Email = entity.Email
}

type CreateUserDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (dto *CreateUserDTO) Validate() error {
	if dto.Name == "" {
		return errors.New(core.EmptyName)
	}
	if dto.Name != "" && len(dto.Name) > 100 {
		return errors.New(core.NameTooLong)
	}
	if dto.Email == "" {
		return errors.New(core.EmptyEmail)
	}
	if _, err := mail.ParseAddress(dto.Email); err != nil {
		return errors.New(core.InvalidEmail)
	}

	return nil
}

func (dto *CreateUserDTO) ApplyToEntity(entity *User) {
	entity.Name = dto.Name
	entity.Email = dto.Email
}

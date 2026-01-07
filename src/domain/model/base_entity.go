package model

import (
	"reflect"

	"github.com/google/uuid"
)

type BaseEntity interface {
	SetID(id uuid.UUID)
	GetId() uuid.UUID
	SetCreatedAt()
	SetUpdatedAt()
}

func NewEntity[T any]() T {
	var zero T
	return reflect.New(reflect.TypeOf(zero).Elem()).Interface().(T)
}

type BaseDTO[E BaseEntity] interface {
	ApplyToEntity(entity E)
	RecieveEntity(entity E)
}

type BaseCreateDTO[E BaseEntity] interface {
	Validate() error
	ApplyToEntity(entity E)
}

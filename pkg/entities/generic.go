package entities

import (
	"context"
)

type Repository[T interface{}] interface {
	Create(ctx context.Context, entity *T) error
	GetAll(ctx context.Context) (*[]T, error)
	GetByID(ctx context.Context, id uint) (*T, error)
	Update(ctx context.Context, entity *T, id uint) error
	Delete(ctx context.Context, id uint) error
}

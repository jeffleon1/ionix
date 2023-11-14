package posgrest

import (
	"context"

	"github.com/jeffleon1/ionix/pkg/entities"
	"gorm.io/gorm"
)

type repository[T interface{}] struct {
	db *gorm.DB
}

func New[T interface{}](db *gorm.DB) entities.Repository[T] {
	return &repository[T]{
		db,
	}
}

func (r *repository[T]) Create(ctx context.Context, entity *T) error {
	return r.db.WithContext(ctx).Create(&entity).Error
}

func (r *repository[T]) GetAll(ctx context.Context) (*[]T, error) {
	var entities []T
	err := r.db.WithContext(ctx).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return &entities, nil
}

func (r *repository[T]) GetByID(ctx context.Context, id uint) (*T, error) {
	var entity T
	if err := r.db.WithContext(ctx).First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *repository[T]) Update(ctx context.Context, entity *T, id uint) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Updates(entity).Error
}

func (r *repository[T]) Delete(ctx context.Context, id uint) error {
	var entity T
	return r.db.WithContext(ctx).Delete(&entity, id).Error
}

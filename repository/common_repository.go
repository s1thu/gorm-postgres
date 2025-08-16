package repository

import "gorm.io/gorm"

type Repository[T any] interface {
	Create(entity *T) error
	GetByID(id uint) (*T, error)
	GetAll() ([]*T, error)
	Update(entity *T) error
	Delete(id uint) error
}

type BaseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) Repository[T] {
	return &BaseRepository[T]{db: db}
}

// Create implements Repository.
func (b *BaseRepository[T]) Create(entity *T) error {
	return b.db.Create(entity).Error
}

// Delete implements Repository.
func (b *BaseRepository[T]) Delete(id uint) error {
	var zero T
	return b.db.Delete(&zero, id).Error
}

// GetAll implements Repository.
func (b *BaseRepository[T]) GetAll() ([]*T, error) {
	var entities []*T
	err := b.db.Find(&entities).Error
	return entities, err
}

// GetByID implements Repository.
func (b *BaseRepository[T]) GetByID(id uint) (*T, error) {
	var entity T
	err := b.db.First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

// Update implements Repository.
func (b *BaseRepository[T]) Update(entity *T) error {
	return b.db.Save(entity).Error
}

package repository

import (
	"github.com/s1thu/gorm-postgres/model"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	Base BaseRepository[model.Author]
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{
		Base: BaseRepository[model.Author]{db: db},
	}
}

// Author-specific methods
func (r *AuthorRepository) GetByEmail(email string) (*model.Author, error) {
	var author model.Author
	err := r.Base.db.Where("email = ?", email).First(&author).Error
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *AuthorRepository) GetWithBooks(id uint) (*model.Author, error) {
	var author model.Author
	err := r.Base.db.Preload("Books").First(&author, id).Error
	if err != nil {
		return nil, err
	}
	return &author, nil
}

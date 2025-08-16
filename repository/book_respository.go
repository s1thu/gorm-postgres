package repository

import (
	"github.com/s1thu/gorm-postgres/model"
	"gorm.io/gorm"
)

type BookRepository struct {
	Base BaseRepository[model.Book]
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		Base: BaseRepository[model.Book]{db: db},
	}
}

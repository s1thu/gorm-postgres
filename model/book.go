package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	ISBN        string         `gorm:"uniqueIndex" json:"isbn"`
	Price       float64        `gorm:"not null" json:"price"`
	Description string         `json:"description"`
	AuthorID    uint           `gorm:"not null" json:"author_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	// Relationship: Each book belongs to one author
	Author Author `gorm:"foreignKey:AuthorID" json:"author"`
}

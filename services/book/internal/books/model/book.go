package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID            string         `json:"id"`
	Title         string         `json:"title"`
	Description   string         `json:"description"`
	ISBN          string         `json:"isbn"`
	AuthorID      string         `json:"author_id"`
	CategoryID    string         `json:"category_id"`
	PublishedYear int            `json:"published_year"`
	BookStock     BookStock      `json:"book_stock" gorm:"foreignKey:BookID"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-"`
}

func (Book) TableName() string {
	return "books"
}

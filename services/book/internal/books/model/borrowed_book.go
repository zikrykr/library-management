package model

import (
	"time"

	"gorm.io/gorm"
)

type BorrowedBook struct {
	ID         string    `json:"id"`
	BookID     string    `json:"book_id"`
	UserID     string    `json:"user_id"`
	BorrowedAt time.Time `json:"borrowed_at"`
	DueAt      time.Time `json:"due_at"`
	ReturnedAt time.Time `json:"returned_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  gorm.DeletedAt
}

func (BorrowedBook) TableName() string {
	return "borrowed_books"
}

package model

import (
	"time"

	"gorm.io/gorm"
)

type BookStock struct {
	ID             string    `json:"id"`
	BookID         string    `json:"book_id"`
	TotalStock     int       `json:"total_stock"`
	AvailableStock int       `json:"available_stock"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      gorm.DeletedAt
}

func (BookStock) TableName() string {
	return "book_stocks"
}

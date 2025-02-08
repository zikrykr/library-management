package model

import (
	"time"

	"gorm.io/gorm"
)

type BookStock struct {
	ID             string         `json:"-"`
	BookID         string         `json:"-"`
	TotalStock     int            `json:"total_stock"`
	AvailableStock int            `json:"available_stock"`
	CreatedAt      time.Time      `json:"-"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"-"`
}

func (BookStock) TableName() string {
	return "book_stocks"
}

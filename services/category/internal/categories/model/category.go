package model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          string         `json:"id" gorm:"column:id"`
	Name        string         `json:"name" gorm:"column:name"`
	Description string         `json:"description" gorm:"column:description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

package model

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	ID        string         `json:"id" gorm:"column:id"`
	Name      string         `json:"name" gorm:"column:name"`
	Bio       string         `json:"bio" gorm:"column:bio"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

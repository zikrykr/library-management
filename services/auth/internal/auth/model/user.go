package model

import (
	"time"
)

type (
	User struct {
		ID           string    `json:"id" gorm:"column:id"`
		FullName     string    `json:"full_name" gorm:"column:full_name"`
		Email        string    `json:"email" gorm:"column:email"`
		Role         string    `json:"role" gorm:"column:role"`
		PasswordHash string    `json:"password_hash" gorm:"column:password_hash"`
		CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
		UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at"`
	}
)

func (User) TableName() string {
	return "users"
}

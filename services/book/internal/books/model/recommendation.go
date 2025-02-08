package model

import "gorm.io/gorm"

type Recommendation struct {
	ID        string `json:"id"`
	BookID    string `json:"book_id"`
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt gorm.DeletedAt
}

func (Recommendation) TableName() string {
	return "recommendations"
}

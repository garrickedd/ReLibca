package entity

import (
	"time"

	_ "gorm.io/gorm"
)

type Book struct {
	BookId     string    `json:"book_id" gorm:"primaryKey"`
	Name       string    `json:"name" gorm:"not null"`
	Author     string    `json:"author" gorm:"not null"`
	Place      string    `json:"place" gorm:"not null"`
	BookImage  string    `json:"book_image" gorm:"not null"`
	Category   Category  `json:"category" gorm:"foreignKey:CategoryId;references:CategoryId;not null"`
	CategoryId string    `json:"category_id" gorm:"uniqueIndex"`
	CreatedAt  time.Time `json:"created_at" gorm:"not null"`
	ModifiedAt time.Time `json:"modified_at" gorm:"not null"`
}

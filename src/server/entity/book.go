package entity

import (
	"time"

	_ "gorm.io/gorm"
)

type Book struct {
	BookId     string    `json:"book_id" gorm:"primaryKey;type:varchar(20)"`
	Name       string    `json:"name" gorm:"type:text;not null"`
	Author     string    `json:"author" gorm:"type:text;not null"`
	Place      string    `json:"place" gorm:"type:text;not null"`
	BookImage  string    `json:"book_image" gorm:"type:text;not null"`
	CategoryId string    `json:"category_id" gorm:"type:varchar(20);not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"not null"`
	ModifiedAt time.Time `json:"modified_at" gorm:"not null"`
}

package entity

import (
	"time"

	_ "gorm.io/gorm"
)

type Book struct {
	BookId     string    `gorm:"primaryKey;type:varchar(20)"`
	Name       string    `gorm:"type:text;not null"`
	Author     string    `gorm:"type:text;not null"`
	Place      string    `gorm:"type:text;not null"`
	BookImage  string    `gorm:"type:text;not null"`
	CategoryId string    `gorm:"type:varchar(20);not null"`
	CreatedAt  time.Time `gorm:"not null"`
	ModifiedAt time.Time `gorm:"not null"`
}

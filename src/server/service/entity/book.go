package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type Book struct {
	gorm.Model
	BookId uuid.UUID `gorm:"primaryKey"`
	
}
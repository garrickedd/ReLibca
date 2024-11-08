package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookCategory struct {
	gorm.Model
	Id 		uuid.UUID
	Name	string		"gorm:text;not null"
	Type	int			
}
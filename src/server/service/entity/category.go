package entity

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryId	string	`json:"category_id" gorm:"primary_key;not null"`
	Name		string	`json:"name" gorm:"not null"`
}
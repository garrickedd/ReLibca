package entity

import _ "gorm.io/gorm"

type Category struct {
	CategoryId string `json:"category_id" gorm:"primary_key;not null"`
	Name       string `json:"name" gorm:"not null"`
}

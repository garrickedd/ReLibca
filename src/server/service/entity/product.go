package entity

import (
	"time"

	_ "gorm.io/gorm"
)

type Product struct {
	// gorm.Model
	ProductId    string    `json:"product_id" gorm:"primary_key;unique"`
	Name         string    `json:"name" gorm:"not null"`
	Description  string    `json:"description" gorm:"not null"`
	ProductImage string    `json:"product_image" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"not null"`
	ModifiedAt   time.Time `json:"modified_at" gorm:"not null"`
}

package entity

import (
	"time"

	_ "gorm.io/gorm"
)

type Product struct {
	ProductId    string    `json:"product_id" gorm:"primaryKey;unique;type:varchar(20)"`
	Name         string    `json:"name" gorm:"type:text;not null"`
	Description  string    `json:"description" gorm:"type:text;not null"`
	ProductImage string    `json:"product_image" gorm:"type:text;not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"not null"`
	ModifiedAt   time.Time `json:"modified_at" gorm:"not null"`
}

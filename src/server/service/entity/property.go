package entity

import (
	_ "gorm.io/gorm"
)

type Property struct {
	// gorm.Model
	Product   Product `json:"product" gorm:"foreignKey:ProductId;references:ProductId;not null"`
	ProductId string  `json:"product_id"`
	Size      string  `json:"size" gorm:"not null"`
	Unit      int     `json:"unit" gorm:"not null"`
	Price     float64 `json:"quantity" gorm:"type:decimal(19,4);not null"`
}

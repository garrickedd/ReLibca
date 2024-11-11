package entity

import (
	_ "gorm.io/gorm"
)

type Property struct {
	ProductId string  `json:"product_id" gorm:"primaryKey;type:varchar(20);not null"`
	Size      string  `json:"size" gorm:"type:varchar(10);not null"`
	Unit      int     `json:"unit" gorm:"type:text;not null"`
	Price     float64 `json:"price" gorm:"type:money;not null"`
}

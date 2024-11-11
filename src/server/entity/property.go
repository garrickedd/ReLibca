package entity

import (
	_ "gorm.io/gorm"
)

type Property struct {
	ProductId string  `gorm:"primaryKey;type:varchar(20);not null"`
	Size      string  `gorm:"type:varchar(10);not null"`
	Unit      int     `gorm:"type:text;not null"`
	Price     float64 `gorm:"type:money;not null"`
}

package entity

import (
	"gorm.io/gorm"
)

type Property struct {
	gorm.Model
	Product		Product		`json:"product_id" gorm:"primary_key;foreignKey:OrderId;not null"`
	ProductId	string
	Size		string		`json:"size" gorm:"not null"`
	Unit		int			`json:"unit" gorm:"not null"`
	Price		float64		`json:"quantity" gorm:"type:decimal(19,4);not null"`
}
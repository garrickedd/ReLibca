package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	Order		Order		`json:"order_id" gorm:"primary_key;foreignKey:OrderId;not null"`
	OrderId		uuid.UUID
	Product		Product		`json:"product_id" gorm:"foreignKey:ProductId;not null"`
	ProductId	string
	Book		Book		`json:"book_id" gorm:"foreignKey:BookId;not null"`
	BookId		string
	Quantity	int			`json:"quantity" gorm:"not null"`
}
package entity

import (
	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type OrderItem struct {
	Order     Order     `json:"order" gorm:"primary_key;foreignKey:OrderId;references:OrderId;not null"`
	OrderId   uuid.UUID `json:"order_id"`
	Product   Product   `json:"product" gorm:"foreignKey:ProductId;not null"`
	ProductId string    `json:"product_id"`
	Book      Book      `json:"book" gorm:"foreignKey:BookId;references:BookId;not null"`
	BookId    string    `json:"book_id"`
	Quantity  int       `json:"quantity" gorm:"not null"`
}

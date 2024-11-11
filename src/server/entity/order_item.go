package entity

import (
	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type OrderItem struct {
	OrderId   uuid.UUID `json:"order_id" gorm:"primaryKey"`
	ProductId string    `json:"product_id" gorm:"type:varchar(20);not null"`
	BookId    string    `json:"book_id" gorm:"type:varchar(20);not null"`
	Quantity  int       `json:"quantity" gorm:"not null"`
}

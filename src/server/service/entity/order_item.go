package entity

import (
	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type OrderItem struct {
	OrderId   uuid.UUID `gorm:"primaryKey"`
	ProductId string    `gorm:"type:varchar(20);not null"`
	BookId    string    `gorm:"type:varchar(20);not null"`
	Quantity  int       `gorm:"not null"`
}

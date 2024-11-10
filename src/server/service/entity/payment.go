package entity

import (
	"time"

	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type Payment struct {
	PaymentId uuid.UUID `json:"payment_id" gorm:"primary_key;unique"`
	Method    int       `json:"method" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	User      User      `json:"user" gorm:"foreignKey:UserId;references:UserId;not null"`
	UserId    string    `json:"user_id"`
	Order     Order     `json:"order" gorm:"foreignKey:OrderId;references:OrderId;not null"`
	OrderId   uuid.UUID `json:"order_id"`
}

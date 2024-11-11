package entity

import (
	"time"

	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type Payment struct {
	PaymentId  uuid.UUID `json:"payment_id" gorm:"primaryKey;unique;type:uuid"`
	Method     int       `json:"method" gorm:"not null"`
	TotalMoney float64   `json:"total_money" gorm:"type:money;not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"not null"`
	UserId     string    `json:"user_id" gorm:"type:varchar(20);not null"`
	OrderId    uuid.UUID `json:"order_id" gorm:"type:uuid;not null"`
}

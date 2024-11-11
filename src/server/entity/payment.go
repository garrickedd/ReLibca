package entity

import (
	"time"

	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type Payment struct {
	PaymentId  uuid.UUID `gorm:"primaryKey;unique;type:uuid"`
	Method     int       `gorm:"not null"`
	TotalMoney float64   `gorm:"type:money;not null"`
	CreatedAt  time.Time `gorm:"not null"`
	UserId     string    `gorm:"type:varchar(20);not null"`
	OrderId    uuid.UUID `gorm:"type:uuid;not null"`
}

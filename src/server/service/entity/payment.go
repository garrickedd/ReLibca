package entity

import (
	"time"

	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type Payment struct {
	// gorm.Model
	PaymentId uuid.UUID `json:"payment_id" gorm:"primary_key;unique"`
	Method    int       `json:"method" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	User      User      `json:"user_id" gorm:"foreignKey:UserId"`
	UserId    string
}

package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderId		uuid.UUID	`json:"order_id" gorm:"primary_key;unique"`
	Note		string		`json:"note" gorm:"not null"`
	Status		int			`json:"status" gorm:"not null"`
	CreatedAt	time.Time	`json:"created_at" gorm:"not null"`
	ModifiedAt	time.Time	`json:"modified_at" gorm:"not null"`
	Promotion	Promotion	`json:"promotion_id" gorm:"foreignKey:PromotionId;not null"`
	PromotionId	string
}
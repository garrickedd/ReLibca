package entity

import (
	"time"

	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type Order struct {
	OrderId     uuid.UUID `json:"order_id" gorm:"primaryKey;unique;type:uuid"`
	Note        string    `json:"note" gorm:"type:text;not null"`
	Status      int       `json:"status" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	ModifiedAt  time.Time `json:"modified_at" gorm:"not null"`
	PromotionId string    `json:"promotion_id" gorm:"type:varchar(20)"`
}

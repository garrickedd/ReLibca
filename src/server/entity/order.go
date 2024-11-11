package entity

import (
	"time"

	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type Order struct {
	OrderId     uuid.UUID `gorm:"primaryKey;unique;type:uuid"`
	Note        string    `gorm:"type:text;not null"`
	Status      int       `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
	ModifiedAt  time.Time `gorm:"not null"`
	PromotionId string    `gorm:"type:varchar(20)"`
}

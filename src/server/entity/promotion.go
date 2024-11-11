package entity

import (
	"time"

	_ "gorm.io/gorm"
)

type Promotion struct {
	PromotionId string    `gorm:"primaryKey;type:string"`
	Code        string    `gorm:"type:text;not null"`
	Name        string    `gorm:"type:text;not null"`
	Discount    float32   `gorm:"not null"`
	StartTime   time.Time `gorm:"not null"`
	EndTime     time.Time `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
	ModifiedAt  time.Time `gorm:"not null"`
}

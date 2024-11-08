package entity

import (
	"time"

	_ "gorm.io/gorm"
)

type Promotion struct {
	// gorm.Model
	PromotionId string    `json:"promotion_id" gorm:"primaryKey"`
	Code        string    `json:"code" gorm:"not null"`
	Discount    string    `json:"discount" gorm:"not null"`
	StartTime   time.Time `json:"start_time" gorm:"not null"`
	EndTime     time.Time `json:"end_time" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	ModifiedAt  time.Time `json:"modified_at" gorm:"not null"`
}

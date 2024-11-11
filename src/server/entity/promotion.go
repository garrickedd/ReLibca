package entity

import (
	"time"

	_ "gorm.io/gorm"
)

type Promotion struct {
	PromotionId string    `json:"promotion_id" gorm:"primaryKey;type:string"`
	Code        string    `json:"code" gorm:"type:text;not null"`
	Name        string    `json:"name" gorm:"type:text;not null"`
	Discount    float32   `json:"discount" gorm:"not null"`
	StartTime   time.Time `json:"start_time" gorm:"not null"`
	EndTime     time.Time `json:"end_time" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	ModifiedAt  time.Time `json:"modified_at" gorm:"not null"`
}

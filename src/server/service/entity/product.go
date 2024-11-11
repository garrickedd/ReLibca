package entity

import (
	"time"

	_ "gorm.io/gorm"
)

type Product struct {
	ProductId    string    `gorm:"primaryKey;unique;type:varchar(20)"`
	Name         string    `gorm:"type:text;not null"`
	Description  string    `gorm:"type:text;not null"`
	ProductImage string    `gorm:"type:text;not null"`
	CreatedAt    time.Time `gorm:"not null"`
	ModifiedAt   time.Time `gorm:"not null"`
}

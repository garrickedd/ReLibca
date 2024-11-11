package entity

import (
	"time"

	_ "gorm.io/gorm"
)

type User struct {
	UserId      string    `gorm:"primaryKey;unique;type:varchar(20)"`
	FirstName   string    `gorm:"type:text;not null"`
	LastName    string    `gorm:"type:text;not null"`
	PhoneNumber string    `gorm:"type:varchar(12);not null"`
	Password    string    `gorm:"type:text;not null"`
	Role        int       `gorm:"not null"`
	IsActive    bool      `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
	ModifiedAt  time.Time `gorm:"not null"`
}

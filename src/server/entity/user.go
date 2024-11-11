package entity

import (
	"time"

	_ "gorm.io/gorm"
)

type User struct {
	UserId      string    `json:"user_id" gorm:"primaryKey;unique;type:varchar(20)"`
	FirstName   string    `json:"first_name" gorm:"type:text;not null"`
	LastName    string    `json:"last_name" gorm:"type:text;not null"`
	PhoneNumber string    `json:"phone_number" gorm:"type:varchar(12);not null"`
	Password    string    `json:"password" gorm:"type:text;not null"`
	Role        int       `json:"role" gorm:"not null"`
	IsActive    bool      `json:"is_active" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	ModifiedAt  time.Time `json:"modified_at" gorm:"not null"`
}

package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserId		string		`json:"user_id" 		gorm:"primary_key;unique"`
	FirstName	string		`json:"first_name" 		gorm:"not null"`
	LastName	string		`json:"last_name" 		gorm:"not null"`
	PhoneNumber	string		`json:"phone_number"	gorm:"not null"`
	Password	string		`json:"password" 		gorm:"not null"`
	Role		int			`json:"role" 			gorm:"not null"`
	IsActive	bool		`json:"is_active" 		gorm:"not null"`
	CreatedAt	time.Time	`json:"role" 			gorm:"not null"`
	ModifiedAt	time.Time	`json:"role" 			gorm:"not null"`
}
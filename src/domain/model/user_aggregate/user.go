package model

import "github.com/garickedd/ReLibca/src/domain/model"

type User struct {
	model.BaseModel
	Username     string `gorm:"type:string;size:20;not null"`
	FirstName    string `gorm:"type:string;size:15"`
	LastName     string `gorm:"type:string;size:15"`
	MobileNumber string `gorm:"type:string;size:11;not null"`
	Email        string `gorm:"type:string;size:64;not null"`
	Password     string `gorm:"type:string;size:64;not null"`
	Enabled      bool   `gorm:"type:text;not null"`
}

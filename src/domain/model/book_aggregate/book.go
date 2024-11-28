package model

import "github.com/garickedd/ReLibca/src/domain/model"

type Book struct {
	model.BaseModel
	Title  string `gorm:"type:text;not null"`
	Author string `gorm:"type:text;null"`
	Place  string `gorm:"type:text;null"`
}

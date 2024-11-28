package model

import "github.com/garickedd/ReLibca/src/domain/model"

type Promotion struct {
	model.BaseModel
	Code         string  `gorm:"type:string;size:255;not null"`
	Name         string  `gorm:"type:text;not null"`
	Discount     float64 `gorm:"type:decimal(5,2);not null"`
	QuantityUnit int     `gorm:"type:int;default 0;not null"`
	StartTime    string  `gorm:"type:timestamp;not null"`
	EndTime      string  `gorm:"type:timestamp;not null"`
}

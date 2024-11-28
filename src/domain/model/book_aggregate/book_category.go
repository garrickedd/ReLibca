package model

import "github.com/garickedd/ReLibca/src/domain/model"

type BookCategory struct {
	model.BaseModel
	Tag string `gorm:"type:string;size:20;not null"`
}

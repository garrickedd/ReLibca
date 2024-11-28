package model

import "github.com/garickedd/ReLibca/src/domain/model"

type Role struct {
	model.BaseModel
	Name string `gorm:"type:string;size:10;not null"`
}

package model

import "github.com/garickedd/ReLibca/src/domain/model"

type ProductProperty struct {
	model.BaseModel
	Size      string `gorm:"type:string;size:10;not null"`
	Unit      int    `gorm:"type:int;not null"`
	UnitPrice string `gorm:"type:money;not null"`
}

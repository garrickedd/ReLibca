package model

import (
	"github.com/garickedd/ReLibca/src/domain/model"
	modelbook "github.com/garickedd/ReLibca/src/domain/model/book_aggregate"
	modelprd "github.com/garickedd/ReLibca/src/domain/model/product_aggregate"
)

type OrderItem struct {
	model.BaseModel
	ProductId int `gorm:"not null"` // FK -> Product
	Quantity  int `gorm:"type:int;not null"`
	BookId    int `gorm:"not null"` // FK -> Book

	// Relation
	Product modelprd.Product `gorm:"foreignKey:ProductId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	Book    modelbook.Book   `gorm:"foreignKey:BookId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
}

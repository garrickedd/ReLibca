package model

import (
	"github.com/garickedd/ReLibca/src/domain/model"
	modeluser "github.com/garickedd/ReLibca/src/domain/model/user_aggregate"
)

type Order struct {
	model.BaseModel
	UserId      int     `gorm:"not null"` // FK -> User
	OrderItemId int     `gorm:"not null"` // FK -> OrderItem
	PromotionId *int    `gorm:"not null"` // FK -> Book
	Total       float64 `gorm:"not null"`
	PaymentType int     `gorm:"not null"`

	// Relation
	User      modeluser.User `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	OrderItem OrderItem      `gorm:"foreignKey:OrderItemId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	Promotion Promotion      `gorm:"foreignKey:PromotionId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
}

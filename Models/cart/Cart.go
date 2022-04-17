package cart

import (
	"FinalProjectGO/Models/users"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	TotalPrice float64     `json:"total_price" gorm:"default:0"`
	Amount     int         `json:"amount" gorm:"default:0"`
	UserId     uint        `json:"user_id"`
	User       users.Users `gorm:"foreignkey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func NewCart(userId uint) *Cart {

	return &Cart{
		UserId: userId,
	}
}

package order

import (
	"FinalProjectGO/Models/users"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	TotalPrice float64     `json:"total_price"`
	Amount     int         `json:"amount"`
	UserId     uint        `json:"user_id"`
	User       users.Users `json:"user" gorm:"foreignkey:UserId"`
}

func NewOrder(totalPrice float64, amount int, userId uint) *Order {
	return &Order{
		TotalPrice: totalPrice,
		Amount:     amount,
		UserId:     userId,
	}
}

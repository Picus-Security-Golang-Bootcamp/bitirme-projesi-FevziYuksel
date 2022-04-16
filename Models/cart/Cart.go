package cart

import (
	Product "FinalProjectGO/Models/product"
	Users "FinalProjectGO/Models/users"
	"gorm.io/gorm"
	"os/user"
)

type Cart struct {
	gorm.Model
	TotalPrice float64   `json:"total_price" gorm:"default:0"`
	Amount     int       `json:"amount" gorm:"default:0"`
	UserId     uint      `json:"user_id"`
	User       user.User `gorm:"foreignkey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
type CartDetails struct {
	gorm.Model
	ProductName string          `json:"product_name"`
	Amount      int             `json:"amount"`
	UnitPrice   float64         `json:"unit_price"`
	TotalPrice  float64         `json:"total_price"`
	ProductId   uint            `json:"product_id"`
	CartId      uint            `json:"cart_id"`
	Product     Product.Product `gorm:"foreignkey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` //
	Cart        Cart            `gorm:"foreignkey:CartId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

//Garbage
type Cart34 struct {
	gorm.Model
	UserId        uint            `json:"user_id"`
	User          Users.Users     `gorm:"foreignkey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductId     uint            `json:"product_id"`
	Product       Product.Product `gorm:"foreignkey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductName   string          `json:"product_name"`
	ProductAmount uint            `json:"amount" gorm:"default:0"`
	//Amount      int             `json:"amount"`
	UnitPrice  float64 `json:"unit_price"`
	TotalPrice float64 `json:"total_price" gorm:"default:0"`
}

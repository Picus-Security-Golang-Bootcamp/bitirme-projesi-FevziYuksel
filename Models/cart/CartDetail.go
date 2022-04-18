package cart

import (
	Product "FinalProjectGO/Models/product"
	"gorm.io/gorm"
)

type CartDetails struct {
	gorm.Model
	ProductName string          `json:"product_name"`
	Amount      int             `json:"amount"`
	UnitPrice   float64         `json:"unit_price"`
	TotalPrice  float64         `json:"total_price"`
	ProductId   uint            `json:"product_id"`
	CartId      uint            `json:"cart_id"`
	Product     Product.Product `gorm:"foreignkey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` //

}

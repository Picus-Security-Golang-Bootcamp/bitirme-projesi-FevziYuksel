package product

import (
	Category "FinalProjectGO/Models/category"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName  string  `json:"product_name" validate:"required"`
	Price        float64 `json:"price" validate:"gt=1"`
	Stock        int     `json:"stock" validate:"gt=0"`
	CategoryId   uint
	CategoryName string            `json:"category_name" validate:"required"`
	SKU          string            `json:"sku" validate:"required"`
	Category     Category.Category `gorm:"foreignkey:CategoryId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func NewProduct(body Product) *Product {

	return &Product{
		ProductName:  body.ProductName,
		Price:        body.Price,
		Stock:        body.Stock,
		CategoryId:   body.CategoryId,
		CategoryName: body.CategoryName,
		SKU:          body.SKU,
	}
}

package order

import (
	"FinalProjectGO/Models/product"
	"database/sql"
	"gorm.io/gorm"
	"os/user"
	"time"
)

type Order struct {
	gorm.Model
	TotalPrice float64   `json:"total_price"`
	Amount     int       `json:"amount"`
	UserId     uint      `json:"user_id"`
	User       user.User `json:"user" gorm:"foreignkey:UserId"`
}
type OrderDetails struct {
	ID          uint            `gorm:"primary_key"`
	OrderId     uint            `json:"order_id"`
	ProductId   uint            `json:"product_id"`
	ProductName string          `json:"product_name"`
	Amount      int             `json:"amount"`
	UnitPrice   float64         `json:"unit_price"`
	TotalPrice  float64         `json:"total_price"`
	Order       Order           `gorm:"foreignkey:OrderId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Product     product.Product `gorm:"foreignkey:ProductId"`
}

//Garbage
type Order4rr struct {
	Id              uint `gorm:"primarykey"`
	NumberOfProduct uint
	TotalPrice      float64
	//OrderDetails uint
	//Product   *Product `gorm:"foreignKey:ProductId"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}

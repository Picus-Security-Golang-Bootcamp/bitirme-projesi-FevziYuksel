package Models

import (
	"database/sql"
	"errors"
	"time"
)

var (
	maxAllowedForBasket             = 20
	maxAllowedQtyPerProduct         = 9
	minCartAmountForOrder   float64 = 50

	ErrNotFound            = errors.New("Item not found")
	ErrCustomerCannotBeNil = errors.New("Customer cannot be nil")
)

type (
	//unique id ???
	//Details for each ordered products
	OrderDetail struct {
		Id                   uint `gorm:"primarykey"`
		ProductId            uint
		Product              *Product `gorm:"foreignKey:ProductId"`
		OrderId              uint
		NumberPerProduct     uint
		TotalPricePerProduct float64
		UnitPrice            float64
		CreatedAt            time.Time
		UpdatedAt            time.Time
		DeletedAt            sql.NullTime `gorm:"index"`
	}
)

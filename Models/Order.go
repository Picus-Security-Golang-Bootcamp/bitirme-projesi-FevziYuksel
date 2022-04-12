package Models

import (
	"database/sql"
	"time"
)

//Transactions
type Order struct {
	Id              uint `gorm:"primarykey"`
	NumberOfProduct uint
	TotalPrice      float64
	//OrderDetails uint
	//Product   *Product `gorm:"foreignKey:ProductId"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}

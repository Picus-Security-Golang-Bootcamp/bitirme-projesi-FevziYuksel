package Models

import (
	"database/sql"
	"time"
)

type Product struct {
	Id        uint `gorm:"primarykey"`
	Name      string
	Sku       string
	UnitPrice float64
	Quantity  uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}

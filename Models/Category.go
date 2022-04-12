package Models

import (
	"database/sql"
	"time"
)

type Category struct {
	Id        uint `gorm:"primarykey"`
	Name      string
	ProductId uint
	Product   *Product `gorm:"foreignKey:ProductId"` //şimdilik kaldırdım
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}

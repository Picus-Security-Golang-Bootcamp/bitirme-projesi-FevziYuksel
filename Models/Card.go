package Models

import (
	"database/sql"
	"time"
)

type Card struct { //Kullancı başı sepet ayır
	Id        uint `gorm:"primarykey"`
	ProductId uint
	Product   *Product `gorm:"foreignKey:ProductId"` //foreign key'i onla ilişkile
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}

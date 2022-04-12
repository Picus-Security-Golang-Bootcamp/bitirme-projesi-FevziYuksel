package Models

import (
	"database/sql"
	"time"
)

type User struct {
	Id        uint `gorm:"primarykey"`
	Email     string
	Password  string
	Role      string
	Card      *Card `gorm:"foreignKey:Id"` //1 basket per user ?
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}

//create special struct for users

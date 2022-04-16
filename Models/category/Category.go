package category

import (
	"gorm.io/gorm"
)

//İlişkilendirmelimiyim ?
type Category struct {
	gorm.Model
	//Id        uint `gorm:"primarykey"`
	Name string `json:"name"`
	//ProductId uint
	//Product   *product.Product `gorm:"foreignKey:ProductId"` //şimdilik kaldırdım
	//CreatedAt time.Time
	//UpdatedAt time.Time
	//DeletedAt sql.NullTime `gorm:"index"`
}

func NewCategory(name string) Category {
	return Category{
		Name: name,
	}
}

func (c *Category) GetName() string {
	return c.Name
}

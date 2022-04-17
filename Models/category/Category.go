package category

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name"`
}

func NewCategory(name string) Category {
	return Category{
		Name: name,
	}
}

func (c *Category) GetName() string {
	return c.Name
}

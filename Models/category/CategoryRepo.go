package category

import (
	config "FinalProjectGO/pkg/config"
	database "FinalProjectGO/pkg/database"
	"gorm.io/gorm"
	"log"
)

type CategoryRepository struct {
	db *gorm.DB
}

var (
	db   *gorm.DB
	repo *CategoryRepository
)

func init() {
	cfg, err := config.LoadConfig("./pkg/config/config-local")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	db = database.Connect(cfg)
	repo = NewCategoryRepository(db)
	repo.Migrations()
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}
func (c *CategoryRepository) Migrations() {
	err := c.db.AutoMigrate(&Category{})
	if err != nil {
		return
	}
}
func IsCategoryExist(name string) bool {
	var category Category
	db.Where("name = ?", name).Find(&category)

	if category.ID == 0 {
		return false
	}
	return true
}
func GetCategoryId(name string) uint {
	var category Category
	db.Where("name = ?", name).Find(&category)
	return category.ID
}

func CreateCategoryTable(insert interface{}) error {
	result := repo.db.Create(insert)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
func FindAllCategories() []Category {
	var category []Category
	db.Find(&category)
	return category
}

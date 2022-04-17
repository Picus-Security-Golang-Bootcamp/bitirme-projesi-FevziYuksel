package product

import (
	config "FinalProjectGO/pkg/config"
	database "FinalProjectGO/pkg/database"
	"gorm.io/gorm"
	"log"
)

type ProductRepository struct {
	db *gorm.DB
}

var (
	db   *gorm.DB
	repo *ProductRepository
)

func init() {
	cfg, err := config.LoadConfig("./pkg/config/config-local")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	db = database.Connect(cfg)
	repo = NewProductRepository(db)
	repo.Migrations()

}
func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Migrations() {
	err := r.db.AutoMigrate(&Product{})
	if err != nil {
		return
	}
}

func IsProductExist(productName string, sku string) bool {
	allProducts := FindAllProducts()

	if len(allProducts) != 0 {
		var product Product
		db.Where("product_name = ? OR sku = ?", productName, sku).Find(&product)
		if product.ID != 0 {
			return true
		}
	}
	return false
}

func FindAllProducts() []Product {
	var products []Product
	db.Find(&products)

	return products
}

func CreateProduct(product *Product) {
	db.Create(product)
}

func Update(product *Product) {
	db.Save(product)
}

func GetAllProducts(pageIndex, pageSize int) ([]Product, int) {
	var products []Product

	allProducts := FindAllProducts()
	db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products)

	return products, len(allProducts)
}

func SearchProduct(queryString string) []Product {
	var products []Product
	db.Where("product_name LIKE ?", "%"+queryString+"%").Or(db.Where("sku LIKE ?", "%"+queryString+"%")).Find(&products)

	return products
}

// SearchProductWithPagination searches products by product name and sku and returns []Product
func SearchProductWithPagination(queryString string, pageIndex, pageSize int) []Product {
	var products []Product
	db.Where("product_name LIKE ?", "%"+queryString+"%").Or(db.Where("sku LIKE ?", "%"+queryString+"%")).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products)

	return products
}

// SearchById searches products by product id and returns Product
func SearchById(id uint) *Product {
	var product Product
	db.Where("id = ?", id).Find(&product)

	return &product
}

// UpdateStock updates product stock
func UpdateStock(p Product, newStock int) {
	p.Stock = newStock
	db.Save(&p)
}

// UpdateName updates product name
func UpdateName(p Product, newProductName string) {
	p.ProductName = newProductName
	db.Save(&p)
}

// UpdatePrice updates product price
func UpdatePrice(p Product, newPrice float64) {
	p.Price = newPrice
	db.Save(&p)
}

// UpdateSKU updates product sku
func UpdateSKU(p Product, newSKU string) {
	p.SKU = newSKU
	db.Save(&p)
}

// DeleteProduct Deletes product from db
func DeleteProduct(p Product) {
	db.Delete(&p)
}

// SearchBySKU searches product by sku and returns Product
func SearchBySKU(sku string) *Product {
	var product Product
	db.Where("sku = ?", sku).Find(&product)

	return &product
}

// SearchByProductName searches products by product name and returns Product
func SearchByProductName(productName string) *Product {
	var product Product
	db.Where("product_name = ?", productName).Find(&product)

	return &product
}

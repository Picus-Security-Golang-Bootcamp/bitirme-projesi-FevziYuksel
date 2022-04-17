package order

import (
	"FinalProjectGO/pkg/config"
	database "FinalProjectGO/pkg/database"
	"gorm.io/gorm"
	"log"
)

type OrderRepository struct {
	db *gorm.DB
}

var (
	db        *gorm.DB
	orderRepo *OrderRepository
)

func init() {
	cfg, err := config.LoadConfig("./pkg/config/config-local")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	db = database.Connect(cfg)
	orderRepo = NewOrderRepository(db)
	orderRepo.Migration()
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (o *OrderRepository) Migration() {
	err := db.AutoMigrate(&Order{})
	if err != nil {
		panic(err)
	}
}

func CreateOrderTable(order *Order) {
	db.Create(order)
}

func SearchById(id uint, userId uint) *Order {
	var model Order
	db.Where("id = ? AND user_id = ?", id, userId).Find(&model)

	return &model
}

func DeleteOrder(id uint) {
	var model Order
	db.Where("id = ?", id).Delete(&model)
}

func FindUserOrders(userId uint) []Order {
	var orders []Order
	db.Where("user_id = ?", userId).Find(&orders)

	return orders
}

package order

import (
	"FinalProjectGO/pkg/config"
	database "FinalProjectGO/pkg/database"
	"gorm.io/gorm"
	"log"
)

type OrderDetailRepository struct {
	db *gorm.DB
}

var (
	orderDetailRepo *OrderDetailRepository
)

func init() {
	cfg, err := config.LoadConfig("./pkg/config/config-local")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	db = database.Connect(cfg)
	orderDetailRepo = NewOrderDetailRepository(db)
	orderDetailRepo.Migration()
}

func NewOrderDetailRepository(db *gorm.DB) *OrderDetailRepository {
	return &OrderDetailRepository{
		db: db,
	}
}

func (d *OrderDetailRepository) Migration() {
	err := db.AutoMigrate(&OrderDetails{})
	if err != nil {
		panic(err)
	}
}

func CreateOrderDetailTable(orderId uint, productId uint, amount int, unitPrice float64, totalPrice float64, productName string) {
	table := &OrderDetails{
		OrderId:     orderId,
		ProductId:   productId,
		Amount:      amount,
		UnitPrice:   unitPrice,
		TotalPrice:  totalPrice,
		ProductName: productName,
	}
	db.Create(table)
}

func FindOrderDetails(orderId uint) []OrderDetails {
	var orderDetails []OrderDetails
	db.Where("order_id = ?", orderId).Find(&orderDetails)

	return orderDetails
}

func DeleteModel(orderDetail OrderDetails) {
	db.Delete(&orderDetail)
}

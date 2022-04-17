package cart

import (
	"FinalProjectGO/pkg/config"
	database "FinalProjectGO/pkg/database"
	"gorm.io/gorm"
	"log"
)

type CartDetailsRepository struct {
	db *gorm.DB
}

var (
	cardDetailsRepo *CartDetailsRepository
)

func init() {
	cfg, err := config.LoadConfig("./pkg/config/config-local")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	db = database.Connect(cfg)
	cardDetailsRepo = NewUserRepository(db)
	cardDetailsRepo.Migrations()
}

func NewUserRepository(db *gorm.DB) *CartDetailsRepository {
	return &CartDetailsRepository{db: db}
}
func (d *CartDetailsRepository) Migrations() {
	err := db.AutoMigrate(&CartDetails{})
	if err != nil {
		return
	}
}

func CreateCartDetails(model *CartDetails) {
	cardDetailsRepo.db.Create(model)
}

func IsProductExist(userId uint, productID uint) bool {
	var cartDetails CartDetails
	cardDetailsRepo.db.Where("cart_id = ? AND product_id = ?", userId, productID).Find(&cartDetails)

	if cartDetails.ID == 0 {
		return false
	}
	return true
}

func DeleteProductInCart(cartID uint, productID uint) {
	var cartDetails CartDetails
	cardDetailsRepo.db.Where("cart_id = ? AND product_id = ?", cartID, productID).Find(&cartDetails)

	cardDetailsRepo.db.Delete(&cartDetails)
}

func GetAllCartDetailsOfUser(cartID uint) *[]CartDetails {
	var cartDetails []CartDetails
	cardDetailsRepo.db.Where("cart_id = ?", cartID).Find(&cartDetails)

	return &cartDetails
}

func GetCartDetailsByCartIdAndProductId(cartID uint, productID uint) *CartDetails {
	var cartDetails CartDetails
	cardDetailsRepo.db.Where("cart_id = ? AND product_id = ?", cartID, productID).Find(&cartDetails)

	return &cartDetails
}

func UpdateProductInCart(cartID uint, productID uint, amount int, totalPrice float64) {
	var cartDetails CartDetails
	cardDetailsRepo.db.Where("cart_id = ? AND product_id = ?", cartID, productID).Find(&cartDetails)

	cardDetailsRepo.db.Model(&cartDetails).Update("amount", amount)
	cardDetailsRepo.db.Model(&cartDetails).Update("total_price", totalPrice)
}

func GetCartDetailsByProductId(productID uint) *[]CartDetails {
	var cartDetails []CartDetails
	cardDetailsRepo.db.Where("product_id = ?", productID).Find(&cartDetails)

	return &cartDetails
}

func UpdateModel(model *CartDetails) {
	cardDetailsRepo.db.Save(model)
}

func DeleteModel(model *CartDetails) {
	cardDetailsRepo.db.Delete(model)
}

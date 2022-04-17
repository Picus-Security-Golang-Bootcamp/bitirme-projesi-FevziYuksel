package cart

import (
	config "FinalProjectGO/pkg/config"
	database "FinalProjectGO/pkg/database"
	"gorm.io/gorm"
	"log"
)

type CartRepository struct {
	db *gorm.DB
}

var (
	db   *gorm.DB
	repo *CartRepository
)

func init() {
	cfg, err := config.LoadConfig("./pkg/config/config-local")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	db = database.Connect(cfg)
	repo = NewCartRepository(db)
	repo.Migrations()
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}
func (d *CartRepository) Migrations() {

	err := d.db.AutoMigrate(&Cart{})
	if err != nil {
		return
	}
}

func (d *CartRepository) InsertCart(Card1 Cart) {
	d.db.FirstOrCreate(&Card1)
}

func CreateCardTable(cart *Cart) {
	db.Create(cart)
}

func SearchById(id uint) *Cart {
	var cart Cart
	db.Where("id = ?", id).First(&cart)
	return &cart
}

func Update(cart *Cart) {
	db.Save(cart)
}

func UpdateUserCart(userId uint, newAmount int, newPrice float64) {
	usersCart := SearchById(userId)
	usersCart.Amount += newAmount
	usersCart.TotalPrice += newPrice
	Update(usersCart)
}

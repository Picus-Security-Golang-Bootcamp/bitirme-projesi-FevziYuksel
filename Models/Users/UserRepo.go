package users

import (
	config "FinalProjectGO/pkg/config"
	database "FinalProjectGO/pkg/database"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type UserRepository struct {
	db *gorm.DB
}

var (
	db   *gorm.DB
	repo *UserRepository
)

func init() {
	cfg, err := config.LoadConfig("./pkg/config/config-local")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	db = database.Connect(cfg)
	repo = NewUserRepository(db)
	repo.Migrations()
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
func (r *UserRepository) Migrations() {
	err := r.db.AutoMigrate(&Users{})
	if err != nil {
		return
	}
}

func CreateUser(insert interface{}) error {
	result := repo.db.Create(insert)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func IsUserExist(email string) bool {
	flag := true
	var user Users
	repo.db.Where("email = ?", email).Find(&user)

	if user.ID == 0 {
		flag = false
	}
	return flag
}
func SearchByEmail(body Users) []Users {
	var users []Users
	db.Find(&users)
	email := body.GetEmail()
	fmt.Println(email)
	db.Where("email LIKE ? ", email).Find(&users)
	return users
}

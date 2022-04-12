package Models

import (
	"errors"
	"gorm.io/gorm"
)

type newDB struct {
	db *gorm.DB
}

//newDB constructor
func NewDBCreate(db *gorm.DB) *newDB {
	return &newDB{db: db}
}
func (d *newDB) Migrations() {
	err := d.db.AutoMigrate(&User{})
	if err != nil {
		return
	}
	err = d.db.AutoMigrate(&Product{})
	if err != nil {
		return
	}
	err = d.db.AutoMigrate(&Card{})
	if err != nil {
		return
	}
	err = d.db.AutoMigrate(&Category{})
	if err != nil {
		return
	}
}
func (d *newDB) Create(insert interface{}) error { //Bunun işlevi ne ???
	result := d.db.Create(insert)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

//-----------------

//Overloading yok go da

func (d *newDB) InsertUser(u User) {
	d.db.Where(User{Id: u.Id}).FirstOrCreate(&u)
}
func (d *newDB) InsertCard(card Card) {
	d.db.Where(Card{Id: card.Id}).FirstOrCreate(&card)
}
func (d *newDB) InsertOrder(o Order) {
	d.db.Where(Order{Id: o.Id}).FirstOrCreate(&o)
}
func (d *newDB) InsertProduct(p Product) {
	d.db.Where(Product{Id: p.Id}).FirstOrCreate(&p)
}
func (d *newDB) InsertCategory(c Category) {
	d.db.Where(Category{Id: c.Id}).FirstOrCreate(&c)
}

//----------

//Product DB Queries
func (d *newDB) ListProducts() []Product {
	var products []Product
	d.db.Find(&products)
	return products
}
func (d *newDB) GetProductByID(id int) (*Product, error) {
	var product Product
	result := d.db.First(&product, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &product, nil
}

func (d *newDB) FindProductByName(name string) []Product {
	var products []Product
	//name = strings.ToLower(name)
	d.db.Where("name LIKE ? ", "%"+name+"%").Find(&products)

	return products
}

func (d *newDB) DeleteProduct(product Product) error {
	result := d.db.Delete(product)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *newDB) ListCategory() []Category {
	var category []Category
	d.db.Find(&category)
	return category
}

package CSV

import (
	"FinalProjectGO/Models/category"
	Product "FinalProjectGO/Models/product"
	User "FinalProjectGO/Models/users"
	"encoding/csv"
	"os"
	"strconv"
)

func ReadCSV(filename string) ([][]string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err //httpErrors.OpenFileError ??
	}
	return records, nil
}

func CSVtoCategory(filename string) ([]category.Category, error) {
	records, err := ReadCSV(filename)
	if err != nil {
		return nil, err
	}

	var newCategory []category.Category

	for _, line := range records[1:] {
		newCategory = append(newCategory, category.Category{
			Name: line[0],
		})
	}
	return newCategory, nil
}

func CSVtoProduct(filename string) error {
	records, err := ReadCSV(filename)
	if err != nil {
		return err
	}
	var newProduct []*Product.Product
	for _, line := range records[1:] {
		price, _ := strconv.ParseFloat(line[1], 10)
		stock, _ := strconv.ParseInt(line[2], 10, 32)
		categoryId, _ := strconv.ParseUint(line[3], 10, 32)
		newProduct = append(newProduct, &Product.Product{
			ProductName:  line[0],
			Price:        price,
			Stock:        int(stock),
			CategoryId:   uint(categoryId),
			CategoryName: line[4],
			SKU:          line[5],
		})
	}
	for _, product := range newProduct {
		Product.InitializeProduct(product)
	}
	return nil
}
func CSVtoUser(filename string) error {
	records, err := ReadCSV(filename)
	if err != nil {
		return err
	}
	var newUsers []*User.Users
	for _, line := range records[1:] {
		newUsers = append(newUsers, &User.Users{
			Email:    line[0],
			Password: line[1],
			Role:     line[2],
		})
	}
	for _, newUser := range newUsers {
		err = User.CreateUser(newUser)
		if err != nil {
			return err
		}
	}
	return nil
}

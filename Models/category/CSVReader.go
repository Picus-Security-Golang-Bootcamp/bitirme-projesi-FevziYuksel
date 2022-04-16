package category

import (
	"encoding/csv"
	"fmt"

	"os"
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

/*
func CSVtoProduct(filename string) ([]Product.Product, error) {
	records, err := ReadCSV(filename)
	if err != nil {
		return nil, err
	}
	var newProduct []Product.Product
	for _, line := range records[1:] {
		newId, _ := strconv.ParseUint(line[0], 10, 32)
		Quantity, _ := strconv.ParseUint(line[3], 10, 32)
		newUnitPrice, _ := strconv.ParseFloat(line[4], 10)
		newProduct = append(newProduct, Product.Product{
			Id:        uint(newId),
			Name:      line[1],
			Sku:       line[2],
			UnitPrice: newUnitPrice,
			Quantity:  uint(Quantity),
		})
	}
	return newProduct, nil
}

*/
func CSVtoCategory(filename string) ([]Category, error) {
	records, err := ReadCSV(filename)
	if err != nil {
		return nil, err
	}

	var newCategory []Category
	fmt.Println(newCategory)
	for _, line := range records[1:] {
		newCategory = append(newCategory, Category{
			Name: line[0],
		})
	}
	return newCategory, nil
}

//Sil
func CSVtoCategory2(filename string) ([]Category, error) {
	records, err := ReadCSV(filename)
	if err != nil {
		return nil, err
	}

	var newCategory []Category
	for _, line := range records[1:] {
		//newId, _ := strconv.ParseUint(line[0], 10, 32)
		//ProductId, _ := strconv.ParseUint(line[2], 10, 32)
		newCategory = append(newCategory, Category{
			//Id:        uint(newId),
			Name: line[1],
			//ProductId: uint(ProductId),
		})
	}
	return newCategory, nil
}

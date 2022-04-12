package Models

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func ReadCSV(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err //httpErrors.OpenFileError ??
	}
	return records, nil
}

func CSVtoProduct(filename string) ([]Product, error) {
	records, err := ReadCSV(filename)
	if err != nil {
		return nil, err
	}
	var newProduct []Product
	for _, line := range records[1:] {
		newId, _ := strconv.ParseUint(line[0], 10, 32)
		Quantity, _ := strconv.ParseUint(line[3], 10, 32)
		newUnitPrice, _ := strconv.ParseFloat(line[4], 10)
		newProduct = append(newProduct, Product{
			Id:        uint(newId),
			Name:      line[1],
			Sku:       line[2],
			UnitPrice: newUnitPrice,
			Quantity:  uint(Quantity),
		})
	}
	return newProduct, nil
}
func CSVtoCategory(filename string) ([]Category, error) {
	records, err := ReadCSV(filename)
	if err != nil {
		return nil, err
	}
	fmt.Println(records)
	var newCategory []Category
	for _, line := range records[1:] {
		newId, _ := strconv.ParseUint(line[0], 10, 32)
		ProductId, _ := strconv.ParseUint(line[2], 10, 32)
		newCategory = append(newCategory, Category{
			Id:        uint(newId),
			Name:      line[1],
			ProductId: uint(ProductId),
		})
	}
	return newCategory, nil
}

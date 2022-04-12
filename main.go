package main

import (
	"FinalProjectGO/Database"
	"FinalProjectGO/Models"
	"fmt"
)

func main() {
	//Update query if items are already exist
	product1 := Models.Product{
		Id:        7,
		Name:      "pro1",
		Sku:       "w3",
		UnitPrice: 15,
		Quantity:  3,
	}
	product2 := Models.Product{
		Id:        8,
		Name:      "pro2",
		Sku:       "w3",
		UnitPrice: 15,
		Quantity:  3,
	}

	n := []Models.Product{product1, product2}

	basket1 := Models.Card{
		Id:        1,
		ProductId: product1.Id,
		Product:   &product1,
	}
	var user1 = Models.User{
		Id:       1,
		Email:    "mail1",
		Password: "123",
		Role:     "admin",
		Card:     &basket1,
	}

	db := Database.InitialMigration()
	userDB1 := *Models.NewDBCreate(db) //Repoların orda
	userDB1.Migrations()

	userDB1.InsertUser(user1)
	userDB1.InsertCard(basket1)

	for _, pro := range n {
		userDB1.InsertProduct(pro)
	}
	list, _ := Models.CSVtoProduct("ProductCSV.csv")

	for _, pro := range list {
		userDB1.InsertProduct(pro)
	}
	//for _, pro := range userDB1.ListProducts() {fmt.Println(pro)}
	/*
		for _, pro := range userDB1.FindProductByName("2") {
			fmt.Println(pro.Name)
		}
		aw, _ := userDB1.GetProductByID(8)
		fmt.Println(aw.Name)

		err := userDB1.DeleteProduct(*aw)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("\n %v", aw.Name)
	*/

	list2, _ := Models.CSVtoCategory("CategoryCSV.csv")
	fmt.Println(list2)
	for _, cat := range list2 {
		fmt.Println(cat)
		fmt.Println()
		//_ = userDB1.Create(cat)
		userDB1.InsertCategory(cat)
	}
	//for _, cat := range userDB1.ListCategory() {fmt.Println(cat.Name)}

}

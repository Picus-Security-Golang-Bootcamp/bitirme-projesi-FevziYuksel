package main

import (
	"FinalProjectGO/API"
	_ "FinalProjectGO/docs"
)

// @title Picus final project
// @description Small shopping card  application.

// @contact.name Fevzi Yüksel

// @contact.email fevziyuksel1996@gmail.com

// @host localhost:5000
// @BasePath /
// @schemes http
func main() {

	API.ServerSetup()

}

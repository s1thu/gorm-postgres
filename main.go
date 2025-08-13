package main

import (
	"log"

	"github.com/s1thu/gorm-postgres/config"
)

func main() {
	// Initialize the database connection
	config.ConnectDB()

	// Additional application logic can go here
	log.Println("Application started")
}

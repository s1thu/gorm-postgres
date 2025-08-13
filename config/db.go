package config

import (
	"fmt"
	"log"
	"os"

	"github.com/s1thu/gorm-postgres/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	fmt.Print("host: ", host, ", port: ", port, ", user: ", user, ", dbname: ", dbname, "\n")
	// Validate required variables
	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Fatal("Database environment variables are not fully set. Please export DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME.")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	err = DB.AutoMigrate(&model.Author{}, &model.Book{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
	log.Println("Database connection established")
}

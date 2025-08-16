package main

import (
	"log"

	"github.com/s1thu/gorm-postgres/config"
	"github.com/s1thu/gorm-postgres/repository"
	"github.com/s1thu/gorm-postgres/service"
)

func main() {
	// Initialize the database connection
	config.ConnectDB()

	authorRepo := repository.NewAuthorRepository(config.DB)
	authorService := service.NewAuthorService(authorRepo)

	// Example usage of the authorService
	newAuthor, err := authorService.CreateAuthor("John Doe", "john@example.com", "Author bio")
	if err != nil {
		log.Fatalf("Error creating author: %v", err)
	}
	log.Printf("Created author: %v", newAuthor)

	log.Println("Application completed")
}

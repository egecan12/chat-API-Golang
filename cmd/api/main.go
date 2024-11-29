package main

import (
	"log"
	"os"
	"path/filepath"

	"go-chat-app/pkg/db"
	"go-chat-app/pkg/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Log the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	log.Printf("Current working directory: %s", cwd)

	// Check if .env file exists
	envPath := filepath.Join(cwd, ".env")
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		log.Fatalf(".env file does not exist at path: %s", envPath)
	}

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Connect to the database
	db.Connect()
	defer db.DB.Close()

	// Setup Gin router
	r := gin.Default()
	routes.SetupRoutes(r)

	// Start server
	r.Run(":8080")
}

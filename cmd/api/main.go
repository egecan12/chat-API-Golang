package main

import (
	"go-chat-app/pkg/db"
	"go-chat-app/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	db.Connect()
	defer db.DB.Close()

	// Setup Gin router
	r := gin.Default()
	routes.SetupRoutes(r)

	// Start server
	r.Run(":8080")
}

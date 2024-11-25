package routes

import (
	"go-chat-app/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Serve the main chat page
	router.GET("/", handlers.AuthMiddleware(), func(c *gin.Context) {
		c.File("index.html") // Ensure the path is correct
	})

	// Serve the login page
	router.GET("/login", func(c *gin.Context) {
		c.File("login.html") // Ensure the path is correct
	})

	// Serve the register page
	router.GET("/register", func(c *gin.Context) {
		c.File("register.html") // Ensure the path is correct
	})

	// WebSocket endpoint for chat
	router.GET("/ws", func(c *gin.Context) {
		go handlers.HandleMessages() // Background goroutine to handle messages
		handlers.HandleConnections(c)
	})

	// Register route
	router.POST("/register", handlers.RegisterUser)

	// Login route
	router.POST("/login", handlers.LoginUser)

	// Serve the lobby page
	router.GET("/lobby", func(c *gin.Context) {
		c.File("lobby.html") // Ensure the path is correct
	})

	// Room endpoints
	router.GET("/rooms", handlers.GetRooms)    // List all rooms
	router.POST("/rooms", handlers.CreateRoom) // Create a new room
}

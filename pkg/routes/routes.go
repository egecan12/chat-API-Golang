package routes

import (
	"go-chat-app/pkg/handlers"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	htmlBasePath := os.Getenv("HTML_BASE_PATH")
	if htmlBasePath == "" {
		htmlBasePath = "." // Default to current directory if not set
	}

	// Serve the main chat page
	router.GET("/", handlers.AuthMiddleware(), func(c *gin.Context) {
		c.File(filepath.Join(htmlBasePath, "index.html"))
	})

	// Serve the login page
	router.GET("/login", func(c *gin.Context) {
		c.File(filepath.Join(htmlBasePath, "login.html"))
	})

	// Serve the register page
	router.GET("/register", func(c *gin.Context) {
		c.File(filepath.Join(htmlBasePath, "register.html"))
	})

	// WebSocket endpoint for chat
	router.GET("/ws", func(c *gin.Context) {
		go handlers.HandleMessages() // Background goroutine to handle messages
		handlers.HandleConnections(c)
	})

	// router.GET("/messages/:roomId", handlers.GetMessages)

	// Serve the room page
	router.GET("/messages", func(c *gin.Context) {
		roomId := c.Query("roomId")
		if roomId == "" {
			c.String(http.StatusBadRequest, "Room ID is required")
			return
		}
		c.File(filepath.Join(htmlBasePath, "room.html"))
	}, handlers.GetMessages)

	// Register route
	router.POST("/register", handlers.RegisterUser)

	// Login route
	router.POST("/login", handlers.LoginUser)

	// Serve the lobby page
	router.GET("/lobby", func(c *gin.Context) {
		c.File(filepath.Join(htmlBasePath, "lobby.html"))
	})

	// Serve the room page
	router.GET("/room", func(c *gin.Context) {
		roomId := c.Query("roomId")
		if roomId == "" {
			c.String(http.StatusBadRequest, "Room ID is required")
			return
		}
		c.File(filepath.Join(htmlBasePath, "room.html"))
	})

	// Room endpoints
	router.GET("/rooms", handlers.AuthMiddleware(), handlers.GetRooms)    // List all rooms
	router.POST("/rooms", handlers.AuthMiddleware(), handlers.CreateRoom) // Create a new room
}

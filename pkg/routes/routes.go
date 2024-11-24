package routes

import (
	"go-chat-app/pkg/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/register", handlers.RegisterUser)
	router.GET("/ws", func(c *gin.Context) {
		go handlers.HandleMessages()
		handlers.HandleConnections(c)
	})
	router.POST("/login", handlers.LoginUser)
	router.GET("/protected", handlers.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the protected route!"})
	})

}

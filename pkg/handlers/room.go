package handlers

import (
	"context"
	"go-chat-app/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateRoom API
func CreateRoom(c *gin.Context) {
	var body struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Token'dan kullanıcı kimliğini al (Auth Middleware ile)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	_, err := db.DB.Exec(context.Background(),
		"INSERT INTO rooms (name, created_by) VALUES ($1, $2)",
		body.Name, userID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create room"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Room created successfully!"})
}

// GetRooms API
func GetRooms(c *gin.Context) {
	rows, err := db.DB.Query(context.Background(), "SELECT id, name FROM rooms")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch rooms"})
		return
	}
	defer rows.Close()

	var rooms []map[string]interface{}
	for rows.Next() {
		var room struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}
		if err := rows.Scan(&room.ID, &room.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning room data"})
			return
		}
		rooms = append(rooms, gin.H{"id": room.ID, "name": room.Name})
	}

	c.JSON(http.StatusOK, gin.H{"rooms": rooms})
}

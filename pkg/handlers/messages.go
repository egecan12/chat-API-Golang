// FILE: handlers/messages.go

package handlers

import (
	"context"
	"go-chat-app/pkg/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMessages(c *gin.Context) {
	roomIdStr := c.Query("roomId")
	if roomIdStr == "" {
		http.Error(c.Writer, "Room ID is required", http.StatusBadRequest)
		return
	}

	roomId, err := strconv.Atoi(roomIdStr)
	if err != nil {
		http.Error(c.Writer, "Invalid Room ID", http.StatusBadRequest)
		return
	}

	rows, err := db.DB.Query(context.Background(),
		`SELECT u.username, m.content 
         FROM messages m
         JOIN users u ON m.user_id = u.id
         WHERE m.room_id=$1 
         ORDER BY m.created_at ASC`,
		roomId,
	)
	if err != nil {
		http.Error(c.Writer, "Error fetching messages", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		if err := rows.Scan(&msg.Username, &msg.Content); err != nil {
			http.Error(c.Writer, "Error scanning message", http.StatusInternalServerError)
			return
		}
		messages = append(messages, msg)
	}

	c.JSON(http.StatusOK, messages)
}

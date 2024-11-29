package handlers

import (
	"context"
	"go-chat-app/pkg/db"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]int) // Connected clients with room IDs
var broadcast = make(chan Message)          // Broadcast channel

type Message struct {
	Username string `json:"username"`
	Content  string `json:"content"`
	RoomID   int    `json:"roomID"`
	UserID   int    `json:"userID"`
}

func HandleConnections(c *gin.Context) {
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

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatalf("Error upgrading connection: %v\n", err)
	}
	defer ws.Close()

	clients[ws] = roomId

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading message: %v\n", err)
			delete(clients, ws)
			break
		}

		msg.RoomID = roomId
		broadcast <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast
		for client, roomId := range clients {
			if roomId == msg.RoomID {
				err := client.WriteJSON(msg)
				if err != nil {
					log.Printf("Error broadcasting message: %v\n", err)
					client.Close()
					delete(clients, client)
				}
			}
		}
		log.Printf("Message received: %+v\n", msg)

		var userExists bool
		err := db.DB.QueryRow(context.Background(),
			"SELECT EXISTS(SELECT 1 FROM users WHERE id=$1)",
			msg.UserID,
		).Scan(&userExists)
		if err != nil {
			log.Printf("Error checking user existence: %v\n", err)
			continue
		}

		if !userExists {
			log.Printf("User ID %d does not exist\n", msg.UserID)
			continue
		}

		_, err = db.DB.Exec(context.Background(),
			"INSERT INTO messages (user_id, room_id, content, created_at) VALUES ($1, $2, $3, now())",
			msg.UserID, msg.RoomID, msg.Content,
		)
		if err != nil {
			log.Printf("Error saving message to database: %v\n", err)
		}
	}
}

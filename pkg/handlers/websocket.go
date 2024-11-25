package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool) // Connected clients
var broadcast = make(chan Message)           // Broadcast channel

type Message struct {
	Username string `json:"username"`
	Content  string `json:"content"` // Ensure this matches "content" from the client
}

func HandleConnections(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatalf("Error upgrading connection: %v\n", err)
	}
	defer ws.Close()

	clients[ws] = true

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading message: %v\n", err)
			delete(clients, ws)
			break
		}

		log.Printf("Received message: %+v\n", msg) // Log the incoming message
		broadcast <- msg
	}

}

func HandleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg) // Broadcast as JSON
			if err != nil {
				log.Printf("Error broadcasting message: %v\n", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

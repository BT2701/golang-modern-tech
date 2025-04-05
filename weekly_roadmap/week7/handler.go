package week7

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	service   *MessageService
	clients   map[*websocket.Conn]bool
	broadcast chan Message
	upgrader  websocket.Upgrader
}

func NewWebSocketHandler(service *MessageService) *WebSocketHandler {
	return &WebSocketHandler{
		service:   service,
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan Message),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (h *WebSocketHandler) HandleConnections(c *gin.Context) {
	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}
	defer conn.Close()

	h.clients[conn] = true

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			delete(h.clients, conn)
			break
		}

		msg.Timestamp = time.Now()
		if err := h.service.SaveMessage(&msg); err != nil {
			log.Printf("Error saving message: %v", err)
		}

		h.broadcast <- msg
	}
}

func (h *WebSocketHandler) HandleMessages() {
	for {
		msg := <-h.broadcast
		for client := range h.clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Error broadcasting message: %v", err)
				client.Close()
				delete(h.clients, client)
			}
		}
	}
}

package socket

import (
	"log"
	"net/http"
	"time"

	"modern-tech/mini_project/domain/models"
	"modern-tech/mini_project/domain/service"

	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	service   *service.MessageService
	clients   map[*websocket.Conn]int // Map client connection to student ID
	broadcast chan models.Message
	upgrader  websocket.Upgrader
}

func NewWebSocketHandler(service *service.MessageService) *WebSocketHandler {
	return &WebSocketHandler{
		service:   service,
		clients:   make(map[*websocket.Conn]int),
		broadcast: make(chan models.Message),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (h *WebSocketHandler) HandleConnections(conn *websocket.Conn, studentID int) {
	defer conn.Close()
	h.clients[conn] = studentID

	for {
		var msg models.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			delete(h.clients, conn)
			break
		}

		msg.Timestamp = time.Now()
		h.broadcast <- msg

		// Lưu tin nhắn vào DB
		if err := h.service.SaveMessage(&msg); err != nil {
			log.Printf("Error saving message: %v", err)
		}
	}
}

func (h *WebSocketHandler) HandleMessages() {
	for {
		msg := <-h.broadcast
		for client, studentID := range h.clients {
			if studentID == msg.ReceiverID {
				err := client.WriteJSON(msg)
				if err != nil {
					log.Printf("Error broadcasting message: %v", err)
					client.Close()
					delete(h.clients, client)
				}
			}
		}
	}
}

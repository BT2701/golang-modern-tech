package week7

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMySQL() (*gorm.DB, error) {
	dsn := "username:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func RunServer() {
	db, err := ConnectMySQL()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&Message{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	repo := NewMessageRepository(db)
	service := NewMessageService(repo)
	handler := NewWebSocketHandler(service)

	go handler.HandleMessages()

	router := gin.Default()
	router.GET("/ws", func(c *gin.Context) {
		handler.HandleConnections(c)
	})

	log.Println("WebSocket server is running on ws://localhost:8080/ws")
	router.Run(":8080")
}

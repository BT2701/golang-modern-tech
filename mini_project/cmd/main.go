package main

import (
	"log"
	"modern-tech/mini_project/adapter/handler"
	"modern-tech/mini_project/adapter/socket"
	"modern-tech/mini_project/domain/models"
	"modern-tech/mini_project/domain/service"
	"modern-tech/mini_project/infrastructure/database"
	// "modern-tech/mini_project/infrastructure/redis"
	"modern-tech/mini_project/infrastructure/repository"
	"modern-tech/mini_project/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		_, err := jwt.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func main() {
	// Kết nối MySQL
	err := database.Connect("root", "root", "modern_tech", "localhost", 3306)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}

	// Kết nối Redis
	// redisClient := redis.NewRedisClient("localhost:6379", "", 0)

	// Tự động migrate
	database.DB.AutoMigrate(&models.Student{}, &models.Message{})

	// Khởi tạo repository và service
	studentRepo := repository.NewStudentRepository(database.DB)
	studentService := service.NewStudentService(studentRepo) // Correctly initialize StudentService
	// authService := service.NewAuthService(studentRepo, redisClient)
	messageRepo := repository.NewMessageRepository(database.DB)
	messageService := service.NewMessageService(messageRepo)

	// Khởi tạo handler
	studentHandler := handler.NewStudentHandler(studentService) // Pass StudentService here
	messageHandler := handler.NewMessageHandler(messageService)

	// Khởi tạo WebSocket handler
	wsHandler := socket.NewWebSocketHandler(messageService)
	go wsHandler.HandleMessages()

	// Thiết lập router
	router := gin.Default()

	// Auth routes
	router.POST("/auth/register", studentHandler.Register)
	router.POST("/auth/login", studentHandler.Login)

	// Protected routes
	protected := router.Group("/")
	protected.Use(AuthMiddleware())
	{
		// protected.GET("/students", studentHandler.GetAll)
		// protected.GET("/students/:id", studentHandler.GetByID)
		// protected.POST("/students", studentHandler.Add)
		// protected.PUT("/students/:id", studentHandler.Update)
		// protected.DELETE("/students/:id", studentHandler.Delete)

		protected.POST("/messages", messageHandler.SendMessage)
		protected.GET("/messages/:receiver_id", messageHandler.GetMessages)
	}

	// WebSocket route
	router.GET("/ws", func(c *gin.Context) {
		// Xử lý WebSocket
	})

	log.Println("Server is running on port 8080")
	router.Run(":8081")
}

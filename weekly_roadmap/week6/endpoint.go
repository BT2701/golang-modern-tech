package week6

import (
	"log"

	"github.com/gin-gonic/gin"
	"modern-tech/weekly_roadmap/week5"
)

func RunServer() {
	db, err := week5.ConnectMySQL()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&User{}, &week5.Student{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	userRepo := NewUserRepository(db)
	authService := NewAuthService(userRepo)
	authHandler := NewAuthHandler(authService)

	studentRepo := week5.NewStudentRepository(db)
	studentService := week5.NewStudentService(studentRepo)
	studentHandler := week5.NewStudentHandler(studentService)

	router := gin.Default()

	// Auth routes
	router.POST("/auth/register", authHandler.Register)
	router.POST("/auth/login", authHandler.Login)

	// Protected student routes
	studentRoutes := router.Group("/students")
	studentRoutes.Use(AuthMiddleware(authService))
	{
		studentRoutes.GET("", studentHandler.GetAll)
		studentRoutes.POST("", studentHandler.Add)
		studentRoutes.PUT("/:id", studentHandler.Update)
		studentRoutes.DELETE("/:id", studentHandler.Delete)
	}

	log.Println("Server is running on port 8080")
	router.Run(":8080")
}

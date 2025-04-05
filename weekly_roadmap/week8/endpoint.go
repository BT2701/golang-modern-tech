package week8

import (
	"log"
	"modern-tech/weekly_roadmap/week3"
	"modern-tech/weekly_roadmap/week5"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	db, err := week5.ConnectMySQL()
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}

	rdb := ConnectRedis()

	if err := db.AutoMigrate(&week3.Student{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	repo := NewStudentRepository(db, rdb)
	handler := NewStudentHandler(repo)

	router := gin.Default()

	router.GET("/students", handler.GetAllStudents)
	router.GET("/metrics", handler.GetMetrics)

	log.Println("Server is running on port 8080")
	router.Run(":8080")
}

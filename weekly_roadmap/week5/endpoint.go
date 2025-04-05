package week5

import (
	"log"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	// Kết nối tới cơ sở dữ liệu MySQL
	db, err := ConnectMySQL()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Tự động migrate bảng Student
	if err := db.AutoMigrate(&Student{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Khởi tạo repository, service và handler
	repo := NewStudentRepository(db)
	service := NewStudentService(repo)
	handler := NewStudentHandler(service)

	// Thiết lập router và đăng ký các route
	router := gin.Default()
	RegisterRoutes(router, handler)

	// Chạy server trên cổng 8080
	log.Println("Server is running on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

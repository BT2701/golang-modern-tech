package week5

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMySQL() (*gorm.DB, error) {
	// Thay đổi thông tin kết nối phù hợp với database của bạn
	dsn := "root:root@tcp(127.0.0.1:3306)/modern_tech?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil, err
	}

	fmt.Println("Successfully connected to MySQL using GORM!")
	return db, nil
}

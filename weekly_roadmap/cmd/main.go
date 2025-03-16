package main

import (
	"fmt"
	"os"
	"weekly_roadmap/week1"
	"weekly_roadmap/week2"
)

func main() {
	fmt.Println("WELCOME TO WEEKLY ROADMAP.")
	println("--------------------------------")
	fmt.Println("Please choose the options:")
	fmt.Println("1. Sum of integers (week1)")
	fmt.Println("2. Calculate area and perimeter of shapes (week2)")

	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Lỗi khi đọc lựa chọn:", err)
		os.Exit(1)
	}

	switch choice {
	case 1:
		err := week1.SumIntegers()
		if err != nil {
			fmt.Println("Lỗi:", err)
			os.Exit(1)
		}
	case 2:
		week2.HandleShape()
	default:
		fmt.Println("Lựa chọn không hợp lệ.")
		os.Exit(1)
	}
}

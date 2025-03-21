package main

import (
	"fmt"
	"modern-tech/weekly_roadmap/week1"
	"modern-tech/weekly_roadmap/week2"
	"modern-tech/weekly_roadmap/week3"
	"modern-tech/weekly_roadmap/week4"
	"os"
)

func main() {
	fmt.Println("WELCOME TO WEEKLY ROADMAP.")
	println("--------------------------------")
	fmt.Println("Please choose the options:")
	fmt.Println("1. Sum of integers (week1)")
	fmt.Println("2. Calculate area and perimeter of shapes (week2)")
	fmt.Println("3. Develop base api (week3)")
	fmt.Println("4. Go routines (week4)")

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
	case 3:
		week3.Week3_Port()
	case 4:
		week4.MainWeek4()
	default:
		fmt.Println("Lựa chọn không hợp lệ.")
		os.Exit(1)
	}
}

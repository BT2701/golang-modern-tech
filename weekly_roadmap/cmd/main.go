package main

import (
	"fmt"
	"os"
	"weekly_roadmap/week1"
)

func main() {
	fmt.Println("Chương trình tính tổng các số nguyên trong danh sách.")
	err := week1.SumIntegers()
	if err != nil {
		fmt.Println("Lỗi:", err)
		os.Exit(1)
	}
}

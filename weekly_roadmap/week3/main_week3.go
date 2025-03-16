package week3

import (
	"fmt"
)

func MainWeek3() {
	var choice int
	fmt.Println("Select a function to execute:")
	fmt.Println("1. Tạo 3 goroutines chạy đồng thời")
	fmt.Println("2. Chương trình tính tổng của một danh sách số bằng nhiều goroutines.")
	fmt.Println("3. Chương trình mô phỏng ngân hàng với nhiều tài khoản")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		GoRoutines()
	case 2:
		SumListGoRoutine()
	case 3:
		HandleBankAccount()
	default:
		fmt.Println("Invalid choice")
	}
}

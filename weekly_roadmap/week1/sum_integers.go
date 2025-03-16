package week1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func inputIntegers() []int {
	fmt.Println("Nhập vào danh sách số nguyên, cách nhau bởi dấu cách:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	strNumbers := strings.Split(input, " ")

	var numbers []int
	for _, str := range strNumbers {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Lỗi: Vui lòng nhập số nguyên hợp lệ.")
			return nil
		}
		numbers = append(numbers, num)
	}
	fmt.Println("Danh sách số nguyên đã nhập:", numbers)
	return numbers
}

func SumIntegers() error {
	numbers := inputIntegers()
	if numbers == nil {
		return fmt.Errorf("Danh sách số nguyên không hợp lệ.")
	}

	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("Tổng các số nguyên trong danh sách:", sum)
	return nil
}

package week4

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func sumPart(numbers []int, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	resultChan <- sum
}

func sumList(numbers []int, numGoroutines int) int {
	var wg sync.WaitGroup
	resultChan := make(chan int, numGoroutines)
	partSize := (len(numbers) + numGoroutines - 1) / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		start := i * partSize
		end := start + partSize
		if end > len(numbers) {
			end = len(numbers)
		}
		wg.Add(1)
		go sumPart(numbers[start:end], &wg, resultChan)
	}

	wg.Wait()
	close(resultChan)

	totalSum := 0
	for sum := range resultChan {
		totalSum += sum
	}

	return totalSum
}

func inputIntegers() []int {
	fmt.Println("Nhập vào danh sách số nguyên, cách nhau bởi dấu cách:")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Lỗi khi nhập danh sách số nguyên:", err)
		return nil
	}
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

func SumListGoRoutine() {
	numbers := inputIntegers()
	numGoroutines := 3
	totalSum := sumList(numbers, numGoroutines)
	fmt.Printf("Total sum: %d\n", totalSum)
}

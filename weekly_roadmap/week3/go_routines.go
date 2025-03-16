package week3

import (
	"fmt"
	"sync"
)

func GoRoutines() {
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 1: Hello from goroutine 1!")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 2: Hello from goroutine 2!")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 3: Hello from goroutine 3!")
	}()

	wg.Wait()
}
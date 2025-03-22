package week4

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func (a *BankAccount) Deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amount
}

func (a *BankAccount) Withdraw(amount int) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if amount > a.balance {
		return fmt.Errorf("insufficient funds")
	}
	a.balance -= amount
	return nil
}

func (a *BankAccount) Balance() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

func inputSimulateDeposit() int {
	var amount int
	fmt.Println("Nhập số tiền gửi vào tài khoản:")
	_, err := fmt.Scanln(&amount)
	if err != nil {
		fmt.Println("Lỗi khi nhập số tiền gửi:", err)
		return 0
	}
	return amount
}

func HandleBankAccount() {
	account := &BankAccount{}

	var wg sync.WaitGroup

	simulate := inputSimulateDeposit()

	// Simulate deposits
	for i := 0; i < simulate; i++ {
		wg.Add(1)
		go func(amount int) {
			defer wg.Done()
			account.Deposit(amount)
			fmt.Printf("Deposited %d, balance: %d\n", amount, account.Balance())
		}(i * 100)
	}

	// Simulate withdrawals
	for i := 0; i < simulate; i++ {
		wg.Add(1)
		go func(amount int) {
			defer wg.Done()
			err := account.Withdraw(amount)
			if err != nil {
				fmt.Printf("Failed to withdraw %d: %s\n", amount, err)
			} else {
				fmt.Printf("Withdrew %d, balance: %d\n", amount, account.Balance())
			}
		}(i * 50)
	}

	wg.Wait()
	fmt.Printf("Final balance: %d\n", account.Balance())
}

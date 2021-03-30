package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("vim-go")
}

var (
	mu      sync.Mutex
	balance int
)

//
func deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

func getBalance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}

package main

import (
	"fmt"
	"time"
)

func init() {
	go teller()
}

func main() {
	deposit(59)
	fmt.Println()
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(balance())
			deposit(100)
			time.Sleep(time.Second * 1)
		}
	}()
	deposit(500)
	time.Sleep(time.Second * 2)

	fmt.Println("before withdraw", balance())
	fmt.Println(withdraw(50))
	fmt.Println("after withdraw", balance()) // around 900
	time.Sleep(time.Second * 2)
	fmt.Println("before withdraw v2", balance())
	fmt.Println(withdraw(5000))                 // does substract and add for quick second
	fmt.Println("after withdraw v2", balance()) // 900
	// messy and multiple routines accessing var balance = bad

}

var deposits = make(chan int)
var balances = make(chan int)
var withdraws = make(chan int)

// substracts val from balance.
func withdraw(amount int) bool {
	withdraws <- amount
	if balance() <= 0 {
		deposit(amount) // subtracts and readds if below 0
		return false
	}
	return true
}

func deposit(amount int) { deposits <- amount }

func balance() int { return <-balances }

func teller() {
	var balance int // for teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case amount := <-withdraws:
			balance -= amount
		}
	}
}

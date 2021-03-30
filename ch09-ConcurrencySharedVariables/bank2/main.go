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
	go func(){
		for i := 0; i < 10 ; i++ {
		fmt.Println(balance())
		deposit(100)
		time.Sleep(time.Second*1)
	}
	}()
	deposit(50)
	time.Sleep(time.Second*5)
	deposit(500)
	time.Sleep(time.Second*5)

}

var deposits = make(chan int)
var balances= make(chan int)
var withdraws= make(chan int)

// TODO: fix withdraw so that is checks balance first

// Should check if balance is enough to withdraw, return true/false
func withdraw(amount int) { withdraws <-amount}

func deposit(amount int) {deposits <- amount}

func balance() int {return <-balances}

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


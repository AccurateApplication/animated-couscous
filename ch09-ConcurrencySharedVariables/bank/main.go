package main

import (
	"fmt"
	//"animated-couscousbank"
	bank "animated-couscous/ch09-ConcurrencySharedVariables/bank/bank123"
	"time"
)

func main( ) {
	bank.Deposit(100)
	fmt.Println("Pre go funcs", bank.Balance()) // 100
	x := bank.Balance() // x = 100
	fmt.Println("x ", x) 
	time.Sleep(time.Second*1)
	go func () {
		bank.Deposit(100) 
		fmt.Println(bank.Balance(),"Go func Bal") // 300
		time.Sleep(time.Second*1)
	}()
	go bank.Deposit(100)
	time.Sleep(time.Second*1)
	fmt.Println("after go bank depsoit", bank.Balance()) // 300
	time.Sleep(time.Second*1)

}

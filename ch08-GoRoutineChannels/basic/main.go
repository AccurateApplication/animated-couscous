package main

import (
	"fmt"
	"time"
)

/*
Goroutine similiar to threads


*/

func main() {
	fmt.Println("go routine start")
	go rotine(1 * time.Second) // calls function rotine, doesnt wait for return

	time.Sleep(7 * time.Second)
	fmt.Println("sleep over")
}

func rotine(delay time.Duration) {
	fmt.Println("before start")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
		time.Sleep(delay)
	}
	fmt.Println("after start")
}

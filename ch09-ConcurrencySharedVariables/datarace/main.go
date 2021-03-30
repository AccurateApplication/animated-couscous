package main

import (
	"fmt"
	"time"
)

func main() {

	go racer()
	time.Sleep(time.Second*1)

	var x []int
	go func() {
		x := make([]int,10) 
		fmt.Println(x)
	}()
	go func() {
		x := make([]int,10000) 
		fmt.Println(x)
	}()
	time.Sleep(time.Second*1)

	x[9999] = 1 // memory corruption possible as x could be slice of 10/nil

}

// Data race happens when 2 goroutines access same var concurrently
func racer() {
	chWait := make(chan struct{})
	n := 0
	go func() {
		n++
		close(chWait)
	}()
	n++
	<-chWait
	fmt.Println(n) // n could be 1/2
}

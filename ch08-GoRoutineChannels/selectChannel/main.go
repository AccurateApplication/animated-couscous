package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; 1 < n; i++ {
		c <- x
		x, y = y, x+y
	}
	defer close(c)

}

func main() {
	kanal := make(chan int, 15)
	go fibonacci(cap(kanal), kanal) // cap of kanal=number of times to loop
	for i := range kanal {
		time.Sleep(1 * time.Second)
		fmt.Printf("%v\n", i)
	}

}

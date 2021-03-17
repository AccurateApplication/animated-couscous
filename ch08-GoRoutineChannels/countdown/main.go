package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	fmt.Println("countdown start")
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("time after 10")
	case <-abort:
		fmt.Println("abort!") // hits here after 10 sec. why?
		return
	}

	launch()
}

func launch() {
	fmt.Println("lets go")
}

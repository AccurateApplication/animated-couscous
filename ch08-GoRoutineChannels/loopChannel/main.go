package main

import "fmt"

func main() {
	queue := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			queue <- i
		}
		close(queue)
	}()
	for elem := range queue {
		fmt.Println(elem) // 1234567..
	}
	// close(queue) = panic , close of closed channel

	kanal := something()
	for elem := range kanal {
		fmt.Println(elem) // 1,2,3,4,5..
	}
}

// returns channel with ints
func something() <-chan int {
	queue := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			queue <- i
		}
		// defer close(queue) works with both defer and normal
		close(queue)
	}()
	return queue
}

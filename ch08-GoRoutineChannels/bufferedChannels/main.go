package main

import "fmt"

func main() {
	// len on buffered channel gets number of currently buffered elements
	// cap on buffered channel, gets capacity

	ch := make(chan string, 5)
	fmt.Println(len(ch), cap(ch)) // 0, 5
	ch <- "A"
	ch <- "B"
	fmt.Println(len(ch), <-ch, <-ch) // 2, A, B
}

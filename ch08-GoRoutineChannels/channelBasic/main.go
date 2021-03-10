package main

import (
	"fmt"
	"time"
)

/*
ch := make(chan int),  ch has type chan int
channel = reference to data structure created by make
zero val of channel = nil

// Operations
ch <- x   | send statement
x = <-ch  | recieve expression, in assignment statement
<-ch      | receive statement, result discard
close(ch) | close channel

*/

func main() {
	// buffered channel
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
	//fmt.Println(<-ch) // can only print twice
	ch <- 3 //
	//ch <- 10 // crash
	fmt.Println(<-ch) // 3
	time.Sleep(1 * time.Second)

	// unbuffered chan
	//ch1 := make(chan int, 0)
	//ch1 <- 1 // chan send, goroutines are asleep deadlock
	//fmt.Println(<-ch1)

	sendChannels2("hello123")
	fmt.Println("msg in main", msg)
	sendChannels("bird is the word")
	fmt.Println(msg) // msg only changed in sendChannels

}

var msg = make(chan string)  // channel msg
var msg2 = make(chan string) // channel msg
func sendChannels(word string) {
	//wg := new(sync.WaitGroup)

	go func() {
		msg <- word
		//wg.Add(1)
	}() // anonymous func
	message := <-msg
	for i := 0; i < 5; i++ {
		fmt.Printf("%d message: %s\tmsg: %v \n", i, message, msg)
		time.Sleep(1 * time.Second)
	}
	//wg.Wait()
	defer close(msg)

}

func sendChannels2(word string) {
	//for i := 0; i < 5; i++ {
	//	msg2 <- word
	//}
	//x := <-msg2

	go func() { msg2 <- word }() //
	message := <-msg2
	for i := 0; i < 5; i++ {
		fmt.Printf(" message: %s\tmsg: %v \n", message, msg2)
		time.Sleep(1 * time.Second)
	}
	defer close(msg2)

}

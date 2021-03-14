package main

import "fmt"

/*
X chan <- int  | send only channel, allow send no receive

X <-chan int  | receive channel., allow receieve not send

*/

// sends to out, not receive. Adds numbers up to 99
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

// send to out, receive from in.
func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

// print 'in'
func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)          // naturals channel now consists of ints 0-99
	go squarer(squares, naturals) // sends squared numbers of ints in naturals to squares channel
	//printer(squares)              // Print squares

	//for elem := range squares {  same as printer function
	//	fmt.Println(elem)
	//}
}

package main

import "fmt"

func main() {

	x := address()

	y := *x
	fmt.Println(x, &x, *x)
}

func address() *int {
	var x int
	return &x
}

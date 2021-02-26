package main

import "fmt"

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x

	}
}

func main() {
	fmt.Println("vim-go")
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	for i := 0; i < 10; i++ {
		fmt.Println(f())

	}
}

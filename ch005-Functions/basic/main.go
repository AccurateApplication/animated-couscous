package main

import "fmt"

/*
func name(parameters-list) result {
	body
}
:wq


parameters are local variables within body of func

*/

func main() {

	fmt.Println(appending(5, 50))
	x := factorial_calc(10)
	fmt.Println(x)
}

func appending(x, y int) int {
	return x + y
}

var count = 10

func factorial_calc(number int) int {
	if number == 0 || number == 1 {
		return 1
	}
	if number < 0 {
		fmt.Println("Invalid number")
		return -1
	}
	return number * factorial_calc(number-1)
}

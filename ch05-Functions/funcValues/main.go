package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("vim-go")
	f := square
	fmt.Println(f(5))
	fmt.Println(product(5, 4))

	fmt.Println(strings.Map(rmSpace, "text with space removed"))
	fmt.Println(strings.Map(add1, "lucas"))
}

func square(n int) int     { return n * n }
func product(m, n int) int { return m * n }
func add1(r rune) rune     { return r + 1 }
func rmSpace(r rune) rune {
	if r == ' ' {
		return 0
	}
	return r
}

package main

import (
	"fmt"
)

type Person interface {
	greet() string
}

type Human struct {
	Name string
}

func (h *Human) greet() string {
	return "Im " + h.Name
}

func isAPerson(h Person) {
	fmt.Println(h.greet())
}

func main() {
	var a interface{}
	fmt.Println(a) // nil

	var b = Human{"StringName"}
	fmt.Println("b greet()\t", b.greet()) // normal method

	// isAPerson(b) // Cannot use type human as type Person

	isAPerson(&b)
}

package main

import "fmt"

type test interface {
	Area() float64
	Perimeter() float64
}

type square struct {
	width  float64
	length float64
}

type squareTest struct {
	width  float64
	length float64
}

func (s square) Area() float64 {
	return s.width * s.length
}
func (s square) Perimeter() float64 {
	return (2 * s.width) + (2 * s.length)
}

func main() {
	fmt.Println()
	var s test
	fmt.Printf("type %T\n", s) // nil
	// fmt.Println(s.Perimeter(), s.Area()) // runtime error, invalid memory addr/nil ptr
	s = square{10, 20}
	fmt.Println(s.Perimeter(), s.Area()) // 60 200
	//s = squareTest(50, 50)               // does not work
	//fmt.Println(s.Perimeter())           // too many arg's to conversion

	// ---------------------------
	// type assertiuon
	var i interface{}
	x, ok := i.(string)
	fmt.Printf("x: %v \tbool: %t \n", x, ok) // "" , false

	var i1 interface{} = "test"
	x1, ok := i1.(string)
	fmt.Printf("x: %v \tbool: %t \n", x1, ok) // test, true

	fl, ok := i.(float64)
	fmt.Printf("f1: %v \tbool: %t \n", fl, ok) // 0, false
}

package main

import (
	"fmt"
	"os"
)

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("%g F or %g C\n", f, c)
	fmt.Println("print func\t", fToC(f))
	var a, b, d = true, 2.3, "string"
	fmt.Printf("%T\t%T\t%T\t\n", a, b, d)
	var file, err = os.Open("doesNotExist")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file) // nil
	file.Close()
	x, y := 50, 100
	fmt.Printf("before swap x:%d\ty:%d\t\n", x, y)
	x, y = y, x
	fmt.Printf("swap x:%d\ty:%d\t\n\n\n", x, y)
	p := &x
	fmt.Println("p of type int, points to X, aboev", *p)
	fmt.Println("this gets the address of X?", p)
	*p = 1237716237163
	fmt.Println("x now changed", x)

}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

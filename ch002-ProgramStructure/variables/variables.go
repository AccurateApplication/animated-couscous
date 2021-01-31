package main

import "fmt"

func main() {
	p := new(int)
	fmt.Printf("p: %d\t*p: %d\tp&: %d\n", p, *p, &p)
	*p = 5
	fmt.Printf("p: %d\t*p: %d\tp&: %d\n", p, *p, &p)
	p = new(int)
	q := new(int)
	fmt.Printf("p: %d\t*p: %d\tp&: %d\n", p, *p, &p)
	fmt.Println(p == q) // false
	// new(T) // unnamed var
	// fmt.Println(T, *T)
	fmt.Printf("global: %d\n", global) // cant add " * " to this one, runtime error. why? (global = 0)
	xd()
	fmt.Printf("global: %d\n", *global) // global = 1
	fmt.Println("\n----\n ")
	medals := []string{"gold", "silver", "bronze"} // slice
	for i, val := range medals {
		fmt.Println(i, val)
	}
}

var global *int

func xd() {
	var x int
	x = 1
	global = &x
}

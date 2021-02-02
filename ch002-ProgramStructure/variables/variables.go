package main

import (
	"fmt"
	"log"
	"os"
)

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
	X := "hello!"
	for i := 0; i < len(X); i++ {
		X := X[i]
		if X != '!' {
			X := X + 'A' - 'a'
			fmt.Printf("%c\n", X)
		}
	}
	// init()
	fmt.Println(cwd)
}

var cwd string

func init() { // this will be used whether called or not, can print cwd in main for example
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}

var global *int

func xd() {
	var x int
	x = 1
	global = &x
}

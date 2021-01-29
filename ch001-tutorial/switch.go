package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(rand.Intn(99))
	fmt.Println(rand.Intn(50))
	x := rand.Intn(52)
	switch {
	case x%2 == 0:
		fmt.Printf("x(%d) is even\n", x)
	default:
		fmt.Printf("x(%d) is odd\n", x)
	}

}

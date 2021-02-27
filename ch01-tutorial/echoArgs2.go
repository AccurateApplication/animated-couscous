package main

// use range os args to get arg given

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str, sep := "", ""
	for _, args := range os.Args[1:] {
		str += sep + args
		sep = " "
	}
	fmt.Println(str)
	fmt.Println("\nis.Args[1:] = ", os.Args[1:])
	fmt.Println("-")
	fmt.Println(strings.Join(os.Args[1:], " "))
}

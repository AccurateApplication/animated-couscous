package main

import (
	"fmt"
	"os"
)

func main() {
	var s, seperator string
	for i := 1; i < len(os.Args); i++ { // add arguments to s?
		seperator = " "
		s += os.Args[i] + seperator
	}
	/*
		for i, val := range s {
			fmt.Printf("indx: %d w/ val: %s\n", i, val)
			fmt.Println("println val", val)
		}
	*/
	for i, val := range os.Args[:] {
		fmt.Printf("indx: %d w/ val: %s\n", i, val)
	}
}

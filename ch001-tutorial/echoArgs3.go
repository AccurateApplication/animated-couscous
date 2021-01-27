package main

import (
	"fmt"
	"os"
	"strings"
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
	fmt.Println("\nstrings join:\t", strings.Join(os.Args[1:], " ")) // This is more efficent that using loop/cat as "s" will be garbage collected when contents of S no longer in use
	for i, val := range os.Args[:] {
		fmt.Printf("indx: %d w/ val: %s\n", i, val)
	}
}

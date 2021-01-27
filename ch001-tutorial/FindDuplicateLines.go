package main

// like unix uniq command, but worse. pipe text with \n to command

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)      // Map holds set of key/val pairs
	input := bufio.NewScanner(os.Stdin) //scanner from bufio pkg easy to process input
	for input.Scan() {
		counts[input.Text()]++ // Append to counts map. Every line = one key/val pair
	}
	// Ignores errors from input.err(), how to add?
	for line, num := range counts { // loop through map and print
		if num > 1 {
			fmt.Printf("%d\t%s\n", num, line)
		}

	}
}

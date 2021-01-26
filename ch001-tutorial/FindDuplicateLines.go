package main

// like unix uniq command, but worse. pipe textfile to command

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// Ignores errors from input.err(), how to add?
	for line, num := range counts {
		if num > 1 {
			fmt.Printf("%d\t%s\n", num, line)
		}

	}
}

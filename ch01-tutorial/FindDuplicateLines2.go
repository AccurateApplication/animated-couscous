package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// will use map this time aswell
	counts := make(map[string]int)
	files := os.Args[1:] // take file as input instead
	if len(files) == 0 {
		countLines(os.Stdin, counts)
		fmt.Println("len file = 0")
	} else {
		for i, arg := range files {
			f, err := os.Open(arg)
			fmt.Println(i, arg)
			if err != nil { // std error handlnig
				fmt.Fprintf(os.Stderr, "dup: %v\n", err) // difference between Fprintf and Printf?
				continue                                 // if we get error, still continue loop above (?)
			}
			countLines(f, counts)
			f.Close() // can "defer" be used?
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) { // when map passed to function, will be seen by func main
	input := bufio.NewScanner(f) // ignoring err from input.Err()
	for input.Scan() {
		counts[input.Text()]++ // append to counts.
	}
}

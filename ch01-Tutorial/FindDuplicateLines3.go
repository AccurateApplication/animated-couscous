package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] { // all args given should now be files?
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err %v\n", err)
			continue // if we get err, still loop. just prnt error
		}
		// readfile gives us byte slice that we convert with to string with: string(data) below
		for _, line := range strings.Split(string(data), "\n") { // if we dont split \n we will only get 1 string
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

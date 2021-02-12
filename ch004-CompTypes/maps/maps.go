package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
 maps = hash map
 keys in given map are same
 all values same type
 but val and key can be different types


*/

func main() {
	ages := map[string]int{
		"alice": 31,
		"bob":   45,
	}
	// x := &ages["bob"] // cannot take address since map element is not variable

	for i, v := range ages {
		fmt.Println(i, ages[i], v)
	}
	// var names []string
	names := make([]string, 0, len(ages)) // since we know final size of names (len ages), we can declare length of slice before hand, more effiecent
	for name := range ages {
		names = append(names, name)
	}
	fmt.Println(ages["bob"], "<-- || ", names)
	age, ok := ages["bob"]
	if !ok {
		fmt.Println("age !ok")
	}
	age2, ok := ages["nobody"]
	if !ok {
		fmt.Println(" age2 !ok") // this print
	}
	fmt.Println("age:", age, "\t age2:", age2)
	var testMap map[string]int
	fmt.Printf("%t\t%t\n", testMap == nil, len(testMap) == 0) // true, true.  zero map val is nil,
	/*
		/
		/
		/
	*/
	seen := make(map[string]bool) // set of strings, bool as value
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line, "<- line in: if !seen")
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

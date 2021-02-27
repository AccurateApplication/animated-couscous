package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

/*
https://en.wikipedia.org/wiki/Depth-first_search


*/

func main() {
	fmt.Println("")
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
	//for _, x := range prereqs {
	//	fmt.Println(x)
	//}
	s := fmt.Sprintf("sprintf test\n")
	io.WriteString(os.Stdout, s)
	fmt.Println(s)
}
func topoSort(m map[string][]string) []string { // m = prereqs map in main
	var result []string
	seen := make(map[string]bool)     // will use this to true/false of string.
	var visitAll func(items []string) // recursive, anonymous function?
	visitAll = func(items []string) {
		for _, item := range items {
			for !seen[item] { // if not seen?  (item = false?)
				seen[item] = true
				visitAll(m[item]) // runs function on m[items]
				//fmt.Println("m item: ", m[item]) //
				result = append(result, item) // append to result
			}
		}
	}
	var keys []string
	for key := range m { // add elements/whatever is called from m (prereqs) to keys
		keys = append(keys, key)
		//fmt.Printf("Key: %s\t", key)
	}
	sort.Strings(keys) // sort keys, why?
	visitAll(keys)     // calls itself but now with keys instead
	return result

}

// to be continued

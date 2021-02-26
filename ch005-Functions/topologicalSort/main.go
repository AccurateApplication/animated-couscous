package main

import "fmt"

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

func main() {
	fmt.Println("vim-go")
	//for i, course := range topoSort(prereqs) {
	//	fmt.Printf("%d:\t%s\n", i+1, course)
	//}
	for _, x := range prereqs {
		fmt.Println(x)
	}
}

//func topoSort(m map[string][]string)[]string {
//	var order []string
//	seen := make(map[string]bool) // will use this to true/false.
//    var visitAll func(items []string)
//    visitAll = func(items []string) {
//        for !seen[item] {
//            seen[item] = true
//            visitAll(m[item]) // unsure what this does...
//
//
//			// to be continued

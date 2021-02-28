package main

import (
	"animated-couscous/ch05/links"
	"os"

	// "animated-couscous/ch05/links"
	"fmt"
	"log"
)

func main() {
	fmt.Println("vim-go")
	breadthFirst(crawl, os.Args[1:])
}

//func breadthFirst(f func (item string) []string, worklist []string {
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil               // hmm. cause we dont need worklist now?
		for _, item := range items { // item=worklist elements
			if !seen[item] { // seen used to only visit items once
				seen[item] = true
				worklist = append(worklist, f(item)...) // f(item)... = all items
				// all items in list added to worklist
				// Func returns when loop done, worklist is then modified with added items
			}
		}
	}
}

func crawler(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

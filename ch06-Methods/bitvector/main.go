package main

import "fmt"

type IntSet struct {
	words []uint64
}

// Reports if set contains non-negative value X
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
	// x/64 less than len intset words AND ......?
	// Confusing, look at another day..
}

func main() {
	fmt.Println("vim-go")
}

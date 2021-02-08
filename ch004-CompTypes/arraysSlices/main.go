package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	// array fixed length sequence of zero/more elements of type
	// Slices are almost the same, but can grow and shirnk, more versatile and used more in  go.

	var a [3]int // array, 3 ints. val's initially set to 0
	// var b [5]int = [5]int{1, 4, 5} // 1,4,5,0,0
	b := [5]int{1, 4, 5}
	var c []int // slice
	fmt.Println(a, b)
	for i := 0; i < 10; i++ { // append 10 numbers to slice
		c = append(c, i)
	}
	fmt.Println(c, "<- slice")
	cr1 := sha256.Sum256([]byte("X"))
	cr2 := sha256.Sum256([]byte("Y"))
	// fmt.Println(len(cr1), cr1)
	fmt.Printf("len: %d\nbase16 X: %x\nbase16 Y: %x\n", len(cr1), cr1, cr2) // %x = base16
	revSlice(c)
	fmt.Println("reveresed slice?:", c)
}

func zero(ptr [32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func revSlice(s []int) {
	for i, x := 0, len(s)-1; i < x; i, x = i+1, x-1 {
		s[i], s[x] = s[x], s[i]
	}
}

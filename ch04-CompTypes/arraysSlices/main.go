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
	//////
	revSlice(c)
	fmt.Printf("capacity C: %v\t reveresed slice: %v\n", cap(c), c)
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("i: %d\tcap y=%d\ty:%v\n", i, cap(y), y)
		x = y
	}
	// testSlice := make([]int, 5, 50) //  slice with 5 elements and cap of 50
	// fmt.Printf("testslice: %v\t cap: %v\t len: %v", testSlice, cap(testSlice), len(testSlice))

	const smpl = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	for i := 0; i < len(smpl); i++ {
		fmt.Printf("%x ", smpl[i])
	}
	fmt.Printf(" \n%x ", smpl)
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

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) { // if capacity of X more than len X, can grow/extend
		z = x[:zlen]
	} else { // we dont have room to extend
		zcap := zlen
		if zcap < 2*len(x) { // if capacity of Z is less than 2*len X then double
			// Grows by doubling
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap) //
		fmt.Println("zlen/zcap:", zlen, zcap)
		// fmt.Printf("zlen:%v \tzcap:%v\n", zlen, zcap)
		copy(z, x) // Copies elements from Z to X.
	}
	z[len(x)] = y //
	return z
}

//func appendInt2(x []int, y ...int) []int { // ... (ellipsis) makes function "variadic", meaning it accepts any number of final arg's

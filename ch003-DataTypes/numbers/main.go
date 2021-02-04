package main

import (
	"fmt"
	"math"
)

func main() {
	var i int8 = 127
	fmt.Println(i, i*1, i*i)
	var u uint8 = 255
	fmt.Println(u, u*1, u*u)
	f := 1.99           // float
	fmt.Println(int(f)) // discard fractional part, down
	var apple int32 = 1
	var banana int16 = 40
	// var fruit int = apple + banana // mismatched
	var fruit = int(apple) + int(banana) // this works
	fmt.Println(fruit)
	nan := math.NaN()
	fmt.Println(nan, nan == nan) // false
	runee := 'r'
	fmt.Println(runee) // 114
	maxFloat := math.MaxFloat32
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	fmt.Println(maxFloat, x, y, x*y)
	fmt.Println(string(65)) // A
	s := "abc"              // string
	b := []byte(s)          // convert string to byte slice
	s2 := string(b)         // string again
	fmt.Println(b, s, s2)
	// ofor ?, x := range Oi

}

/*
comparision
== equal to
!= not equal to
< less than
<= less than / equal
> more than
>= greater than/ equal
*/

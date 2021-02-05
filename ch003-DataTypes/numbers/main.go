package main

import (
	"fmt"
	"math"
)

// constants is something that wont change. Useful for example math stuff
type Weekday int

const (
	Sunday Weekday = iota // this declares sunday to be 0, monday 1, and so on
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	a = 1
	b
	c = 2
	d
	Pi = 3.14159265358979323846264338327950288419716939937510582097494459
)

func main() {
	fmt.Println(Sunday, Monday, "<- su mon")
	fmt.Println(a, b, c, d, "<- a,b,c,d") // 1,1,2,2
	fmt.Println(Pi)

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
	fmt.Println(b, s, s2, "\n\n\n")
	fmt.Println(Waa())
	hej := Waa()
	fmt.Println(hej)

}

const (
	_ = 1 << (10 * iota) // powers of 1024
	KiB
	MiB
)

func Waa() int {
	fmt.Println(KiB, MiB)
	hej := (KiB + MiB)
	return hej

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

package main

import (
	"fmt"
	"math"
)

func main() {
	p := Point{1, 20}
	q := Point{10, 20}
	fmt.Println(Dist(p, q))
	fmt.Println(p.Dist(p))
	fmt.Println(p.Dist(q)) // p.Dist == selector
	perim := Path{
		{1, 5},
		{3, 8},
		{5, 4},
		{5, 5},
	}
	fmt.Println(perim.Dist())
	// Methods with pointers
	// Using ScaleBy Func
	x := &Point{1, 2}      //pointer to Point ?
	x.ScaleBy(5)           // use func
	fmt.Println("*x:", *x) // Point=5,10

	x1 := Point{5, 10}
	x1.ScaleBy(5)          // this also works. compiler adds & before var x. This doesnt work directly on Point, as we cant get address of it
	fmt.Println("x1:", x1) // Point=25,50
	// Point{5, 70}.ScaleBy(5) = err: cannot call pointer method on point literal
}

type Point struct{ X, Y float64 }

// standard function
func Dist(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Method of point type
func (p Point) Dist(q Point) float64 { // p Point = receiver (like self in python?)
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

func (path Path) Dist() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Dist(path[i])
		}
	}
	return sum
}

// Calling functions makes copy of each arg value, if needs to update variable/variable too large to copy, use pointer

// Method called '  (*Point).ScaleBy.   '
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

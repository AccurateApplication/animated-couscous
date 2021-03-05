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

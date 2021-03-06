package main

import (
	"fmt"
	"image/color"
	"math"
)

func main() {
	//y := []int{12, 3, 5}
	x := IntList{5, nil}
	fmt.Println(x.Sum())

	var cp ColoredPoint
	cp.X = 1
	cp.Y = 5
	red := color.RGBA{255, 0, 0, 0}
	fmt.Println(cp.Point)
	//var q = ColoredPoint{Point{1, 5}, red}
	var p = ColoredPoint{Point{5, 1}, red}
	fmt.Println(p)
	p.ScaleBy(5)
	fmt.Println(p)

}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type Point struct{ X, Y float64 }
type ColoredPoint struct {
	Point
	Color color.RGBA
}

func Dist(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// IntList linked list of int. nil *IntList = empty list
type IntList struct {
	Value int
	Tail  *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

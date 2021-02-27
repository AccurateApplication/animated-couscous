package main

import "fmt"

type Employee struct { // name of struct field is exported if begin with Capital letter
	ID            int
	Name          string
	Motivated     bool
	IsEmployed    bool
	GoodWorkEthic bool
}

type Point struct{ X, Y int }

// Struct embedding . supposedly called inheritance
type Circle struct {
	Center Point // Point is embedded within Circle
	Radius int
}
type Wheel struct {
	Circle // Circle embeddded within Wheel
	Spokes int
}

var w Wheel

var Isak Employee
var Lucas Employee

// const Isak = Employee

func main() {
	Isak.ID = 50
	Isak.IsEmployed = true
	Isak.Motivated = true
	Lucas.Name = "Lucas"
	Isak.Name = "Isak"
	Isak.fire()
	Lucas.employ()
	Lucas.talk()
	fmt.Println(Isak, Lucas)

	var a int = 100
	var b int = 200
	fmt.Printf("Before swap: A: %d\tB: %d\n", a, b)
	swap(a, b)
	fmt.Printf("after swap: A: %d\tB: %d\n", a, b) // still same

	// Compare structs:
	p := Point{1, 2}
	q := Point{2, 1}
	fmt.Println("Compare struct:", p.X == q.X && p.Y == q.Y) // false
	fmt.Println("Compare struct:", p.Y == q.X && p.X == q.Y) // true

	w.Circle.Center.X = 8
	w.Circle.Center.Y = 5
	fmt.Println(w)
}

func (e *Employee) fire() {
	e.IsEmployed = false
}

func (e *Employee) employ() {
	e.IsEmployed = true
	e.GoodWorkEthic = true
}

func (e *Employee) talk() {
	/* pointer is required if function must modify argument. Otherwise only gets copy of arg, not reference to original arg
	shown below in func swap
	*/
	fmt.Println("My name is: ", e.Name)
}

// Does not actually change x,y as no pointers
func swap(x, y int) int {
	var temp int
	temp = x
	x = y
	y = temp
	return temp
}

package main

import "fmt"

type Employee struct { // name of struct field is exported if begin with Capital letter
	ID            int
	Name          string
	Motivated     bool
	IsEmployed    bool
	GoodWorkEthic bool
}

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

}

func (e *Employee) fire() {
	e.IsEmployed = false
}

func (e *Employee) employ() {
	e.IsEmployed = true
	e.GoodWorkEthic = true
}

func (e Employee) talk() {
	fmt.Println("My name is: ", e.Name)
}

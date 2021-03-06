package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary int
}

// func (receiver receiver_type) some_func_name(arguments) return_values

func (e employee) setNewName(newName string) {
	e.name = newName
}

// now with pointer, which change value directly and no copy
func (e *employee) setNewNamePtr(newName string) {
	e.name = newName
}

func (e employee) getName() string {
	return e.name
}

func (e employee) PrintEmployee() employee {
	fmt.Printf(" \nprint method\nName: %s\nAge: %d\nSalary: %d", e.name, e.age, e.salary)
	return e
}

func main() {
	emp1 := employee{name: "Andres", age: 51, salary: 25000}
	fmt.Println(emp1)
	emp1.setNewName("NoName")
	fmt.Println(emp1.name) // same name. (Andres)
	// This is because func passed it as copy and old values are printed

	emp2 := employee{name: "Rodriguez", age: 500, salary: 50000}
	emp2.setNewNamePtr("NoName")
	fmt.Println(emp2)

	fmt.Printf("Get employee name: %s\n", emp1.getName())

	(&emp2).setNewNamePtr("Name2")
	fmt.Println(emp2.getName(), "emp2 changed name again\n\n ")

	emp2.PrintEmployee()

}

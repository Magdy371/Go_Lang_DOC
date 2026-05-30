package main

import (
	"fmt"
)

func main() {
	var person1 Person = Person{
		name:   "John Doe",
		age:    30,
		salary: 50000.00,
	}
	fmt.Println(person1.personInfo())
}

type Person struct {
	name   string
	age    int
	salary float64
}

func (p Person) personInfo() string {
	return fmt.Sprintf("Name: %s, Age: %d, Salary: %.2f", p.name, p.age, p.salary)
}

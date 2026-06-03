package main

//Struct in Go meets Classes in c++ but no inheritance concept is found
//every struct is stand it alone with his own types and methods
import (
	"fmt"
)

func main() {
	var human human = human{"magdy", 26, "Damanhur", 200000}
	fmt.Println(human)
	//to print names and values
	fmt.Printf("%+v\n", human)
	fmt.Printf("Name is:%v\nSalary is: %v\n", human.name, human.salary)
}

type human struct {
	name   string
	age    int
	addr   string
	salary int
	//create a method to set these data

}

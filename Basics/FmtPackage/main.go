package main

import (
	"fmt"
)

func main() {
	const aValue = 10
	const aNumber = 10
	name := "magdy"
	fmt.Println("my value is: ", aValue)
	fmt.Printf("my number is: %v and its type is a number \n", aNumber)
	fmt.Printf("my name is: %v and its type is: %T\n", name, name)
}

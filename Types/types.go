package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello from Go Lang types")
	//1- Explicit type
	var name string = "my name is magdy"
	fmt.Println(`my name is: `, name)
	var phone int = 0101011
	fmt.Print("my phone is: ", phone)

	//implicit types
	address := "Damanhur CITY"
	postalCode := 22512
	fmt.Println("my postal code is: ", postalCode)
	fmt.Println("my address is: ", address)

	//Constant is a simple unchanged value
	//Explicit version is
	const aNumber int = 10
	//implicit version
	const aValue = 10
	fmt.Println("my value is: ", aValue)
	fmt.Println("my number is: ", aNumber)

	//Numeric Types
	/*
		byte[int8, uint8]
		uint16, int 16
		int, uint]->[int 64, uint64, uint23, int32] based on os
		uintptr
		float64, float32
		complex64, complex128
	*/
	stringLenth, err := fmt.Println("my value is: ", aValue)
	if err == nil {
		fmt.Println("the sting length is", stringLenth)
	}
}

package main

import "fmt"

func main() {
	var colors [3]string
	colors[0] = "red"
	colors[1] = "green"
	colors[2] = "blue"
	fmt.Printf("Existing colors are %v, and the array length is %v \n", colors, len(colors))

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Existing numbers are %v, and the array length is %v \n", numbers, len(numbers))

	// Printing the array
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("index %v = %v \n", i, numbers[i])
	}
}

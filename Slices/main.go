package main

/**
slices
are abstraction layer set in top of the array
But it is resizable , and easily sorted
it must contain data of the same type
**/
import (
	"fmt"
	"sort"
)

func main() {
	//Declare a slice like array
	colors := []string{"red", "blue", "green"}
	fmt.Println(colors)

	//method 2
	var color2 = make([]string, 0, 3)
	color2 = append(color2, "red", "blue", "green")
	fmt.Println(color2)
	fmt.Println(len(colors))
	color2 = append(color2, "Fuschia", "Purple", "Purple")
	fmt.Println(color2)
	fmt.Println(len(color2))
	color2 = remove(color2, 0)
	color2 = remove(color2, 3)
	fmt.Println(color2)
	//Sorting a slice
	sort.Strings(color2)
	fmt.Println(color2)
}

func remove(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}

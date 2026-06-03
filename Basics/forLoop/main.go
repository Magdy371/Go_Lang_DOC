package main

import (
	"fmt"
)

func main() {
	var numbers = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("index %v = %v \n", i, numbers[i])
	}

	var listGroceries map[string]string = make(map[string]string)
	listGroceries["fruits"] = "apple"
	listGroceries["vegetables"] = "Cucamber"
	listGroceries["meats"] = "beef"
	fmt.Printf("listGroceries are: %v \n", listGroceries)

	for k, v := range listGroceries {
		fmt.Printf("key: %v, value: %v \n", k, v)
	}

	//For like while loop
	var val int = 0
	var sum int = 0
	for val <= 5 {
		sum += val
		fmt.Printf("val: %v \n", val)
		fmt.Printf("sum: %v \n", sum)
		val++
	}
}

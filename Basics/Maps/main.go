package main

import "fmt"

func main() {
	var listGroceries map[string]string = make(map[string]string)
	listGroceries["fruits"] = "apple"
	listGroceries["vegetables"] = "Cucamber"
	listGroceries["meats"] = "beef"
	fmt.Printf("listGroceries are: %v \n", listGroceries)

	for k, v := range listGroceries {
		fmt.Printf("key: %v, value: %v \n", k, v)
	}

}

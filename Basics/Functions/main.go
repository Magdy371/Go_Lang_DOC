package main

import "fmt"

func main() {
	sum, counter, average := multiValuereturn(12, 3, 4, 57, 776)
	fmt.Printf("sum = %v\n", sum)
	fmt.Printf("average = %v\n", average)
	fmt.Printf("counter = %v\n", counter)
}

func multiValuereturn(value ...int) (int, int, float64) {
	var sum int = 0
	var avarage float64 = 0
	var counter int = 0
	for i := 0; i < len(value); i++ {
		sum += value[i]
		counter++
	}
	avarage = float64(sum) / float64(counter)
	return sum, counter, avarage
}

package main

import (
	"fmt"
	"time"
)

func main() {
	var weekDay time.Weekday = time.Now().Weekday()
	var weeDayN int = int(weekDay)
	fmt.Printf("Todat is %v and its number = %v \n", weekDay, weeDayN)
	var result string
	switch weekDay {
	case 1:
		result = " today Monday"
	case 2:
		result = " today Tuesday"
	case 3:
		result = " today Wendsday"
	case 4:
		result = " today Thrusday"
	case 5:
		result = " today Friday"
	default:
		result = " today is weekend"
	}
	fmt.Println(result)
}

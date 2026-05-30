package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var result float64 = calculate("445", "10", "*")
	fmt.Println(result)

}

func calculate(input1 string, input2 string, operation string) float64 {
	val1 := convertInputToValue(input1)
	val2 := convertInputToValue(input2)
	var result float64 = 0
	switch operation {
	case "+":
		result = addValues(val1, val2)
	case "*":
		result = multiplyValues(val1, val2)
	case "-":
		result = subtractValues(val1, val2)
	case "/":
		result = divideValues(val1, val2)
	default:
		fmt.Println("invalid math operation")
	}
	return result
}

func convertInputToValue(input string) float64 {
	val, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("Error converting input to float")
	}
	return val
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your age")
	var str, _ = reader.ReadString('\n')
	age, _ := strconv.ParseInt(strings.TrimSpace(str), 0, 64)
	var pointer *int64 = &age
	fmt.Printf("Pointer is %d\n", *pointer)
	///
	fmt.Println("Enter your salary")
	str, _ = reader.ReadString('\n')
	salary, _ := strconv.ParseFloat(strings.TrimSpace(str), 64)
	pointer2 := &salary
	fmt.Printf("Pointer is %f\n", *pointer2)
	*pointer2 += 24445532
	fmt.Printf("Pointer is %f\n", *pointer2)
	fmt.Printf("your salary is %f\n", salary)
}

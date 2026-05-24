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
	fmt.Print("Enter your name: ")
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)

	fmt.Print("Enter your price: ")
	str, err = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Printf("Invalid Operation: %v\n and its type is %T", err, str)
	}
	fmt.Println(f)
}

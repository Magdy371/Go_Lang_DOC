package main

import (
	"fmt"
	"time"
)

func main() {
	go syaHello()
	fmt.Println("Hello from main function")
	time.Sleep(5 * time.Second)
	fmt.Println("All done")
}
func syaHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello World")
}

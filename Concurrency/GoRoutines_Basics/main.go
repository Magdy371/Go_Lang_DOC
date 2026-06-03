package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello()
	time.Sleep(1 * time.Second)
	sayGoodBye()
}

func sayHello() {
	fmt.Println("Hello To Go")
}
func sayGoodBye() {
	fmt.Println("Good Bye !!")
}

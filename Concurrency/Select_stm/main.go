package main

import "fmt"

func main() {
	fmt.Println("Main started")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go selectWork(ch1, ch2)

	// Send messages
	ch1 <- "Hello from main to ch1"
	ch2 <- "Hello from main to ch2"

	msg1 := <-ch1
	msg2 := <-ch2
	fmt.Println("msg1:", msg1, "msg2:", msg2)
	fmt.Println("Main finished")
}

func selectWork(ch1, ch2 chan string) {
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println("ch1: message 1 received", msg1)
			ch1 <- "response from selectWork"
		case msg2 := <-ch2:
			fmt.Println("ch2: message 2 received", msg2)
			ch2 <- "response from selectWork"
		default:
			// Sleep or remove default to avoid busy loop
			fmt.Println("no message received")
		}
	}
}

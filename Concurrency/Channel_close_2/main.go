package main

import "fmt"

func main() {
	var greetings []string = []string{"Hello", "Hi", "Hey"}
	var messages chan string = make(chan string, len(greetings))
	fmt.Println("main started")
	go greeting(messages, greetings)

	// for {
	// 	msg, ok := <-messages
	// 	if !ok {
	// 		fmt.Println("main finished")
	// 		return
	// 	}
	// 	fmt.Println(msg)
	// }
	//
	for msg := range messages {
		fmt.Println(msg)
	}
	fmt.Println("main finished")
}
func greeting(message chan<- string, array []string) {
	for err, v := range array {
		if err < 0 {
			continue
		}
		message <- v
	}
	close(message)
	fmt.Println("messages recieved")
}

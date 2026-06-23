package main

import (
	"fmt"
	"time"
)

func main() {
	/**
	 * Channel can be buffered e.g. make(chan string, 1) means recieve on argument
	 * Buffered channels is async
	 * Non-buffered channels are sync operation e.g make(chan string)
	**/
	var message chan string = make(chan string, 1)
	go gereeting(message)
	time.Sleep(5 * time.Second)
	fmt.Println("Main Ready")
	var reciever string = <-message
	fmt.Println(reciever)
	time.Sleep(2 * time.Second)
	fmt.Println("Main Done")
}

func gereeting(message chan<- string) {
	fmt.Println("Hello World")
	message <- "agent name is magdy"
	fmt.Println("Message encapsulated")
}

/***
 *
 * Channels are used to communicate between goroutines
 * send only channel: chan<- T  (T is generic can be any type), unidirectional
 * recieve only channel: <-chan T, unidirectional
 * bidirectional channel: chan T implicitly cast to uniderectional channel but opposite is imposible
 * Unidirectional channels are used to enforce type safety and prevent accidental data flow in the wrong direction.
 *
***/

/***
 * Channel Close Operations:
 * Closing a channel is done using the close() function.
 * means no more values will be sent on the channel. e.g close(message)
 * we can close bidirectional channel and send only channel
 * we cannot close recieve only channel there will cause a panic
***/

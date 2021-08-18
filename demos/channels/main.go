package main

import "fmt"

func main() {
	messages := make(chan string)

	go func () { messages <- "ping" }()

	go func() { messages <- "pong" } ()
	msg := <-messages

	fmt.Println(msg)

	msg1 := <- messages
	fmt.Println(msg1)
}

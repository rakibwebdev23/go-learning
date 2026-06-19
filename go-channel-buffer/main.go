package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string, 1) // Buffered channel (capacity = 1)

	go func() {
		fmt.Println("Sending message to messages channel")

		messages <- "Hello from message channel"

		fmt.Println("Message sent successfully")
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("About to get message from channel")

	msg := <-messages

	fmt.Println(msg)
}
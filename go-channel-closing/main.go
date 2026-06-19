package main

import (
	"fmt"
	"sync"
)

func main() {
	messages := make(chan int, 3) // Buffered channel for string messages
	var wg sync.WaitGroup // eta goroutine track kore

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			m, ok := <-messages // Receive a message from the messages channel
			if ok {
				fmt.Println("Received message:", m)
			} else {
				fmt.Println("Channel closed")
				return
			}
		}
	}(&wg)

	for i := 1; i <= 3; i++{
		messages <- i;
		fmt.Println("Sending this message");
	}

	close(messages)
	wg.Wait()
}

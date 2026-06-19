package main

import (
	"fmt"
	"time"
)

type User struct {
	Name string
	Age int
}

func main (){
	messages := make (chan string) //Unbuffered channel for string messages
	users := make(chan User) //Unbuffered channel for User struct

	go func (){
		fmt.Println("Sending message to messages channel")
		messages <- "Hello from message channel" // Send a message to the messages channel
	}();

	go func (){
		users <- User{Name: "Alice", Age: 30}
	}();

	time.Sleep(1 * time.Second)
	fmt.Println("About to get message from channel")
	mgg := <-messages // Receive a message from the messages channel and store it in mgg
	fmt.Println(mgg)

	u := <-users // Receive a user from the users channel and store it in u
	fmt.Println("Name:",u.Name, "Age:",u.Age)

}
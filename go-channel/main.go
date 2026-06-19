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
	messages := make (chan string)
	users := make(chan User)

	go func (){
		fmt.Println("Sending message to messages channel")
		messages <- "Hello from message channel"
	}();

	go func (){
		users <- User{Name: "Alice", Age: 30}
	}();

	time.Sleep(1 * time.Second)
	fmt.Println("About to get message from channel")
	mgg := <-messages
	fmt.Println(mgg)

	u := <-users
	fmt.Println("Name:",u.Name, "Age:",u.Age)

}
package main 

import "fmt"

type Person struct {
	name string
	age  int
}

func userInfo (data Person){
	fmt.Println("Name:", data.name)
	fmt.Println("Age:", data.age)
}

func main() {
	user1:= Person{name: "Alice", age: 30}
	user2:= Person{name: "Bob", age: 25}
	userInfo(user1)
	userInfo(user2)
}
package main

import "fmt"

type User struct {
	Name string
	Age  int
	Hobby string
}

func main() {
	user :=User{
		Name: "Rakib",
		Age:  28,
		Hobby: "Reading",
	}

	user2 := User{
		Name: "Rakib",
		Age:  28,
		Hobby: "Reading",
	}

	user3 := User{
		Name: "Rakib",
		Age:  28,
		Hobby: "Reading",
	}

	fmt.Println(user.Name);
	fmt.Println(user);
	fmt.Println(user2.Name);
	fmt.Println(user2);
	fmt.Println(user3.Name);
	fmt.Println(user3);
}

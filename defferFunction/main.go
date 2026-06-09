package main

import "fmt"

// defer function agee call kora hoy jeta last e ghotbe
func sayGoodbye() {
	fmt.Println("Goodbye!")
}

func main() {
	defer sayGoodbye()

	fmt.Println("Hello")
	fmt.Println("Learning Go")
}
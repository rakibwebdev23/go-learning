package main

import (
	"fmt"
	"strings"
)

func main() {
	numbers := []int{10, 20, 30, 40, 50}		
	numbers = append(numbers, 60) // add an element to the slice
	
	fmt.Println(numbers)
	fmt.Println(numbers[0]) // first element

	items:= []string{"apple", "banana", "cherry"}
	items = append(items, "watermellon")
	items = append(items, "grape")
	items = append(items, "orange")

	fmt.Println(strings.Join(items, ", "))
	fmt.Println(items[2]) // third element
}
package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for _, v := range nums {
		total += v
	}

	return total
}

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(10, 20))
	fmt.Println(sum(5))
}

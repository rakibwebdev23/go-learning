// package main

// import "fmt"

// func factorial (n int) int {
// 	if n <= 0 {
// 		return 1;
// 	}
// 	return n * factorial(n-1)
// }

// func main() {
// 	fmt.Println(factorial(5))
// }

package main

import "fmt"

func factorial(n int) int {
	fmt.Println("Calling factorial(", n, ")")

	if n == 0 {
		fmt.Println("Base case reached, return 1")
		return 1
	}

	result := n * factorial(n-1)

	fmt.Println("Returning:", n, "* =", result)

	return result
}

func main() {
	fmt.Println("Final Result:", factorial(5))
}
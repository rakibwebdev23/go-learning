package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		// fmt.Println("for loop output", i)
	}

	//while loop
	k := 3

	for k > 0 {
		// fmt.Println("while loop output", k)
		k-- // this is decrement statement otherwise it will be infinite loop
	}

	// infinite loop
	counter := 0

	for {
		// fmt.Println("infinite loop output", counter)
		counter++
		if counter >= 10 {
			break
		}
	}

	//Skipping loop iterations
	for j := 0; j <= 10; j++ {
		if j%2 == 0 {
			continue
		}
		// fmt.Println("Skipping output", j)
	}

	//Nested loops
	for m := 1; m <= 3; m++ {
		for n := 1; n <= 2; n++ {
			// fmt.Printf("Nested loop output: m=%d, n=%d\n", m, n)
		}
	}
	// loop for array
	// arr := []string{"c", "c++", "java", "python"}
	// for index, value := range arr {
	// 	fmt.Println("array index & value output", index, value)
	// }

	// or, 
	// arr := []string{"c", "c++", "java", "python"}
	// for index,_ := range arr {
	// 	fmt.Println("array index output", index)
	// }

	arr := []string{"c", "c++", "java", "python"}
	for _, value := range arr {
		fmt.Println("array value output", value)
	}

	// loop for array with nested loop
	// arr := []string{"c", "c++", "java", "python"}
	// for index,_ := range arr {
	// 	for _, value := range arr {
	// 		fmt.Println("array index & value output", index, value)
	// 	}
	// }
}

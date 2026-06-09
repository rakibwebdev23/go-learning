// package main

// import "fmt"

// func main() {
    // arr:= [6]string{"This", "is", "a", "Go", "interview", "question"}
    // fmt.Println(arr)

	// s:= arr[1:4]
	// fmt.Println(s)

	// s1:= arr[1:2]
	// fmt.Println(s1)

	// fmt.Println(len(s1));
	// fmt.Println(cap(s1));
	// s1=append(s1, "not js only go")
	// fmt.Println(s1)

	// slice make function 

// 	arr:= make([]int,3, 10) //3 is the length of the slice, 10 is the capacity
// 	fmt.Println(arr)

// }


// slice veriatic function 
package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(cap(nums))
}

func main() {
	sum(1, 2, 5, 6, 9)
}





// slice all function 
// 1. slice from existing array
// 2. slice from existing slice
// 3. slice literal
// 4. slice using make function
package main

import "fmt"

func calculate() (result int) {
	fmt.Println("first print", result)

	defer func() {
		result = result + 10
		fmt.Println("second print", result)
	}()

	result = 5
	return
}

func main() {
	res := calculate()
	fmt.Println("final result:", res)
}




// package main

// import "fmt"

// func deferFunc (a int, b int) int {
// 	i:= a + b
// 	fmt.Println("this is first output", i)
// 	defer fmt.Println("this is second output", i)
// 	fmt.Println("this is third output", i)
// return  i;
// }

// func main () {
// 	deferFunc(2, 3)
// }

// func sum (a int64, b int64) (i int64) {
// 	i = a + b
// 	return i
// }

// func main (){
// 	res := sum (3, 5)
// 	fmt.Println(res)
// }
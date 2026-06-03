// package main

// import "fmt"

// func modifyPointer (val int){
// 	val = val * 10
// 	fmt.Printf("Value inside function: %d\n", val)
// 	fmt.Printf("Age address: %p\n", &val)
// }

// func main() {

// 	age:= 30
// 	fmt.Printf("My age is: %d\n", &age)
// 	fmt.Printf("My age is: %d\n", age)
// 	modifyPointer(age)
// }


// package main

// import "fmt"

// func main (){
// 	x := 10
// 	p := &x
// 	q := *p
// 	*p = 20
// 	fmt.Println(x)
// 	fmt.Println(p)
// 	fmt.Println(q)
// }


package main

import "fmt"

func Print(numbers *[3]int){
	fmt.Println(numbers)
}

func main(){
	numbers := [3]int{1, 2, 3}
	Print(&numbers);
}
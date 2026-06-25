package main

import "fmt"

func calculate() (result int) {
	fmt.Println("calculate first output", result)

	sum := func ()  {
		result = result + 30;
		fmt.Println("calculate second output", result);
	}

	defer sum()

	result = 10;
	fmt.Println("calculate third output", result);

	return result;
}

func calc () int{
	result := 0;
	fmt.Println("first output");

	show := func ()  {
		result = result + 20;
		fmt.Println("second output", result);
	}

	defer show()

	result = 5;
	fmt.Println("third output", result)

	return result;
}

func main() {
    a := calculate()
	fmt.Println("final result:", a)

	b := calc()
	fmt.Println("final result",b);
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
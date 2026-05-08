package main

import fmt "fmt"

const HOST string = "localhost"
func main () {
	fmt.Println("Hello, World!")
	fmt.Println("Host:", HOST)

	const PORT int = 8080
	fmt.Println("Port:", PORT)

	const AppName string = "MyApp"
	fmt.Println("App Name:", AppName)

	appName := "GO_APP"
	fmt.Println("App Name:", appName)

	const PI float32 = 3.1422221213333
	fmt.Println("PI:", PI)

	const isProduction bool = false
	fmt.Println("Is Production:", isProduction)

	const float64 float64 = 3.141592653589793
	fmt.Println("Float64:", float64)
}
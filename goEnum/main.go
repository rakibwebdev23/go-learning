package main

import "fmt"

const (
	Saturday = 1

	Sunday = 2

	Monday = 3

	Tuesday = 4

	Wednesday = 5

	Thursday = 6

	Friday = 7	
)

type logLevel int

const (
	Debug logLevel = iota
	Info
	Warning
	Error
)
func main() {
	fmt.Println(Saturday)
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Debug)
	fmt.Println(Info)
	fmt.Println(Warning)
	fmt.Println(Error)
}
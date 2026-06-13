package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Display() {
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}

type Employee struct {
	Person
	Salary float64
}

func main() {
	emp := Employee{
		Person: Person{
			Name: "Rakib",
			Age:  25,
		},
		Salary: 50000,
	}

	emp.Display()
	fmt.Println("Salary:", emp.Salary)
}
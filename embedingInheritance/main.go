package main

import "fmt"

type Address struct {
	City    string
	Country string
}

type ContactInfo struct {
	Email string
	Phone string
}

type Person struct {
	Name        string
	Age         int
	Address
	ContactInfo
}

func (p Person) GetProfile() {
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("City:", p.City)
	fmt.Println("Country:", p.Country)
	fmt.Println("Email:", p.Email)
	fmt.Println("Phone:", p.Phone)
}

func main() {
	person := Person{
		Name: "Rakib",
		Age:  30,
		Address: Address{
			City:    "Dhaka",
			Country: "Bangladesh",
		},
		ContactInfo: ContactInfo{
			Email: "rakib@example.com",
			Phone: "01712345678",
		},
	}
	person.GetProfile();
}
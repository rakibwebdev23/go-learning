package main

import "fmt"

func main() {
	age := 20
	if age < 18 {
		fmt.Println("Minor")
	} else {
		fmt.Println("Adult")
	}

	// if-else if-else
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else if score >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}

	// map condiont 
	useAccess := map[string]bool{
		"admin": true,
		"user":  false,
	}

	role := "admin"
	if hasAccess, exists := useAccess[role]; exists && hasAccess {
		fmt.Printf("%s has access\n", role)
	} else {
		fmt.Printf("%s does not have access\n", role)
	}
}

package main

import (
	"fmt"
)

func main() {
	studentsGrades := map[string]int{
		"Alice": 85,
		"Bob":   92,
		"Charlie": 78,
	}
	fmt.Println(studentsGrades)
	fmt.Println(studentsGrades["Alice"]) // access Alice's grade

	// Adding a new student
	studentsGrades["David"] = 88
	fmt.Println(studentsGrades)

	// Updating an existing student's grade
	studentsGrades["Bob"] = 95
	fmt.Println(studentsGrades)

	// Deleting a student
	delete(studentsGrades, "Charlie")
	fmt.Println(studentsGrades)
}
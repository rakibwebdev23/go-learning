package main

import (
	"fmt"
	"strings"
)

type MathError struct {
	Operation string
	InputA    int
	InputB    int
	Message   string
}

func (e MathError) Error() string {
	var inputs []string

	if e.Operation == "Division" {
		inputs = append(inputs, fmt.Sprintf("a=%d", e.InputA))
		inputs = append(inputs, fmt.Sprintf("b=%d", e.InputB))
	}

	return fmt.Sprintf(
		"Math error in %s (%s): %s",
		e.Operation,
		strings.Join(inputs, ", "),
		e.Message,
	)
}

func main() {
	err := MathError{
		Operation: "Division",
		InputA:    10,
		InputB:    0,
		Message:   "cannot divide by zero",
	}

	fmt.Println(err)
}

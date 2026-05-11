package main

func main() {
	var a int = 10
	switch a {
	case 1:
		println("a is 1")
	case 10:
		println("a is 10")
	default:
		println("a is not 1 or 10")
	}

	var b string = "hello"
	switch b {
	case "hello":
		println("b is hello")
	case "world":
		println("b is world")
	default:
		println("b is not hello or world")
	}

	var day string = "Monday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		println("It's a weekday")
	case "Saturday", "Sunday":
		println("It's a weekend")
	default:
		println("Invalid day")
	}

	checkType := func(i interface{}) {
		switch v := i.(type) {
		case int:
			println("i is an int:", v)
		case string:
			println("i is a string:", v)
		default:
			println("i is of unknown type")
		}
	}

	checkType(42)
	checkType("Go")
	checkType(3.14)
}
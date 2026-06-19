package main

import (
	"fmt"
	"regexp"
)

func main() {

	text := "Hello World! Welcome to Go 2025. Email: test@gmail.com"

	// ====================================================
	// 1. Compile Regex
	// ====================================================

	regGo, err := regexp.Compile(`Go`)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("MatchString:", regGo.MatchString(text))

	// ====================================================
	// 2. MustCompile
	// ====================================================

	regWorld := regexp.MustCompile(`World`)
	fmt.Println("MustCompile Match:", regWorld.MatchString(text))

	// ====================================================
	// 3. FindString
	// ====================================================

	fmt.Println("FindString:",
		regGo.FindString(text))

	// ====================================================
	// 4. FindStringIndex
	// ====================================================

	fmt.Println("FindStringIndex:",
		regGo.FindStringIndex(text))

	// ====================================================
	// 5. FindAllString
	// ====================================================

	data := "Go Java Go Python Go Rust"

	reg := regexp.MustCompile(`Go`)

	fmt.Println("FindAllString:",
		reg.FindAllString(data, -1))

	// ====================================================
	// 6. FindAllStringIndex
	// ====================================================

	fmt.Println("FindAllStringIndex:",
		reg.FindAllStringIndex(data, -1))

	// ====================================================
	// 7. ReplaceAllString
	// ====================================================

	fmt.Println("ReplaceAllString:",
		reg.ReplaceAllString(data, "Golang"))

	// ====================================================
	// 8. Split
	// ====================================================

	regSpace := regexp.MustCompile(`\s+`)

	fmt.Println("Split:")
	fmt.Println(regSpace.Split(data, -1))

	// ====================================================
	// 9. Digit Search
	// ====================================================

	regDigit := regexp.MustCompile(`[0-9]+`)

	fmt.Println("Digits Found:")
	fmt.Println(regDigit.FindAllString(text, -1))

	// ====================================================
	// 10. Email Match
	// ====================================================

	regEmail := regexp.MustCompile(
		`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`,
	)

	fmt.Println("Email Found:")
	fmt.Println(regEmail.FindString(text))

	// ====================================================
	// 11. Starts With
	// ====================================================

	regStart := regexp.MustCompile(`^Hello`)

	fmt.Println("Starts With Hello:",
		regStart.MatchString(text))

	// ====================================================
	// 12. Ends With
	// ====================================================

	regEnd := regexp.MustCompile(`gmail\.com$`)

	fmt.Println("Ends With gmail.com:",
		regEnd.MatchString("test@gmail.com"))

	// ====================================================
	// 13. Word Boundary
	// ====================================================

	regWord := regexp.MustCompile(`\bGo\b`)

	fmt.Println("Word Boundary:")
	fmt.Println(regWord.FindAllString(
		"Go Golang GoLang Go", -1))

	// ====================================================
	// 14. FindStringSubmatch
	// ====================================================

	user := "Name: John Age: 25"

	regUser := regexp.MustCompile(
		`Name:\s+(\w+)\s+Age:\s+(\d+)`,
	)

	match := regUser.FindStringSubmatch(user)

	fmt.Println("SubMatches:")
	fmt.Println(match)

	if len(match) > 2 {
		fmt.Println("Name:", match[1])
		fmt.Println("Age :", match[2])
	}

	// ====================================================
	// 15. Match Multiple Words
	// ====================================================

	regLang := regexp.MustCompile(`Go|Java|Python`)

	fmt.Println("Languages:")
	fmt.Println(regLang.FindAllString(
		data, -1))

}
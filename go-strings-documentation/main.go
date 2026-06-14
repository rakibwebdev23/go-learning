package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "  Hello, Go World! hello go  "

	// 1. Original string
	fmt.Println("Original:", s)

	// 2. TrimSpace (remove leading & trailing spaces)
	fmt.Println("TrimSpace:", strings.TrimSpace(s))

	// 3. ToLower
	fmt.Println("ToLower:", strings.ToLower(s))

	// 4. ToUpper
	fmt.Println("ToUpper:", strings.ToUpper(s))

	// 5. Title (deprecated but still works)
	fmt.Println("Title:", strings.Title(s))

	// 6. Contains
	fmt.Println("Contains 'Go':", strings.Contains(s, "Go"))

	// 7. ContainsAny
	fmt.Println("ContainsAny 'xyzGo':", strings.ContainsAny(s, "xyzGo"))

	// 8. Count
	fmt.Println("Count 'o':", strings.Count(s, "o"))

	// 9. Index
	fmt.Println("Index of 'Go':", strings.Index(s, "Go"))

	// 10. LastIndex
	fmt.Println("LastIndex of 'o':", strings.LastIndex(s, "o"))

	// 11. Replace
	fmt.Println("Replace first hello:", strings.Replace(s, "hello", "hi", 1))

	// 12. ReplaceAll
	fmt.Println("ReplaceAll:", strings.ReplaceAll(s, "o", "0"))

	// 13. Trim (custom chars remove)
	fmt.Println("Trim:", strings.Trim(s, " h!"))

	// 14. TrimLeft
	fmt.Println("TrimLeft:", strings.TrimLeft(s, " "))

	// 15. TrimRight
	fmt.Println("TrimRight:", strings.TrimRight(s, " "))

	// 16. Split
	parts := strings.Split(s, " ")
	fmt.Println("Split:", parts)

	// 17. Join
	fmt.Println("Join:", strings.Join(parts, "-"))

	// 18. Fields (auto split by space)
	fmt.Println("Fields:", strings.Fields(s))

	// 19. HasPrefix
	fmt.Println("HasPrefix '  He':", strings.HasPrefix(s, "  He"))

	// 20. HasSuffix
	fmt.Println("HasSuffix 'go  ':", strings.HasSuffix(s, "go  "))

	// 21. Repeat
	fmt.Println("Repeat:", strings.Repeat("Go ", 3))

	// 22. Clone
	cloned := strings.Clone(s)
	fmt.Println("Clone:", cloned)

	// 23. Builder (efficient string build)
	var b strings.Builder
	b.WriteString("Hello ")
	b.WriteString("from Builder")
	fmt.Println("Builder:", b.String())

	// 24. Equal
	fmt.Println("Equal:", strings.EqualFold("Go", "go")) // case-insensitive compare

	// 25. Map (transform each rune)
	result := strings.Map(func(r rune) rune {
		if r == 'o' {
			return '0'
		}
		return r
	}, s)

	fmt.Println("Map result:", result)
}
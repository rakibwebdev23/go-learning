package main

import (
	"fmt"
	"unicode"
)

// func main (){
// 	username:= "admin"
// 	fmt.Println(username[0])
// }

func main() {
	data := []rune{'然', '文', '字'}
	for _, v:= range data {
		fmt.Println(string(v), unicode.IsLetter(v))
	}
}
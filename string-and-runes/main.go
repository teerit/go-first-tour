package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	const s = "สวัสดี"
	const t = "hello"

	fmt.Println("Len: ", len(s))
	fmt.Println("Len(eng): ", len(t))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x", s[i])
	}

	fmt.Println()
	fmt.Println("Rune count: ", utf8.RuneCountInString(s))

	fmt.Println()
	fmt.Println("Rune count (eng): ", utf8.RuneCountInString(t))
}

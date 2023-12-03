package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Person struct {
	name string
	age  int
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	x, y := 5, 7
	result1 := cmp.Compare[int](x, y)
	fmt.Println(result1) // output: -1

	result2 := cmp.Less[int](x, y)
	fmt.Println(result2) // output: true

	a, b := "xx", "zy"
	result3 := cmp.Compare[string](a, b)
	fmt.Println(result3)

	// sort string by their length
	fruits := []string{"peach", "banana", "kiwi"}
	lenCmp := func(a, b string) int {
		return cmp.Compare[int](len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)
	fmt.Println(factorial(5))
}

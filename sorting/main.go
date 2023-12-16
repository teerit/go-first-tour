package main

import (
	"fmt"
	"slices"
)

func main() {

	strs := []string{"c", "b", "a"}
	slices.Sort(strs)
	fmt.Println("Sort strings: ", strs)

	ints := []int{7, 2, 3}
	slices.Sort(ints)
	fmt.Println("Sort integers: ", ints)

	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}

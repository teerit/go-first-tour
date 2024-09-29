package main

import "fmt"

func MultplicationTable(size int) [][]int {
	var result [][]int
	multiplier := 1
	for i := 1; i <= size; i++ {
		temp := []int{}
		for j := 1; j <= size; j++ {
			temp = append(temp, multiplier*j)
		}
		result = append(result, temp)
		multiplier++
	}

	return result
}

func main() {
	fmt.Println(MultplicationTable(3))
}

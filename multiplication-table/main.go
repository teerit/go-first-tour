package main

import "fmt"

func MultiplicationTable(size int) [][]int {
	matrix := make([][]int, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
		for y := 0; y < size; y++ {
			matrix[i][y] = (i + 1) * (y + 1)
		}
	}
	return matrix
}

func main() {
	fmt.Println(MultiplicationTable(3))
}

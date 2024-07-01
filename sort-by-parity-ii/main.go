package main

import "fmt"

func sortArrayByParityII(nums []int) []int {
	size := len(nums)
	var result = make([]int, size)

	e := 0
	o := 1
	for _, v := range nums {
		if v%2 == 0 {
			result[e] = v
			e += 2
		}

		if v%2 != 0 {
			result[o] = v
			o += 2
		}
	}

	return result
}

func main() {
	nums := []int{4, 2, 5, 7}
	fmt.Println(sortArrayByParityII(nums))
}

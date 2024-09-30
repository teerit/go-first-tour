package main

import "fmt"

func DigitalRoot(n int) int {
	total := Sum(n)
	for total > 10 {
		total = Sum(total)
	}
	return total
}

func Sum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n = n / 10
	}
	return sum
}

func main() {
	fmt.Println(DigitalRoot(942))
}

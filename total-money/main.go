package main

import "fmt"

func totalMoney(n int) int {
	if n < 7 {
		return sumSeq(1, n)
	} else {
		m := n % 7
		w := int(n / 7)
		sum := sumWeek(w)
		return sum + sumSeq(w+1, w+m)
	}
}

func sumWeek(n int) int {
	return (28 * n) + 7*(n-1)
}

func sumSeq(s, n int) int {
	return ((n - s + 1) * (n + s)) / 2
}

func main() {
	fmt.Println(totalMoney(10))
}

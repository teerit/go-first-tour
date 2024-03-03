package main

import "fmt"

func createCounter() func() int {
	count := 0
	increment := func() int {
		count++
		return count
	}
	return increment
}

func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return x
	}
}

func main() {
	counter1 := createCounter()
	counter2 := createCounter()

	fmt.Println(counter1())
	fmt.Println(counter1())
	fmt.Println(counter1())
	fmt.Println(counter1())
	fmt.Println(counter2())
	fmt.Println(counter2())
	fmt.Println(counter2())

	f := fibonacci()
	for i := f(); i < 3; i = f() {
		fmt.Println(i)
	}
}

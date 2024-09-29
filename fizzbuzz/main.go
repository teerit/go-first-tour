package main

import "fmt"

func main() {
	fizzBuzz(15)
}

func fizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 {
			if i%3 == 0 && i%5 == 0 {
				fmt.Println("FizzBuzz")
				continue
			}

			if i%3 == 0 {
				fmt.Println("Fizz")
				continue
			}

			if i%5 == 0 {
				fmt.Println("Buzz")
				continue
			}
		} else {
			fmt.Println(i)
		}
	}
}

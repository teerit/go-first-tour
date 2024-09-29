package main

import "fmt"

/*
1  => 1
2  => 2
3  => FooFoo (divisible by 3, contains 3)
4  => 4
5  => BarBar (divisible by 5, contains 5)
6  => Foo (divisible by 3)
7  => QixQix (divisible by 7, contains 7)
8  => 8
9  => Foo
10 => Bar
13 => Foo
15 => FooBarBar (divisible by 3, divisible by 5, contains 5)
21 => FooQix
33 => FooFooFoo (divisible by 3, contains two 3)
51 => FooBar
53 => BarFoo
*/

func main() {
	fooBarQix(10)
}

func fooBarQix(n int) {
	for i := 1; i <= n; i++ {
		countThree := containNumber(i, 3)
		countFive := containNumber(i, 5)
		// countSeven := containNumber(i, 7)
		if i%3 == 0 || i%5 == 0 {

			if i%3 == 0 {
				s := "Foo"
				for j := 0; j < countThree; j++ {
					s += "Foo"
				}
				fmt.Println(s)
				continue
			}

			if i%5 == 0 {
				s := "Bar"
				for j := 0; j < countFive; j++ {
					s += "Bar"
				}
				fmt.Println(s)
				continue
			}

		} else {
			fmt.Println(i)
		}
	}
}

func containNumber(i int, c int) int {
	count := 0
	for i > 0 {
		n := i % 10
		i = i / 10
		if n == c {
			count++
		}
	}
	return count
}

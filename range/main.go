package main

import "fmt"

func main() {

	sum := 0
	nums := []int{2, 3, 4}
	for _, num := range nums {
		sum += num
	}

	fmt.Println(sum)

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key: ", k)
	}
}

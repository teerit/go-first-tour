package main

import "fmt"

// typed collections of fields
type person struct {
	name string
	age  int
}

func NewPerson(name string) *person {
	p := person{name: name}
	// access struct field with dot
	p.age = 42
	return &p
}

func main() {
	john := NewPerson("John")
	fmt.Println("Name: ", john.name)
	fmt.Println("John age: ", john.age)
}

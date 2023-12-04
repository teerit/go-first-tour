package main

import "fmt"

type User struct{}
type UserRepository interface {
	Save(user User) error
}

// MySQLUserRepositopry implements UserRepository
type MySQLUserRepository struct{}

func (m *MySQLUserRepository) Save(user User) error {
	return nil
}

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t *T) M() {
	fmt.Println(t.S)
}

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.height + 2*r.width
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	var userRepo UserRepository
	userRepo = &MySQLUserRepository{}
	userRepo.Save(User{})

	var i I = &T{"hello"}
	i.M()

	r := rect{width: 3, height: 5}
	measure(r)
}

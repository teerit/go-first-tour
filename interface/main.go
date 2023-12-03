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

func main() {
	var userRepo UserRepository
	userRepo = &MySQLUserRepository{}
	userRepo.Save(User{})

	var i I = &T{"hello"}
	i.M()
}

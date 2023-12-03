package main

import (
	"fmt"
	"strings"
)

type Profile struct {
	Name string
	Age  int
}

func upperProfileName(p *Profile) {
	p.Name = strings.ToUpper(p.Name)
}

func main() {
	a := 10
	b := 20
	c := add(a, b)
	fmt.Println(c)

	name := "teerit"
	upperAllLetter(&name)
	fmt.Println(name)

	p := &Profile{
		Name: "Teerit",
		Age:  26,
	}
	upperProfileName(p)
	fmt.Println(p.Name)
}

func upperAllLetter(str *string) {
	*str = strings.ToUpper(*str)
}

// ทุกการกำหนดค่าของ Go คือ การ copy
// จากด้านขวาไปทางด้านซ้ายทั้งหมด
func add(a, b int) int {
	c := a + b
	return c
}

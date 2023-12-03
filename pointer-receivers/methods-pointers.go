package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// value receiver
// copy value และสร้างตัวแปรใหม่เพื่อเก็บค่า
// และจะไม่มี side effect ไปยัง memory address
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// pointer receiver
// สร้างตัวแปรเก็บ Address เพื่อทำการ Deference
// ตัวแปรเดิมเสมอ วีธีการนี้จะใช้ตำแหน่งของ
// memory addres เดิมในการอ่านและ เขียน
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// pointer indirection
func main() {
	v := Vertex{3, 4}
	// Go interprets the statement v.Scale(5)
	// as (&v).Scale(5) since the `Scale` method
	// has a pointer receiver.
	v.Scale(2)
	// ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	// ScaleFunc(p, 8)

	fmt.Println(v, p)
}

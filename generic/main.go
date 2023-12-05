package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) RemoveTail() {

	for e := lst.head; e != nil; e = e.next {
		if e.next == lst.tail {
			lst.tail = e
			lst.tail.next = nil
		}
	}

}

func (lst *List[T]) RemoveHead() {
	lst.head = lst.head.next
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}
	fmt.Println("keys: ", MapKeys[int, string](m))

	lst := List[int]{}
	lst.Push(1)
	lst.Push(2)
	lst.Push(3)
	fmt.Println("list: ", lst.GetAll())

	lst.RemoveHead()
	fmt.Println("list: ", lst.GetAll())
}

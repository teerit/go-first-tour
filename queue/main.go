package main

import "fmt"

type MyQueue struct {
	Max int
	T   int
	S   [10]int
}

func Constructor() MyQueue {
	return MyQueue{
		Max: 10,
		T:   -1,
		S:   [10]int{},
	}
}

func (this *MyQueue) Push(x int) {
	if this.T >= 0 {
		this.S[this.T+1] = x
	} else {
		this.S[0] = x
	}
	this.T = this.T + 1
}

func (this *MyQueue) Pop() int {
	tmp := this.S[0]
	this.S[0] = 0
	return tmp
}

func (this *MyQueue) Peek() int {
	return this.S[0]
}

func (this *MyQueue) Empty() bool {
	for _, v := range this.S {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	// ["MyQueue","push","push","push","push","pop","push","pop","pop","pop","pop"]
	// [[],[1],[2],[3],[4],[],[5],[],[],[],[]]
	q := Constructor()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)

	q.Pop()
	fmt.Println(q.S)

	q.Push(5)
	fmt.Println(q.S)
	q.Pop()
	q.Pop()
	q.Pop()
	q.Pop()
	fmt.Println(q.S)
}

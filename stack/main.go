package main

import "fmt"

type MyStack struct {
	Max int
	T   int
	S   [10]int
}

func Constructor() MyStack {
	return MyStack{
		Max: 10,
		T:   -1,
		S:   [10]int{},
	}
}

func (this *MyStack) Push(x int) {
	if this.T >= 0 {
		this.S[this.T+1] = x
	} else {
		this.S[0] = x
	}
	this.T = this.T + 1
}

func (this *MyStack) Pop() int {
	tmp := this.S[this.T]
	this.S[this.T] = 0
	this.T--
	return tmp
}

func (this *MyStack) Top() int {
	return this.S[this.T]
}

func (this *MyStack) Empty() bool {
	for _, v := range this.S {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)

	obj.Pop()

	fmt.Println(obj.Top())
	fmt.Println(obj.S)
	fmt.Println("Is stack empty? ", obj.Empty())
}

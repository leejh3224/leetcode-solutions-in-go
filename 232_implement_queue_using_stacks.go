package main

import "fmt"

// stack implementation

// Stack ...
type Stack []int

// Push ...
func (s *Stack) Push(x int) { *s = append(*s, x) }

// Pop ...
func (s *Stack) Pop() int {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}

// Peek ...
func (s *Stack) Peek() int { return (*s)[len(*s)-1] }

// Empty ...
func (s *Stack) Empty() bool {
	return len(*s) == 0
}

// MyQueue ...
type MyQueue struct {
	in, out Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.in.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.out.Empty() {
		for !this.in.Empty() {
			this.out.Push(this.in.Pop())
		}
	}
	return this.out.Pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.out.Empty() {
		for !this.in.Empty() {
			this.out.Push(this.in.Pop())
		}
	}
	return this.out.Peek()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.in.Empty() && this.out.Empty()
}

func main() {
	var q = Constructor()
	q.Push(1)
	q.Push(2)
	fmt.Printf("%v %v\n", q.Peek(), q)
	q.Push(3)
	fmt.Printf("%v %v\n", q.Peek(), q)
	fmt.Printf("%v %v\n", q.Pop(), q)
	fmt.Printf("%v %v\n", q.Pop(), q)
}

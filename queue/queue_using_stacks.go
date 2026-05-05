package queue

import "errors"

// ErrStackEmpty occurs when popping an empty stack.
var ErrStackEmpty = errors.New("empty stack")

type (
	// UsingStacks is a queue that is made using two stacks.
	UsingStacks struct {
		stack1 Stack
		stack2 Stack
	}
	// Stack is a stack of integers.
	Stack struct {
		stack []int
	}
)

// enqueue solves the problem in O(1) time and O(n) space.
func (usingStacks *UsingStacks) enqueue(n int) {
	panic("not implemented")

}

// dequeue solves the problem in O(1) time and O(n) space.
func (usingStacks *UsingStacks) dequeue() int {
	panic("not implemented")

}

func (stack *Stack) push(element int) {
	panic("not implemented")

}

func (stack *Stack) pop() int {
	panic("not implemented")

}

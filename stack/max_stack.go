package stack

import "errors"

// MaxStack is a stack that can return the maximum of the integers pushed.
type MaxStack struct {
	stack1, stack2 []int
}

// ErrEmptyStack occurs when a trying to pop a stack that is empty.
var ErrEmptyStack = errors.New("stack is empty")

// Max solves the problem in O(1) time and O(n) space.
func (maxStack *MaxStack) Max() int {
	panic("not implemented")

}

// Push adds an integer to the stack.
func (maxStack *MaxStack) Push(i int) {
	panic("not implemented")

}

// Pop returns the last inserted integer.
func (maxStack *MaxStack) Pop() (int, error) {
	panic("not implemented")

}

func max(a, b int) int {
	panic("not implemented")

}

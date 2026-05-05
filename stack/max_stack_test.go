package stack

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"errors"
)

/*
TestMaxStack tests solution(s) with the following signature and problem description:

	func (maxStack *MaxStack) Push(i int)
	func (maxStack *MaxStack) Max() int

Implement a stack that can return the max of element it contains.

For example if we push {2,4,5} to the stack, max should return 5.
*/
var _ = DescribeTable("MaxStack",
	func(push []int, pop int, max int) {
		stack := new(MaxStack)
		for _, p := range push {
			stack.Push(p)
		}
		for range pop {
			_, err := stack.Pop()
			if err != nil {
				Expect(err).ToNot(HaveOccurred())
			}
		}
		got := stack.Max()
		if got != max {
			Expect(got).To(Equal(max))
		}
	},
	Entry("MaxStack #1", []int{}, 0, -1),
	Entry("MaxStack #2", []int{1, 2, 3, 4, 5}, 2, 3),
	Entry("MaxStack #3", []int{1, 2, 3, 4, 5}, 0, 5),
	Entry("MaxStack #4", []int{5, 4, 3, 2, 1}, 2, 5),
	Entry("MaxStack #5", []int{1, 5, 3, 4}, 1, 5),
)

var _ = Describe("EmptyStackError", func() {
	It("works correctly", func() {
		stack := new(MaxStack)
		if _, err := stack.Pop(); !errors.Is(err, ErrEmptyStack) {
			Expect(err).To(Equal(ErrEmptyStack))
		}
	})
})

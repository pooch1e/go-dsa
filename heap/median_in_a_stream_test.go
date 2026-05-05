package heap

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"fmt"
)

/*
TestMedianInAStream tests solution(s) with the following signatures and problem description:

	func newMedianKeeper() medianKeeper
	func (m *medianKeeper) addNumber(num int)

Given a stream of integers, return the median of the set after each input.

For example given {1,2,3,4}, return {1, 1.5, 2, 2.5} after each input respectively.
*/
var _ = DescribeTable("MedianInAStream",
	func(numbers []int, median float64) {
		heap := newMedianKeeper()
		for _, number := range numbers {
			heap.addNumber(number)
		}
		got := heap.median()
		if got != median {
			Expect(got).To(Equal(median))
		}
	},
	Entry("MedianInAStream #1", []int{1}, 1),
	Entry("MedianInAStream #2", []int{2, 3, 4}, 3),
	Entry("MedianInAStream #3", []int{1, 2, 3, 4, 5}, 3),
	Entry("MedianInAStream #4", []int{1, 2, 3, 4, 5, 6}, 3.5),
	Entry("MedianInAStream #5", []int{6, 5, 4, 3, 2, 1}, 3.5),
	Entry("MedianInAStream #6", []int{1, 4, 5, -2, 4, 6}, 4),
)

var _ = Describe("Heap", func() {
	It("works correctly", func() {
		minHeap := new(minHeap)
		minHeap.Push(1)
		got := minHeap.Pop().(int)
		if got != 1 {
			Fail(fmt.Sprintf("expected %d got %d", got, 1))
		}
	})
})

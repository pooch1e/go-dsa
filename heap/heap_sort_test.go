package heap

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"math"
	"slices"
)

/*
TestSort tests solution(s) with the following signature and problem description:

	func HeapSort(list []int) []int

Given a slice of integers, return a sorted slice using Heap Sort.

For example given {3,1,2}, return {1,2,3}.

Heap sort works by inserting all items to a heap and then popping them one by. As heap returns the minimum
element with each pop the outcome is a sorted list of items..
*/
var _ = DescribeTable("Sort",
	func(list []int, sorted []int) {
		if got := Sort(list); !slices.Equal(got, sorted) {
			Expect(got).To(Equal(sorted))
		}
	},
	Entry("Sort #1", []int{}, []int{}),
	Entry("Sort #2", []int{1, 2}, []int{1, 2}),
	Entry("Sort #3", []int{2, 1}, []int{1, 2}),
	Entry("Sort #4", []int{1, 2, 3}, []int{1, 2, 3}),
	Entry("Sort #5", []int{3, 2, 1}, []int{1, 2, 3}),
	Entry("Sort #6", []int{1, 3, 2}, []int{1, 2, 3}),
	Entry("Sort #7", []int{-1, 3, 2, 0, 4}, []int{-1, 0, 2, 3, 4}),
)

var _ = Describe("MinHeapImplementation", func() {
	It("works correctly", func() {
		h := NewMinHeap()
		Expect(h.Pop()).To(Equal(math.MinInt))
		h.Push(1)
		Expect(h.Pop()).To(Equal(1))
	})
})

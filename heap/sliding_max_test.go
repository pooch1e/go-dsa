package heap

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	containerheap "container/heap"
	"slices"
)

/*
TestMaxSlidingWindow tests solution(s) with the following signature and problem description:

	func MaxSlidingWindow(numbers []int, k int) []int

Given a slice of integers and a positive integer k, return the maximum of each sub-slice of the
slice when a window of length k is moved from left to right and we can only see k elements that
are in the window.

For example given {1,4,5,-2,4,6} and k=3, we will have the following windows:

	{1,4,5} -> 5
	{4,5,-2} -> 5
	{5,-2,4} -> 5
	{-2,4,6} -> 6

So return {5,5,5,6}.
*/
var _ = DescribeTable("MaxSlidingWindow",
	func(numbers []int, k int, maxSliding []int) {
		if got := MaxSlidingWindow(numbers, k); !slices.Equal(got, maxSliding) {
			Expect(got).To(Equal(maxSliding))
		}
	},
	Entry("MaxSlidingWindow #1", []int{}, 2, []int{}),
	Entry("MaxSlidingWindow #2", []int{1, 4, 5, -2, 4, 6}, 1, []int{1, 4, 5, -2, 4, 6}),
	Entry("MaxSlidingWindow #3", []int{1, 4, 5, -2, 4, 6}, 2, []int{4, 5, 5, 4, 6}),
	Entry("MaxSlidingWindow #4", []int{1, 4, 5, -2, 4, 6}, 3, []int{5, 5, 5, 6}),
	Entry("MaxSlidingWindow #5", []int{1, 4, 5, -2, 4, 6}, 4, []int{5, 5, 6}),
	Entry("MaxSlidingWindow #6", []int{1, 4, 5, -2, 4, 6}, 6, []int{6}),
)

var _ = Describe("MaxSlidingWindowPop", func() {
	It("works correctly", func() {
		pq := make(slidingWindow, 5)
		containerheap.Init(&pq)
		containerheap.Push(&pq, 5)
		if got := containerheap.Pop(&pq).(int); got != 5 {
			Expect(5).To(Equal(got))
		}
	})
})

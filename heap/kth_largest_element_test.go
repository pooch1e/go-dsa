package heap

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"fmt"
)

/*
TestKthLargestElement tests solution(s) with the following signature and problem description:

	func KthLargestElement(elements []int, k int) int

Given a slice of integers and an integer k, return the kth largest element.

For example given {3,5,6,3,1,2,9} and k=3 return 5, because largest integers are 9, 6, 5 and
the 3rd largest is 5.
*/
var _ = DescribeTable("KthLargestElement",
	func(list []int, k int, kthLargest int) {
		got := KthLargestElement(list, k)
		if got != kthLargest {
			Expect(got).To(Equal(kthLargest))
		}
	},
	Entry("KthLargestElement #1", []int{1}, 1, 1),
	Entry("KthLargestElement #2", []int{1, 2, 3, 4, 5, 6}, 3, 4),
	Entry("KthLargestElement #3", []int{6, 1, 2, 3, 4, 5}, 1, 6),
	Entry("KthLargestElement #4", []int{6, 1, 2, 3, 4, 5}, 6, 1),
)

var _ = Describe("MinHeap", func() {
	It("works correctly", func() {
		minHeap := new(minimumHeap)
		minHeap.Push(1)
		got := minHeap.Pop().(int)
		if got != 1 {
			Fail(fmt.Sprintf("expected %d got %d", got, 1))
		}
	})
})

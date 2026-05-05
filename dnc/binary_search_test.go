package dnc

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestBinarySearch tests solution(s) with the following signature and problem description:

	func BinarySearch(list []int, search int) int

Given a sorted slice of integers like and a target int find the position of the target in the
slice.

In Binary Search we start with the middle element of the set and compare it with
the target. If the target is greater than the middle element we search the right
half of the set, otherwise we search the left half of the set. We repeat this
process until we find the target or we exhaust the set.

For example given {1,2,3,4,6} and 4 as a target, return 3 because 4 is at index 3.
*/
var _ = DescribeTable("BinarySearch",
	func(sortedNumbers []int, search int, position int) {
		if got := BinarySearch(sortedNumbers, search); got != position {
			Expect(got).To(Equal(position))
		}
	},
	Entry("BinarySearch #1", []int{}, 3, -1),
	Entry("BinarySearch #2", []int{1}, 1, 0),
	Entry("BinarySearch #3", []int{1}, 2, -1),
	Entry("BinarySearch #4", []int{1, 2, 3, 4, 5}, 1, 0),
	Entry("BinarySearch #5", []int{1, 2, 3, 4, 5}, 3, 2),
	Entry("BinarySearch #6", []int{1, 2, 3, 5, 10}, 4, -1),
	Entry("BinarySearch #7", []int{1, 2, 3, 8, 10}, 8, 3),
	Entry("BinarySearch #8", []int{1, 2, 3, 8, 10, 11}, 8, 3),
	Entry("BinarySearch #9", []int{1, 2, 3, 4, 5, 9}, 4, 3),
)

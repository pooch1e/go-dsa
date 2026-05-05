package hashtable

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestSmallestMissingPositiveInteger tests solution(s) with the following signature and problem description:

	func SmallestMissingPositiveInteger(input []int) int

Given a slice of integers, return the smallest missing positive integer.

For example given {-9, 0, 1, 2, 3, 4, 5, 6} return 7 because -9 is not a positive integer and
all numbers from 0 to 6 are present in the list.
*/
var _ = DescribeTable("SmallestMissingPositiveInteger",
	func(numbers []int, missing int) {
		if got := SmallestMissingPositiveInteger(numbers); got != missing {
			Expect(got).To(Equal(missing))
		}
	},
	Entry("SmallestMissingPositiveInteger #1", []int{}, 1),
	Entry("SmallestMissingPositiveInteger #2", []int{0}, 1),
	Entry("SmallestMissingPositiveInteger #3", []int{-1}, 1),
	Entry("SmallestMissingPositiveInteger #4", []int{2}, 1),
	Entry("SmallestMissingPositiveInteger #5", []int{1, 2, 3}, 4),
	Entry("SmallestMissingPositiveInteger #6", []int{1, 3}, 2),
	Entry("SmallestMissingPositiveInteger #7", []int{-9, 0, 1, 2, 3, 4, 5, 6}, 7),
)

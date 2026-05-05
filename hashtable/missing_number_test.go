package hashtable

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestMissingNumber tests solution(s) with the following signature and problem description:

	func MissingNumber(numbers []int) int

Given an unsorted slice of numbers like return the missing integer.

For example given {7,5,3,4,1,2,0,-1}, return the missing integer like 6.
*/
var _ = DescribeTable("MissingNumber",
	func(numbers []int, missing int) {
		got := MissingNumber(numbers)
		if got != missing {
			Expect(got).To(Equal(missing))
		}
	},
	Entry("MissingNumber #1", []int{}, -1),
	Entry("MissingNumber #2", []int{2}, -1),
	Entry("MissingNumber #3", []int{1}, -1),
	Entry("MissingNumber #4", []int{1, 2}, -1),
	Entry("MissingNumber #5", []int{1, 2, 3, 4, 6}, 5),
	Entry("MissingNumber #6", []int{1, 2, 3, 4, 5, 6}, -1),
	Entry("MissingNumber #7", []int{6, 3, 2, 1, 4}, 5),
	Entry("MissingNumber #8", []int{-3, -1, 0, 1, 2}, -2),
	Entry("MissingNumber #9", []int{-3, -2, -1, 1, 2}, 0),
)

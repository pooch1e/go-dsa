package bit

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestOddlyRepeatedNumber tests solution(s) with the following signature and problem description:

	func OddlyRepeatedNumber(list []int) int

Given a slice of integers that are all repeated an even number of times except one,
find the oddly repeated element.

For example given {1, 2, 2, 3, 3} return 1. Given {1, 2, 1, 2, 3} return 3.
*/
var _ = DescribeTable("OddlyRepeatedNumber",
	func(list []int, oddlyRepeated int) {
		got := OddlyRepeatedNumber(list)
		if got != oddlyRepeated {
			Expect(got).To(Equal(oddlyRepeated))
		}
	},
	Entry("OddlyRepeatedNumber #1", []int{}, -1),
	Entry("OddlyRepeatedNumber #2", []int{1, 2, 2}, 1),
	Entry("OddlyRepeatedNumber #3", []int{1, 2, 1, 2, 3}, 3),
	Entry("OddlyRepeatedNumber #4", []int{1, 2, 2, 3, 3}, 1),
)

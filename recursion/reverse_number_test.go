package recursion

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestReverseDigits tests solution(s) with the following signature and problem description:

	func ReverseDigits(n int) int

Given an integer reverse the order of the digits.

For example given 123 return 321.
*/
var _ = DescribeTable("ReverseDigits",
	func(n int, reversed int) {
		if got := ReverseDigits(n); got != reversed {
			Expect(got).To(Equal(reversed))
		}
	},
	Entry("ReverseDigits #1", 0, 0),
	Entry("ReverseDigits #2", 1, 1),
	Entry("ReverseDigits #3", 12, 21),
	Entry("ReverseDigits #4", 112, 211),
	Entry("ReverseDigits #5", 110, 11),
	Entry("ReverseDigits #6", 123, 321),
)

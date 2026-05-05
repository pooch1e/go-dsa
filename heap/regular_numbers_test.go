package heap

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"slices"
)

/*
TestRegularNumbers tests solution(s) with the following signature and problem description:

	func RegularNumbers(n int) []int

Regular numbers are numbers whose only prime divisors are 2, 3, and 5.
For example 24 is a regular number because it can be factored into 2^3 * 3^1 * 5^0.

Given a positive integer n, return first n regular numbers.

For example given n=4:

	1 = 2^0
	2 = 2^1
	3 = 3^1
	4 = 2^2

So the first 3 regular numbers are {1, 2, 3, 4}.
*/
var _ = DescribeTable("RegularNumbers",
	func(n int, regularNumbers []int) {
		if got := RegularNumbers(n); !slices.Equal(got, regularNumbers) {
			Expect(got).To(Equal(regularNumbers))
		}
	},
	Entry("RegularNumbers #1", 0, []int{}),
	Entry("RegularNumbers #2", 1, []int{1}),
	Entry("RegularNumbers #3", 2, []int{1, 2}),
	Entry("RegularNumbers #4", 3, []int{1, 2, 3}),
	Entry("RegularNumbers #5", 4, []int{1, 2, 3, 4}),
	Entry("RegularNumbers #6", 5, []int{1, 2, 3, 4, 5}),
	Entry("RegularNumbers #7", 10, []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 12}),
	Entry("RegularNumbers #8", 15, []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15, 16, 18, 20, 24}),
)

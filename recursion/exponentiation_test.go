package recursion

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestPowerOf tests solution(s) with the following signature and problem description:

	func PowerOf(x, n int) int

Given x and n, return x raised to the power of n in an efficient manner.

For example given x=2 and n=3 the algorithm should return 8.
*/
var _ = DescribeTable("PowerOf",
	func(x int, n int, result int) {
		if got := PowerOf(x, n); got != result {
			Expect(got).To(Equal(result))
		}
	},
	Entry("PowerOf #1", 1, 0, 1),
	Entry("PowerOf #2", 1, 1, 1),
	Entry("PowerOf #3", 1, 2, 1),
	Entry("PowerOf #4", 2, 1, 2),
	Entry("PowerOf #5", 2, 2, 4),
	Entry("PowerOf #6", 2, 3, 8),
	Entry("PowerOf #7", 3, 3, 27),
	Entry("PowerOf #8", 5, 2, 25),
	Entry("PowerOf #9", 5, 3, 125),
	Entry("PowerOf #10", 5, 4, 625),
)

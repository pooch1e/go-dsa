package bit

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestMiddleWithoutDivision tests solution(s) with the following signature and problem description:

	func MiddleWithoutDivision(min, max int)

Given two integers return the integer that is in the middle of the two integers without
using any arithmetic operators such as {+,-,/,*,++,--,+=,…}.

For example given 1 and 5, return 3. This is because 3 is in the middle of integers 1 to 5.
*/
var _ = DescribeTable("MiddleWithoutDivision",
	func(a int, b int, mid int) {
		if got := MiddleWithoutDivision(a, b); got != mid {
			Expect(got).To(Equal(mid))
		}
	},
	Entry("MiddleWithoutDivision #1", 0, 0, 0),
	Entry("MiddleWithoutDivision #2", 0, 1, 0),
	Entry("MiddleWithoutDivision #3", 0, 5, 2),
	Entry("MiddleWithoutDivision #4", 1, 5, 3),
	Entry("MiddleWithoutDivision #5", 0, 6, 3),
	Entry("MiddleWithoutDivision #6", 1, 5, 3),
	Entry("MiddleWithoutDivision #7", 2, 20, 11),
	Entry("MiddleWithoutDivision #8", 2, -2, 0),
	Entry("MiddleWithoutDivision #9", -3, -2, -3),
)

package bit

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestMax tests solution(s) with the following signature and problem description:

	func Max(x, y int) int

Given two integers, return the larger of the two without using any comparison
operations like {if, switch,…}.

For example given 20 and 2 return 20.
*/
var _ = DescribeTable("Max",
	func(a int, b int, max int) {
		got := Max(a, b)
		if got != max {
			Expect(got).To(Equal(max))
		}
	},
	Entry("Max #1", 0, 1, 1),
	Entry("Max #2", 20, 2, 20),
	Entry("Max #3", 2, 20, 20),
	Entry("Max #4", 2, -2, 2),
	Entry("Max #5", -3, -2, -2),
)

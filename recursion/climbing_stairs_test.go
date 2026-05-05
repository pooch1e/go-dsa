package recursion

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestClimbingStairs tests solution(s) with the following signature and problem description:

	func ClimbingStairs(n int) int

Given n the number of steps, return in how many ways you can climb these stairs if you are
only able to climb 1 or 2 steps at a time.

For example given 5 we can climb the stairs in the following ways:

	1, 1, 1, 1, 1
	1, 1, 1, 2
	1, 1, 2, 1
	1, 2, 1, 1
	2, 1, 1, 1
	2, 2, 1,
	1, 2, 2,
	2, 1, 2,

So the algorithm should return 8.
*/
var _ = DescribeTable("ClimbingStairs",
	func(n int, ways int) {
		if got := ClimbingStairs(n); got != ways {
			Expect(got).To(Equal(ways))
		}
	},
	Entry("ClimbingStairs #1", 0, 1),
	Entry("ClimbingStairs #2", 2, 2),
	Entry("ClimbingStairs #3", 3, 3),
	Entry("ClimbingStairs #4", 5, 8),
	Entry("ClimbingStairs #5", 7, 21),
	Entry("ClimbingStairs #6", 10, 89),
)

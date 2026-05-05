package greedy

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestJumpGame tests solution(s) with the following signature and problem description:

	func JumpGame(jumps []int) bool

Given a list of integers that representing the maximum posit ions one can jump at any
given position like {1,2,4,2,1} (at position 1 we can jump up to 1, at position 2 we can
jump up to 2), return true if one can reach the last position of the list and false
otherwise. The response is true for this example.
*/
var _ = DescribeTable("JumpGame",
	func(jumps []int, canReachTheEnd bool) {
		if got := JumpGame(jumps); got != canReachTheEnd {
			Expect(got).To(Equal(canReachTheEnd))
		}
	},
	Entry("JumpGame #1", []int{}, true),
	Entry("JumpGame #2", []int{1}, true),
	Entry("JumpGame #3", []int{1, 0, 1}, false),
	Entry("JumpGame #4", []int{3, 1, 0, 0, 1}, false),
	Entry("JumpGame #5", []int{2, 3, 1, 1, 4}, true),
	Entry("JumpGame #6", []int{1, 2, 4, 2, 1}, true),
	Entry("JumpGame #7", []int{2, 3, 1, 1, 2, 0, 1}, true),
)

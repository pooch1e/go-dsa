package dp

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestMaxHouseRobber tests solution(s) with the following signature and problem description:

	func MaxHouseRobber(wealth []int) int

Given an array representing the amount of wealth inside houses like {1,2,3,4} return the
maximum wealth the robber can steal like `6` with the condition that he is not allowed to
rob two consecutive houses.
*/
var _ = DescribeTable("MaxHouseRobber",
	func(houses []int, max int) {
		if got := MaxHouseRobber(houses); got != max {
			Expect(got).To(Equal(max))
		}
	},
	Entry("MaxHouseRobber #1", []int{}, 0),
	Entry("MaxHouseRobber #2", []int{1}, 1),
	Entry("MaxHouseRobber #3", []int{1, 2}, 2),
	Entry("MaxHouseRobber #4", []int{1, 2, 3, 4}, 6),
	Entry("MaxHouseRobber #5", []int{1, 3, 5, 10, 15}, 21),
	Entry("MaxHouseRobber #6", []int{1, 5, 8, 9, 10}, 19),
	Entry("MaxHouseRobber #7", []int{1, 5, 8, 9, 10, 20, 30}, 49),
)

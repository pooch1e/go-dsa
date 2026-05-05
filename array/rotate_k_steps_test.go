package array

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"slices"
)

/*
TestRotateKSteps tests solution(s) with the following signature and problem description:

	RotateKSteps(list []int, k int)

Given an list of integers and a number k, rotate the array k times.

For example given {1,2,3} and 3, return {1,2,3} because:

After 1 rotation: {3,1,2}.
After 2 rotations: {2,3,1}.
After 3 rotations: {1,2,3}.
*/
var _ = DescribeTable("RotateKSteps",
	func(list []int, steps int, rotatedList []int) {
		RotateKSteps(list, steps)
		if !slices.Equal(list, rotatedList) {
			Expect(list).To(Equal(rotatedList))
		}
	},
	Entry("RotateKSteps #1", []int{}, 0, []int{}),
	Entry("RotateKSteps #2", []int{1, 2, 3}, 1, []int{3, 1, 2}),
	Entry("RotateKSteps #3", []int{1, 2, 3}, 2, []int{2, 3, 1}),
	Entry("RotateKSteps #4", []int{1, 2, 3}, 3, []int{1, 2, 3}),
	Entry("RotateKSteps #5", []int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}),
)

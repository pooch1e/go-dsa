package hashtable

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"slices"
)

/*
TestSumUpToK tests solution(s) with the following signature and problem description:

	func SumUpToK(numbers []int, k int) []int

Given a slice of integers, output the indices of the first two elements that sum up to K.

For example given {1, 2, 3, 4} and K = 5, return {1, 2} because 2 + 3 = 5.
*/
var _ = DescribeTable("SumUpToK",
	func(k int, numbers []int, indices []int) {
		got := SumUpToK(numbers, k)
		if !slices.Equal(got, indices) {
			Expect(got).To(Equal(indices))
		}
	},
	Entry("SumUpToK #1", 6, []int{3, 3}, []int{0, 1}),
	Entry("SumUpToK #2", 3, []int{1, 2, 3, 4}, []int{0, 1}),
	Entry("SumUpToK #3", 5, []int{1, 2, 3, 6}, []int{1, 2}),
	Entry("SumUpToK #4", 7, []int{1, 2, 3, 4}, []int{2, 3}),
)

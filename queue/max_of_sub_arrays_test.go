package queue

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"slices"
)

/*
TestMaxOfKLengthSubArrays tests solution(s) with the following signature and problem description:

	func MaxOfKLengthSubArrays(numbers []int, k int) ([]int, error)

Given a slice of integers and an integer k, return a slice containing the maximum in each k-sized
sub-array (sub-slice) of the input.

For example given {1,2,3,4,5} and k=2 return {2,3,4,5} because:

* Sub-arrays of the input with length 2 are {{1,2},{2,3},{3,4},{4,5}}
* The maximum in each of the sub-arrays is {2,3,4,5}.
*/
var _ = DescribeTable("MaxOfKLengthSubArrays",
	func(k int, list []int, maximums []int) {
		got, err := MaxOfKLengthSubArrays(list, k)
		if err != nil {
			Expect(err).ToNot(HaveOccurred())
		}
		if !slices.Equal(got, maximums) {
			Expect(got).To(Equal(maximums))
		}
	},
	Entry("MaxOfKLengthSubArrays #1", 3, []int{5, 3, 5, 6, 7, 8}, []int{5, 6, 7, 8}),
	Entry("MaxOfKLengthSubArrays #2", 2, []int{5, 3, 5, 6, 7, 8}, []int{5, 5, 6, 7, 8}),
	Entry("MaxOfKLengthSubArrays #3", 3, []int{1, 2, 3, 4, 5}, []int{3, 4, 5}),
)

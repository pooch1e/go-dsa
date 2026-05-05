package array

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"reflect"
)

/*
TestEqualSumSubArrays tests solution(s) with the following signature and problem description:

	func EqualSubArrays(list []int) [][]int

Given an list of integers A, return two sub-arrays with equal sums that are a partition of the
original list, without changing the order of the elements in the list.

For example given {1,7,3,5}, return {1,7} and {3,5} because 1+7 = 3+5 = 8.
*/
var _ = DescribeTable("EqualSumSubArrays",
	func(list []int, subArrays [][]int) {
		if got := EqualSubArrays(list); !reflect.DeepEqual(got, subArrays) {
			Expect(got).To(Equal(subArrays))
		}
	},
	Entry("EqualSumSubArrays #1", []int{}, [][]int{}),
	Entry("EqualSumSubArrays #2", []int{1}, [][]int{}),
	Entry("EqualSumSubArrays #3", []int{1, 2, 2, 3}, [][]int{}),
	Entry("EqualSumSubArrays #4", []int{1, 2, 3}, [][]int{{1, 2}, {3}}),
	Entry("EqualSumSubArrays #5", []int{2, 3, 5}, [][]int{{2, 3}, {5}}),
	Entry("EqualSumSubArrays #6", []int{1, 7, 3, 5}, [][]int{{1, 7}, {3, 5}}),
	Entry("EqualSumSubArrays #7", []int{-4, 1, 1, 1, 1}, [][]int{}),
	Entry("EqualSumSubArrays #8", []int{4, 1, 1, 1, 1}, [][]int{{4}, {1, 1, 1, 1}}),
	Entry("EqualSumSubArrays #9", []int{1, 1, 1, 1, 4}, [][]int{{1, 1, 1, 1}, {4}}),
)

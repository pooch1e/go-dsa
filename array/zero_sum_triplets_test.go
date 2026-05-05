package array

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"reflect"
)

/*
TestZeroSumTriplets tests solution(s) with the following signature and problem description:

	ZeroSumTriplets(list []int) [][]int

Given an array of numbers like {1, 2, -4, 6, 3} returns unique triplets from the numbers
with sum that equals zero like {-4, 1, 3} because -4+1+3=0.
*/
var _ = DescribeTable("ZeroSumTriplets",
	func(list []int, triplets [][]int) {
		if got := ZeroSumTriplets(list); !reflect.DeepEqual(got, triplets) {
			Expect(got).To(Equal(triplets))
		}
	},
	Entry("ZeroSumTriplets #1", []int{}, [][]int{}),
	Entry("ZeroSumTriplets #2", []int{1}, [][]int{}),
	Entry("ZeroSumTriplets #3", []int{1, 2}, [][]int{}),
	Entry("ZeroSumTriplets #4", []int{1, 2, 3}, [][]int{}),
	Entry("ZeroSumTriplets #5", []int{1, -4, 3}, [][]int{{-4, 1, 3}}),
	Entry("ZeroSumTriplets #6", []int{0, 0, 0, 0}, [][]int{{0, 0, 0}}),
	Entry("ZeroSumTriplets #7", []int{1, 2, -4, 6, 3}, [][]int{{-4, 1, 3}}),
	Entry("ZeroSumTriplets #8", []int{-1, -2, -8, 6, 2, 3}, [][]int{{-8, 2, 6}, {-2, -1, 3}}),
	Entry("ZeroSumTriplets #9", []int{1, -2, -4, 5, -2, 4, 1, 3}, [][]int{{-4, 1, 3}, {-2, -2, 4}, {-2, 1, 1}}),
)

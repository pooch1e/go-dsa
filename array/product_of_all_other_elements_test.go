package array

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"slices"
)

/*
TestProductOfAllOtherElements tests solution(s) with the following signature and problem description:

	ProductOfAllOtherElements(list []int) []int

Given an array of integers A, construct a new array B such that B[i] = product of all items
in A except A[i] without using division in O(n) time.

For example given {1,2,3,4}, return {24,12,8,6} because:
* 24=2*3*4.
* 12=1*3*4.
* 8=1*2*4.
* 6=1*2*3.
*/
var _ = DescribeTable("ProductOfAllOtherElements",
	func(list []int, products []int) {
		if got := ProductOfAllOtherElements(list); !slices.Equal(got, products) {
			Expect(got).To(Equal(products))
		}
	},
	Entry("ProductOfAllOtherElements #1", []int{}, []int{}),
	Entry("ProductOfAllOtherElements #2", []int{2, 3}, []int{3, 2}),
	Entry("ProductOfAllOtherElements #3", []int{1, 2, 3}, []int{6, 3, 2}),
	Entry("ProductOfAllOtherElements #4", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}),
	Entry("ProductOfAllOtherElements #5", []int{1, 2, 3, 4, 5}, []int{120, 60, 40, 30, 24}),
)

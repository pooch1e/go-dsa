package array

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"slices"
)

/*
TestAddSliceOfTwoNumbers tests solution(s) with the following signature and problem description:

	AddTwoNumbers(num1, num2 []int) []int

A slice representation of a positive integer like 283 looks like {2,8,3}. Given two positive
integers represented in this format return their sum in the same format.

For example given {2,9} and {9,9,9}, return {1,0,2,8}.
Because 29+999=1028.
*/
var _ = DescribeTable("AddSliceOfTwoNumbers",
	func(num1 []int, num2 []int, sum []int) {
		if got := AddSliceOfTwoNumbers(num1, num2); !slices.Equal(got, sum) {
			Expect(got).To(Equal(sum))
		}
	},
	Entry("AddSliceOfTwoNumbers #1", []int{1}, []int{}, []int{1}),
	Entry("AddSliceOfTwoNumbers #2", []int{1}, []int{0}, []int{1}),
	Entry("AddSliceOfTwoNumbers #3", []int{1}, []int{1}, []int{2}),
	Entry("AddSliceOfTwoNumbers #4", []int{1}, []int{9}, []int{1, 0}),
	Entry("AddSliceOfTwoNumbers #5", []int{2, 5}, []int{3, 5}, []int{6, 0}),
	Entry("AddSliceOfTwoNumbers #6", []int{2, 9}, []int{9, 9, 9}, []int{1, 0, 2, 8}),
	Entry("AddSliceOfTwoNumbers #7", []int{9, 9, 9}, []int{9, 9, 9}, []int{1, 9, 9, 8}),
)

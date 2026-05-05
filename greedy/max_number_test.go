package greedy

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"slices"
)

/*
TestMaxNumber tests solution(s) with the following signature and problem description:

	func MaxNumber(number1, number2 []int, n int) []int

Given two numbers represented as list of positive integers, and an integer n, return the
largest possible integer with n digits that can be constructed by merging digits from two
numbers while respecting the order of digits in each number.

For example given {5, 4, 3, 2, 1}, {9, 8, 7, 6} and 9, return
{9, 8, 7, 6, 5, 4, 3, 2, 1} because it is the largest 9 digit combination of all elements
of the two given slices.
*/
var _ = DescribeTable("MaxNumber",
	func(number1 []int, number2 []int, k int, largestNumber []int) {
		if got := MaxNumber(number1, number2, k); !slices.Equal(got, largestNumber) {
			Expect(got).To(Equal(largestNumber))
		}
	},
	Entry("MaxNumber #1", []int{}, []int{}, 1, []int{}),
	Entry("MaxNumber #2", []int{}, []int{1}, 1, []int{1}),
	Entry("MaxNumber #3", []int{0}, []int{1}, 1, []int{1}),
	Entry("MaxNumber #4", []int{0}, []int{1}, 2, []int{1, 0}),
	Entry("MaxNumber #5", []int{1}, []int{0}, 2, []int{1, 0}),
	Entry("MaxNumber #6", []int{1}, []int{2}, 4, []int{}),
	Entry("MaxNumber #7", []int{2}, []int{1}, 4, []int{}),
	Entry("MaxNumber #8", []int{3, 2, 1}, []int{1}, 3, []int{3, 2, 1}),
	Entry("MaxNumber #9", []int{6, 9}, []int{6, 0, 4}, 5, []int{6, 9, 6, 0, 4}),
	Entry("MaxNumber #10", []int{1, 2, 3}, []int{1, 2, 3}, 4, []int{3, 1, 2, 3}),
	Entry("MaxNumber #11", []int{5, 4, 3, 2, 1}, []int{9, 8, 7, 6}, 9, []int{9, 8, 7, 6, 5, 4, 3, 2, 1}),
	Entry("MaxNumber #12", []int{5, 4, 3, 2, 1}, []int{9, 8, 7, 6}, 8, []int{9, 8, 7, 6, 5, 4, 3, 2}),
)

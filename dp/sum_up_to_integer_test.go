package dp

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestSumUpToInteger tests solution(s) with the following signature and problem description:

	func SumUpToInteger(numbers []int, sum int) bool

Given a set of positive integers like {1,2,3,4,5} and an integer like 7 write a
function that returns true if there are two numbers in the list that sum up to the given
integer and false otherwise.
*/
var _ = DescribeTable("SumUpToInteger",
	func(numbers []int, sum int, sumsUp bool) {
		if got := SumUpToInteger(numbers, sum); got != sumsUp {
			Expect(got).To(Equal(sumsUp))
		}
	},
	Entry("SumUpToInteger #1", []int{1}, 1, true),
	Entry("SumUpToInteger #2", []int{1, 2}, 2, true),
	Entry("SumUpToInteger #3", []int{1, 2}, 3, true),
	Entry("SumUpToInteger #4", []int{1, 2}, 4, false),
	Entry("SumUpToInteger #5", []int{1, 2, 3, 4, 5}, 7, true),
	Entry("SumUpToInteger #6", []int{1, 2, 3, 4, 5}, 8, true),
	Entry("SumUpToInteger #7", []int{1, 2, 3, 4, 5}, 15, true),
	Entry("SumUpToInteger #8", []int{1, 2, 3, 4, 5}, 16, false),
	Entry("SumUpToInteger #9", []int{1, 5, 8, 9, 10, 20, 30}, 25, true),
)

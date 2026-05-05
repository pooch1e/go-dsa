package dp

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestRodCutting tests solution(s) with the following signature and problem description:

	CutRod(prices []int, n int) int

Given a list containing price a table such as {1,5,8,9,10} indicating the price of
a rod of a given length (1 inch rod is $1, 2 inch rod is $5, 5 inch rod is $10) and
number n like 3, indicating the length of a given rod, calculate maximum revenue that
can be earned by cutting the rod and selling the pieces when cutting is free.
*/
var _ = DescribeTable("RodCutting",
	func(snacks []int, n int, solution int) {
		if got := CutRod(snacks, n); got != solution {
			Expect(got).To(Equal(solution))
		}
	},
	Entry("RodCutting #1", []int{1}, 1, 1),
	Entry("RodCutting #2", []int{1, 2}, 2, 2),
	Entry("RodCutting #3", []int{1, 3, 5, 10, 15}, 2, 3),
	Entry("RodCutting #4", []int{1, 3, 5, 10, 15}, 3, 5),
	Entry("RodCutting #5", []int{1, 5, 8, 9, 10}, 4, 10),
	Entry("RodCutting #6", []int{1, 5, 8, 9, 10, 20, 30}, 4, 10),
)

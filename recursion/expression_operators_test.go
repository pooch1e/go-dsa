package recursion

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestExpressionOperators tests solution(s) with the following signature and problem description:

	func ExpressionOperators(list []int, target int) string

Given a slice of numbers representing operands in an equation, and a target integer representing
the result of the equation, return a string representing operators that can be inserted between
the operands to form the equation and yield the target result.
Only + and - operators are allowed and the are assumed to have the same priority

For example given {1,5,3} and 3 return {+,-} because 1+5-3 = 3.
*/
var _ = DescribeTable("ExpressionOperators",
	func(operands []int, target int, operators string) {
		if got := ExpressionOperators(operands, target); got != operators {
			Expect(got).To(Equal(operators))
		}
	},
	Entry("ExpressionOperators #1", []int{}, 1, ""),
	Entry("ExpressionOperators #2", []int{1}, 1, ""),
	Entry("ExpressionOperators #3", []int{1, 2}, 3, "+"),
	Entry("ExpressionOperators #4", []int{1, 2, 3}, 0, "+-"),
	Entry("ExpressionOperators #5", []int{1, 5, 3}, 3, "+-"),
	Entry("ExpressionOperators #6", []int{1, 2, 3, 4, 5, 6}, 5, "+-+-+"),
)

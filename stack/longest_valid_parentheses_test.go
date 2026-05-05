package stack

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestLongestValidParentheses tests solution(s) with the following signature and problem description:

	func LongestValidParentheses(input string) int

Given a string containing parentheses, find the length of the longest valid (well-formed)
parentheses substring.

For example given "(()", return 2 because the longest valid parentheses substring is "()".
*/
var _ = DescribeTable("LongestValidParentheses",
	func(input string, longestValidParentheses int) {
		if got := LongestValidParentheses(input); got != longestValidParentheses {
			Expect(got).To(Equal(longestValidParentheses))
		}
	},
	Entry("LongestValidParentheses #1", "", 0),
	Entry("LongestValidParentheses #2", "(", 0),
	Entry("LongestValidParentheses #3", "()", 2),
	Entry("LongestValidParentheses #4", "())", 2),
	Entry("LongestValidParentheses #5", "(()", 2),
	Entry("LongestValidParentheses #6", "((()", 2),
	Entry("LongestValidParentheses #7", "((())", 4),
	Entry("LongestValidParentheses #8", "()(()", 2),
	Entry("LongestValidParentheses #9", "((())()", 6),
	Entry("LongestValidParentheses #10", "((())()(((", 6),
	Entry("LongestValidParentheses #11", "))((())()(((", 6),
)

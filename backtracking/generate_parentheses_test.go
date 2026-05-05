package backtracking

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"slices"
	"sort"
)

/*
TestGenerateParentheses tests solution(s) with the following signature and problem description:

	GenerateParentheses(n int) []string

Given an integer n produce all valid variations of arranging
n pair of parentheses. e.g. for `2` it should return `()()` and `(())`.
*/
var _ = DescribeTable("GenerateParentheses",
	func(n int, validParentheses []string) {
		got := GenerateParentheses(n)
		if len(got) > 0 {
			sort.Strings(got)
		}
		if !slices.Equal(validParentheses, got) {
			Expect(got).To(Equal(validParentheses))
		}
	},
	Entry("GenerateParentheses #1", 0, []string{""}),
	Entry("GenerateParentheses #2", 1, []string{"()"}),
	Entry("GenerateParentheses #3", 2, []string{"(())", "()()"}),
	Entry("GenerateParentheses #4", 3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}),
	Entry("GenerateParentheses #5", 4, []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"}),
	Entry("GenerateParentheses #6", 5, []string{"((((()))))", "(((()())))", "(((())()))", "(((()))())", "(((())))()", "((()(())))", "((()()()))", "((()())())", "((()()))()", "((())(()))", "((())()())", "((())())()", "((()))(())", "((()))()()", "(()((())))", "(()(()()))", "(()(())())", "(()(()))()", "(()()(()))", "(()()()())", "(()()())()", "(()())(())", "(()())()()", "(())((()))", "(())(()())", "(())(())()", "(())()(())", "(())()()()", "()(((())))", "()((()()))", "()((())())", "()((()))()", "()(()(()))", "()(()()())", "()(()())()", "()(())(())", "()(())()()", "()()((()))", "()()(()())", "()()(())()", "()()()(())", "()()()()()"}),
)

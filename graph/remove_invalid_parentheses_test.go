package graph

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"slices"
)

/*
TestRemoveInvalidParentheses tests solution(s) with the following signature and problem description:

	func RemoveInvalidParentheses(input string) []string

Given a string containing parentheses and other alphabet letters like `(z)())()`, remove the minimum
amount of parentheses to make the string valid like `(z())()` and `(z)()()`.
*/
var _ = DescribeTable("RemoveInvalidParentheses",
	func(input string, outputs []string) {
		if got := RemoveInvalidParentheses(input); !slices.Equal(got, outputs) {
			Expect(got).To(Equal(outputs))
		}
	},
	Entry("RemoveInvalidParentheses #1", "", []string{}),
	Entry("RemoveInvalidParentheses #2", ")(", []string{""}),
	Entry("RemoveInvalidParentheses #3", ")v(", []string{"v"}),
	Entry("RemoveInvalidParentheses #4", "(v)", []string{"(v)"}),
	Entry("RemoveInvalidParentheses #5", "(z)())()", []string{"(z())()", "(z)()()"}),
	Entry("RemoveInvalidParentheses #6", "()()z)()", []string{"(()z)()", "()(z)()", "()()z()"}),
	Entry("RemoveInvalidParentheses #7", "()())()", []string{"(())()", "()()()"}),
)

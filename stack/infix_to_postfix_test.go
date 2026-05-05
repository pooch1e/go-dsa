package stack

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"slices"
)

/*
TestInfixToPostfix tests solution(s) with the following signature and problem description:

	func InfixToPostfix(infix []string) []string

Given an infix expression convert it to a postfix expression supporting the four basic arithmetic
operations and parentheses.

Infix expression is how humans typically write arithmetic expressions like 1*2+3+4*5 which
is equivalent of (1*2) + 3 + (4*5).

For example given 1*2+3+4*5, return 1 2 * 3 + 4 5 * which both evaluate to 25.
*/
var _ = DescribeTable("InfixToPostfix",
	func(infix []string, postfix []string) {
		if got := InfixToPostfix(infix); !slices.Equal(got, postfix) {
			Expect(got).To(Equal(postfix))
		}
	},
	Entry("InfixToPostfix #1", []string{""}, []string{""}),
	Entry("InfixToPostfix #2", []string{"a", "+", "b"}, []string{"a", "b", "+"}),
	Entry("InfixToPostfix #3", []string{"a", "-", "b", "+", "c"}, []string{"a", "b", "c", "+", "-"}),
	Entry("InfixToPostfix #4", []string{"a", "-", "(", "b", "+", "c", ")"}, []string{"a", "b", "c", "+", "-"}),
	Entry("InfixToPostfix #5", []string{"a", "+", "b", "-", "c"}, []string{"a", "b", "c", "-", "+"}),
	Entry("InfixToPostfix #6", []string{"a", "/", "b"}, []string{"a", "b", "/"}),
	Entry("InfixToPostfix #7", []string{"1", "*", "2", "+", "3", "+", "4", "*", "5"}, []string{"1", "2", "3", "4", "5", "*", "+", "+", "*"}),
	Entry("InfixToPostfix #8", []string{"1", "*", "(", "2", "+", "3", ")", "+", "4", "*", "5"}, []string{"1", "2", "3", "+", "4", "5", "*", "+", "*"}),
)

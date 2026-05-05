package stack

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestEvaluatePostfix tests solution(s) with the following signature and problem description:

	func EvaluatePostfixExpression(expression []string) (float64, error)

Given a postfix expression calculate its the value.

The postfix expression is a list of strings where each string is either an operator like
arithmetic symbols or an operand that are numbers.

To evaluate the postfix expression, we start scanning the expression from left to right.
If the current element is an operator we apply the operand to the last two operands we read
we then remove the operator and replace the two operands with the results of the operation.

For example given 1 2 3 + *, return 5 because 1 * (2 + 3) = 5.
*/
var _ = DescribeTable("EvaluatePostfix",
	func(postfix []string, expectErr bool, outcome float64) {
		got, err := EvaluatePostfixExpression(postfix)
		if err != nil && !expectErr {
			Expect(err).ToNot(HaveOccurred())
		}
		if got != outcome {
			Expect(got).To(Equal(outcome))
		}
	},
	Entry("EvaluatePostfix #1", []string{""}, true, -1),
	Entry("EvaluatePostfix #2", []string{"+"}, true, -1),
	Entry("EvaluatePostfix #3", []string{"A", "B", "+"}, true, -1),
	Entry("EvaluatePostfix #4", []string{"1", "B", "+"}, true, -1),
	Entry("EvaluatePostfix #5", []string{"1", "2", "+"}, false, 3),
	Entry("EvaluatePostfix #6", []string{"1", "2", "3", "+", "*"}, false, 5),
	Entry("EvaluatePostfix #7", []string{"1", "2", "3", "+", "+"}, false, 6),
	Entry("EvaluatePostfix #8", []string{"1", "3", "2", "-", "+"}, false, 2),
	Entry("EvaluatePostfix #9", []string{"9", "3", "/"}, false, 3),
	Entry("EvaluatePostfix #10", []string{"3", "9", "3", "/", "-"}, false, 0),
	Entry("EvaluatePostfix #11", []string{"1", "2", "3", "4", "5", "*", "+", "+", "*"}, false, 25),
	Entry("EvaluatePostfix #12", []string{"1", "2", "3", "+", "4", "5", "*", "+", "*"}, false, 25),
)

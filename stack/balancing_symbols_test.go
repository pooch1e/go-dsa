package stack

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestIsExpressionBalanced tests solution(s) with the following signature and problem description:

	func IsExpressionBalanced(s string) bool

Given a set of symbols including []{}(), determine if the input is is balanced or not.

For example {{(}} is not balanced because there is no closing `)` while {{()}} is balanced.
*/
var _ = Describe("IsExpressionBalanced", func() {
	It("works correctly", func() {
		tests := []struct {
			expression string
			isValid    bool
		}{
			{"", true},
			{"()", true},
			{"(){", false},
			{"(){}", true},
			{"(){}]", false},
			{"(){}][", false},
			{"(){}[]", true},
			{"({}[])", true},
			{"({[]})", true},
			{"({[])", false},
			{"(({[]))", false},
			{")({[])", false},
		}

		for i, test := range tests {
			if got := IsExpressionBalanced(test.expression); got != test.isValid {
				Expect(test.isValid, got).To(Equal(i))
			}
		}
	})
})

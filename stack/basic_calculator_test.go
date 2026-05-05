package stack

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestBasicCalculator tests solution(s) with the following signature and problem description:

func BasicCalculator(input string) (float64, error)

Given an expression containing integers, parentheses and the four basic arithmetic operations
{*,/,+,-} starting from the highest priority to lowest.

For example given 1*2+3+4*5 return 25 because (4*5) + (1*2) + 3 = 25.
*/
var _ = DescribeTable("BasicCalculator",
	func(expression string, expectErr bool, outcome float64) {
		got, err := BasicCalculator(expression)
		if err != nil && !expectErr {
			Expect(err).ToNot(HaveOccurred())
		}
		if got != outcome {
			Expect(got).To(Equal(outcome))
		}
	},
	Entry("BasicCalculator #1", "", true, -1),
	Entry("BasicCalculator #2", "1++", true, -1),
	Entry("BasicCalculator #3", "1+2", false, 3),
	Entry("BasicCalculator #4", "1*(2+3)", false, 5),
	Entry("BasicCalculator #5", "1+2+3", false, 6),
	Entry("BasicCalculator #6", "1+(3-2)", false, 2),
	Entry("BasicCalculator #7", "9/3", false, 3),
	Entry("BasicCalculator #8", "3-9/3", false, 0),
	Entry("BasicCalculator #9", "1*(2+3+(4*5))", false, 25),
	Entry("BasicCalculator #10", "1*(2+3)+(4*5)", false, 25),
	Entry("BasicCalculator #11", "5.10/2", false, 2.55),
)

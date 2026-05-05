package recursion

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestRegularExpressions tests solution(s) with the following signature and problem description:

	func IsRegularExpressionMatch(input, pattern string) bool

Given an input and a regular expression pattern where

	`.` denotes to any character
	`*` denotes to zero or more of the proceeding characters

Write a recursive function to return true if the input matches the pattern and false otherwise.

For example given input `aa` and pattern `a*` the algorithm should return true, but given the same
pattern and "ba" it should return false.
*/
var _ = DescribeTable("RegularExpressions",
	func(input string, pattern string, match bool) {
		if got := IsRegularExpressionMatch(input, pattern); got != match {
			Expect(got).To(Equal(match))
		}
	},
	Entry("RegularExpressions #1", "", "", true),
	Entry("RegularExpressions #2", "", "*", false),
	Entry("RegularExpressions #3", "a", "", false),
	Entry("RegularExpressions #4", "a", ".", true),
	Entry("RegularExpressions #5", "a", "*", false),
	Entry("RegularExpressions #6", "aa", "*", false),
	Entry("RegularExpressions #7", "aa", "*a", false),
	Entry("RegularExpressions #8", "aa", "a*", true),
	Entry("RegularExpressions #9", "ba", "a*", false),
	Entry("RegularExpressions #10", "aa", ".", false),
	Entry("RegularExpressions #11", "ab", ".", false),
	Entry("RegularExpressions #12", "ad", "d", false),
	Entry("RegularExpressions #13", "ad", ".d", true),
	Entry("RegularExpressions #14", "da", "*d", false),
	Entry("RegularExpressions #15", "da", ".*", true),
	Entry("RegularExpressions #16", "ad", "d", false),
	Entry("RegularExpressions #17", "abdef", "*d*", false),
)

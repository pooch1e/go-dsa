package backtracking

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"slices"
	"sort"
)

/*
TestPhoneLetterCombinations tests solution(s) with the following signature and problem description:

	func PhoneLetterCombinations(digits string) []string

Intakes a string of digits from 2 to 9 inclusive and returns all possible combinations of letters
that could be generated from those. For example for input 2 it should return abc.
*/
var _ = DescribeTable("PhoneLetterCombinations",
	func(digits string, combinations []string) {
		got := PhoneLetterCombinations(digits)
		if len(got) > 0 {
			sort.Strings(got)
		}
		if !slices.Equal(combinations, got) {
			Expect(got).To(Equal(combinations))
		}
	},
	Entry("PhoneLetterCombinations #1", "", []string{}),
	Entry("PhoneLetterCombinations #2", "1", []string{}),
	Entry("PhoneLetterCombinations #3", "2", []string{"a", "b", "c"}),
	Entry("PhoneLetterCombinations #4", "9", []string{"w", "x", "y", "z"}),
	Entry("PhoneLetterCombinations #5", "12", []string{}),
	Entry("PhoneLetterCombinations #6", "23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}),
	Entry("PhoneLetterCombinations #7", "232", []string{"ada", "adb", "adc", "aea", "aeb", "aec", "afa", "afb", "afc", "bda", "bdb", "bdc", "bea", "beb", "bec", "bfa", "bfb", "bfc", "cda", "cdb", "cdc", "cea", "ceb", "cec", "cfa", "cfb", "cfc"}),
)

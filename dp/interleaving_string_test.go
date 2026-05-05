package dp

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestIsInterleavingString tests solution(s) with the following signature and problem description:

	IsInterleavingString(a, b, input string) bool

Given three strings a,b, and input, return true if input is made by interleaving a and b such that
it contains all their characters while maintaining their order.

Example:
"a", "bcd", "abcd" would return true.
"a", "bdc", "abd" would return false, because c is missing.
"a", "bdc", "abc" would return false, because d is missing.
*/
var _ = DescribeTable("IsInterleavingString",
	func(a string, b string, input string, isInterleavingString bool) {
		if got := IsInterleavingString(a, b, input); isInterleavingString != got {
			Expect(got).To(Equal(isInterleavingString))
		}
	},
	Entry("IsInterleavingString #1", "a", "b", "ab", true),
	Entry("IsInterleavingString #2", "abc", "def", "abcdef", true),
	Entry("IsInterleavingString #3", "ac", "bdef", "abcdef", true),
	Entry("IsInterleavingString #4", "abe", "cdf", "abcdef", true),
	Entry("IsInterleavingString #5", "abd", "cef", "abcdef", true),
	Entry("IsInterleavingString #6", "a", "bdc", "abdc", true),
	Entry("IsInterleavingString #7", "a", "bdc", "abd", false),
	Entry("IsInterleavingString #8", "a", "bdc", "abc", false),
	Entry("IsInterleavingString #9", "bdc", "a", "abc", false),
	Entry("IsInterleavingString #10", "bcd", "a", "abcd", true),
	Entry("IsInterleavingString #11", "abcd", "a", "aabcd", true),
	Entry("IsInterleavingString #12", "aaac", "aaab", "aaaaaabc", true),
)

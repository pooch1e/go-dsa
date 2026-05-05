package recursion

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestIsPalindrome tests solution(s) with the following signature and problem description:

	func IsPalindrome(s string) bool

Given a string return true if it's a palindrome and false otherwise.

For example given `abba` return true. Given `abca` return false.
*/
var _ = DescribeTable("IsPalindrome",
	func(s string, isPalindrome bool) {
		if got := IsPalindrome(s); got != isPalindrome {
			Expect(got).To(Equal(isPalindrome))
		}
	},
	Entry("IsPalindrome #1", "", true),
	Entry("IsPalindrome #2", "1", true),
	Entry("IsPalindrome #3", "a", true),
	Entry("IsPalindrome #4", "aa", true),
	Entry("IsPalindrome #5", "ab", false),
	Entry("IsPalindrome #6", "aba", true),
	Entry("IsPalindrome #7", "abba", true),
	Entry("IsPalindrome #8", "abca", false),
	Entry("IsPalindrome #9", "acba", false),
	Entry("IsPalindrome #10", "acca", true),
	Entry("IsPalindrome #11", "acdca", true),
	Entry("IsPalindrome #12", "acedeca", true),
	Entry("IsPalindrome #13", "acedec", false),
)

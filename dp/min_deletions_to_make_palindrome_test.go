package dp

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestMinDeletionsToMakePalindrome tests solution(s) with the following signature and problem description:

	func MinimumDeletionsToMakePalindrome(input string) int

Given a string like abccb return the minimum number of character deletions that can be done on the string
to make it a palindrome like 1 (by removing a, we will have bccb).
*/
var _ = DescribeTable("MinDeletionsToMakePalindrome",
	func(input string, minChangeNeededToMakePalindrome int) {
		if got := MinDeletionsToMakePalindrome(input); got != minChangeNeededToMakePalindrome {
			Expect(got).To(Equal(minChangeNeededToMakePalindrome))
		}
	},
	Entry("MinDeletionsToMakePalindrome #1", "", 0),
	Entry("MinDeletionsToMakePalindrome #2", "a", 0),
	Entry("MinDeletionsToMakePalindrome #3", "ab", 1),
	Entry("MinDeletionsToMakePalindrome #4", "acb", 2),
	Entry("MinDeletionsToMakePalindrome #5", "abccb", 1),
	Entry("MinDeletionsToMakePalindrome #6", "abcba", 0),
	Entry("MinDeletionsToMakePalindrome #7", "abdcba", 1),
	Entry("MinDeletionsToMakePalindrome #8", "abdecba", 2),
	Entry("MinDeletionsToMakePalindrome #9", "zabdecbaz", 2),
)

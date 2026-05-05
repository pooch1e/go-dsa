package strings

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestLongestSubstrings tests solution(s) with the following signature and problem description:

	func LongestSubstringOfTwoChars(input string) string

Given a string return the longest substring of two unique characters.

For example given "aabbc", return "aabb" because it's the longest substring that has two unique
characters "a" and "b".

Other substrings of "aabc" include:

* "aabbc", contains more than 2 unique characters.
* "bbc", shorter than "aabb".
*/
var _ = DescribeTable("LongestSubstrings",
	func(input string, longestSubstring string) {
		got := LongestSubstringOfTwoChars(input)
		if got != longestSubstring {
			Expect(got).To(Equal(longestSubstring))
		}
	},
	Entry("LongestSubstrings #1", "", ""),
	Entry("LongestSubstrings #2", "a", ""),
	Entry("LongestSubstrings #3", "ab", "ab"),
	Entry("LongestSubstrings #4", "abcc", "bcc"),
	Entry("LongestSubstrings #5", "aabbc", "aabb"),
	Entry("LongestSubstrings #6", "ada", "ada"),
	Entry("LongestSubstrings #7", "dafff", "afff"),
	Entry("LongestSubstrings #8", "abbdeeeddfffha", "deeedd"),
)

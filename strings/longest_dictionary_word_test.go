package strings

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestLongestDictionaryWordContainingKey tests solution(s) with the following signature and problem description:

	func LongestDictionaryWordContainingKey(key string, dic []string) string

Given a key as string, and a slice of strings containing a dictionary of words, return the longest
word that contains all letters of the key. Input will only contain lowercase letters a-z.

For example given "car" and {"rectify", "race", "archeology", "racoon"}, it should return "archeology",
because "archeology" is the longest word in the given set that contains "c","a",and "r".
*/
var _ = DescribeTable("LongestDictionaryWordContainingKey",
	func(input string, dictionary []string, longestWordContainingCharacters string) {
		got := LongestDictionaryWordContainingKey(input, dictionary)
		if got != longestWordContainingCharacters {
			Expect(got).To(Equal(longestWordContainingCharacters))
		}
	},
	Entry("LongestDictionaryWordContainingKey #1", "a", []string{}, ""),
	Entry("LongestDictionaryWordContainingKey #2", "a", []string{"c"}, ""),
	Entry("LongestDictionaryWordContainingKey #3", "ab", []string{"cd"}, ""),
	Entry("LongestDictionaryWordContainingKey #4", "ab", []string{"acd"}, ""),
	Entry("LongestDictionaryWordContainingKey #5", "", []string{"abc"}, "abc"),
	Entry("LongestDictionaryWordContainingKey #6", "a", []string{"a", "b", "c"}, "a"),
	Entry("LongestDictionaryWordContainingKey #7", "a", []string{"a", "ba", "c", "cc"}, "ba"),
	Entry("LongestDictionaryWordContainingKey #8", "a", []string{"a", "baa", "c"}, "baa"),
	Entry("LongestDictionaryWordContainingKey #9", "abc", []string{"abc", "aabc", "aabbcc", "bbccaaccbbaa"}, "bbccaaccbbaa"),
	Entry("LongestDictionaryWordContainingKey #10", "abc", []string{"abc", "abcdefghijklmn", "abcdefghijkl", "abcdef", "abcijkl"}, "abcdefghijklmn"),
	Entry("LongestDictionaryWordContainingKey #11", "car", []string{"rectify", "race", "archeology", "racoon"}, "archeology"),
)

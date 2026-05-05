package hashtable

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"reflect"
)

/*
TestFindAnagrams tests solution(s) with the following signature and problem description:

	func FindAnagrams(words []string) [][]string

Given a dictionary, return lists of words that are anagrams of each other. Two words are
anagrams if rearranging the letters of one results in the other.

For example given {"cat", "tac", "bar", "baz", "act"} return {"cat", "tac", "act"} because they are
the only anagrams in the list.
*/
var _ = DescribeTable("FindAnagrams",
	func(words []string, anagrams [][]string) {
		got := FindAnagrams(words)
		if !reflect.DeepEqual(got, anagrams) {
			Expect(got).To(Equal(anagrams))
		}
	},
	Entry("FindAnagrams #1", []string{"foo", "bar", "baz"}, [][]string{}),
	Entry("FindAnagrams #2", []string{"foo", "cat", "tac", "act"}, [][]string{{"cat", "tac", "act"}}),
	Entry("FindAnagrams #3", []string{"car", "cap", "arc", "tac"}, [][]string{{"car", "arc"}}),
)

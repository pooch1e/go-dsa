package tree

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"slices"
)

/*
TestAutocompletion tests solution(s) with the following signature and problem description:

	func (t *trie) AutoComplete(word string) []string

Given a dictionary of words and a word, return autocomplete suggestions from the dictionary that
start with the given word.

For example given {"car","caravan","card","carpet","cap","ca"}, and the word "car", return
{"avan","d","pet"} because they are all words that start with "car".
*/
var _ = DescribeTable("Autocompletion",
	func(input string, dict []string, suggestions []string) {
		trie := newTrie(dict)
		got := trie.Autocompletion(input)
		if !slices.Equal(got, suggestions) {
			Expect(got).To(Equal(suggestions))
		}
	},
	Entry("Autocompletion #1", "a", []string{}, []string{}),
	Entry("Autocompletion #2", "a", []string{"a", "aa", "aaa", "aaaa", "aaaaa"}, []string{"a", "aa", "aaa", "aaaa"}),
	Entry("Autocompletion #3", "car", []string{"car", "caravan", "card", "carpet", "cap", "ca"}, []string{"avan", "d", "pet"}),
	Entry("Autocompletion #4", "a", []string{"aaaaa", "aaaa", "aaa", "aa", "a"}, []string{"a", "aa", "aaa", "aaaa"}),
	Entry("Autocompletion #5", "a", []string{"abc", "abd", "abe", "bbb"}, []string{"bc", "bd", "be"}),
	Entry("Autocompletion #6", "a", []string{"abc", "abd", "abe"}, []string{"bc", "bd", "be"}),
	Entry("Autocompletion #7", "a", []string{"aac", "abc", "acc"}, []string{"ac", "bc", "cc"}),
	Entry("Autocompletion #8", "ab", []string{"abcdfg", "abdefg", "abefgh"}, []string{"cdfg", "defg", "efgh"}),
	Entry("Autocompletion #9", "ab", []string{"abc", "abd", "abe"}, []string{"c", "d", "e"}),
	Entry("Autocompletion #10", "ab", []string{"abc", "abcc", "abccd", "abd", "abe"}, []string{"c", "cc", "ccd", "d", "e"}),
	Entry("Autocompletion #11", "ab", []string{"abc", "abcc", "abcd", "abd", "abe"}, []string{"c", "cc", "cd", "d", "e"}),
	Entry("Autocompletion #12", "ab", []string{"abc", "abcd", "abd", "abcc", "abe"}, []string{"c", "cc", "cd", "d", "e"}),
	Entry("Autocompletion #13", "car", []string{"car", "caravan", "card", "carpet", "cap", "ca"}, []string{"avan", "d", "pet"}),
)

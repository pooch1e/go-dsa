package queue

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"slices"
)

/*
TestStringPermutations tests solution(s) with the following signature and problem description:

	func StringPermutations(input string) []string

Given a string return all possible permutations that can be made by rearranging the letters in the string
using a queue.

For example given "abc" return "abc,acb,bac,bca,cab,cba".
*/
var _ = DescribeTable("StringPermutations",
	func(input string, permutations []string) {
		if got := StringPermutations(input); !slices.Equal(got, permutations) {
			Expect(got).To(Equal(permutations))
		}
	},
	Entry("StringPermutations #1", "", []string{""}),
	Entry("StringPermutations #2", "a", []string{"a"}),
	Entry("StringPermutations #3", "ab", []string{"ab", "ba"}),
	Entry("StringPermutations #4", "ba", []string{"ba", "ab"}),
	Entry("StringPermutations #5", "abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"}),
	Entry("StringPermutations #6", "123", []string{"123", "132", "213", "231", "312", "321"}),
)

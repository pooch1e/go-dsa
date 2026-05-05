package strings

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestReverseVowels tests solution(s) with the following signature and problem description:

	func ReverseVowels(str string) (string, error)

Given a string, return the same string while reversing the vowels {"a", "e", "i", "o", "u"}
appear in it.

For example given "coat", return "caot", because the vowels are "o" and "a" and their positions
are reversed.
*/
var _ = DescribeTable("ReverseVowels",
	func(word string, reversed string) {
		got, err := ReverseVowels(word)
		if err != nil {
			Expect(err).ToNot(HaveOccurred())
		}
		if got != reversed {
			Expect(got).To(Equal(reversed))
		}
	},
	Entry("ReverseVowels #1", "umbrella", "ambrellu"),
	Entry("ReverseVowels #2", "coat", "caot"),
	Entry("ReverseVowels #3", "eventually", "avunteelly"),
	Entry("ReverseVowels #4", "sooner rather than later", "seanar rethar then lotor"),
)

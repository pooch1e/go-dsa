package dp

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestWordDistance tests solution(s) with the following signature and problem description:

	func WordDistance(input1, input2 string) int

Given a string like abc, and another string like abcde return how many character
modifications (insert, delete, edit) have to be done on the first string to become
identical to the second string. In this case, the answer is 2.
*/
var _ = DescribeTable("WordDistance",
	func(input1 string, input2 string, distance int) {
		if got := WordDistance(input1, input2); got != distance {
			Expect(got).To(Equal(distance))
		}
	},
	Entry("WordDistance #1", "", "", 0),
	Entry("WordDistance #2", "", "abcde", 5),
	Entry("WordDistance #3", "a", "a", 0),
	Entry("WordDistance #4", "a", "b", 1),
	Entry("WordDistance #5", "ab", "ac", 1),
	Entry("WordDistance #6", "ab", "dc", 2),
	Entry("WordDistance #7", "ab", "dcd", 3),
	Entry("WordDistance #8", "abc", "abcde", 2),
	Entry("WordDistance #9", "abcdef", "abcde", 1),
	Entry("WordDistance #10", "gabcdef", "abcde", 2),
)

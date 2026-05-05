package strings

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestFindDuplicate tests solution(s) with the following signature and problem description:

	func LookAndTell(depth int) []string

Given a positive integer n, return the output of the Look and Tell algorithm until the nth depth.

The Look and Tell algorithm starts by outputting 1 at first level, then at each subsequent level
it reads the previous line by counting the number of times a digit is repeated and then writes
the count and the digit.

For example given 4, return:

1
11
21
1211

Which reads:
one
one one
two ones
one two one one.

The third output (two ones) reads the level before it which is 11. Two ones means repeat
1 two times i.e. 11.
*/
var _ = DescribeTable("FindDuplicate",
	func(depth int, lastLine string) {
		tell := LookAndTell(depth)
		got := tell[len(tell)-1]
		if got != lastLine {
			Expect(got).To(Equal(lastLine))
		}
	},
	Entry("FindDuplicate #1", 0, "-1"),
	Entry("FindDuplicate #2", 1, "1"),
	Entry("FindDuplicate #3", 2, "11"),
	Entry("FindDuplicate #4", 3, "21"),
	Entry("FindDuplicate #5", 4, "1211"),
	Entry("FindDuplicate #6", 5, "111221"),
	Entry("FindDuplicate #7", 6, "312211"),
	Entry("FindDuplicate #8", 7, "13112221"),
	Entry("FindDuplicate #9", 8, "1113213211"),
	Entry("FindDuplicate #10", 9, "31131211131221"),
	Entry("FindDuplicate #11", 10, "13211311123113112211"),
)

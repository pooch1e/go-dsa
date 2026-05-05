package queue

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestGenerateBinaryNumbers tests solution(s) with the following signature and problem description:

	func GenerateBinaryNumbers(n int) []string

Given a positive integer like n count from 0 to n in binary.

For example given 3 return {"0", "1", "10", "11"}.
*/
var _ = DescribeTable("GenerateBinaryNumbers",
	func(n int, lastLine string) {
		binaryNumbers := GenerateBinaryNumbers(n)
		got := binaryNumbers[len(binaryNumbers)-1]
		if got != lastLine {
			Expect(got).To(Equal(lastLine))
		}
	},
	Entry("GenerateBinaryNumbers #1", 0, "0"),
	Entry("GenerateBinaryNumbers #2", 1, "1"),
	Entry("GenerateBinaryNumbers #3", 2, "10"),
	Entry("GenerateBinaryNumbers #4", 3, "11"),
	Entry("GenerateBinaryNumbers #5", 4, "100"),
	Entry("GenerateBinaryNumbers #6", 5, "101"),
	Entry("GenerateBinaryNumbers #7", 6, "110"),
	Entry("GenerateBinaryNumbers #8", 7, "111"),
	Entry("GenerateBinaryNumbers #9", 8, "1000"),
	Entry("GenerateBinaryNumbers #10", 9, "1001"),
	Entry("GenerateBinaryNumbers #11", 10, "1010"),
)

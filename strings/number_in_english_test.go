package strings

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestReadNumberInEnglish tests solution(s) with the following signature and problem description:

	func NumberInEnglish(num int) string

Given n a positive integer smaller than a Trillion, return the number in English words.

For example given 0, return "Zero".
For example given 34, return "Thirty Four".
For example given 10, return "Ten".
For example given 900000000001, return "Nine Hundred Billion One".
*/
var _ = DescribeTable("ReadNumberInEnglish",
	func(number int, english string) {
		got := NumberInEnglish(number)
		if got != english {
			Expect(got).To(Equal(english))
		}
	},
	Entry("ReadNumberInEnglish #1", 0, "Zero"),
	Entry("ReadNumberInEnglish #2", 1, "One"),
	Entry("ReadNumberInEnglish #3", 2, "Two"),
	Entry("ReadNumberInEnglish #4", 10, "Ten"),
	Entry("ReadNumberInEnglish #5", 34, "Thirty Four"),
	Entry("ReadNumberInEnglish #6", 123456789, "One Hundred Twenty Three Million Four Hundred Fifty Six Thousand Seven Hundred Eighty Nine"),
	Entry("ReadNumberInEnglish #7", 1000000000, "One Billion"),
	Entry("ReadNumberInEnglish #8", 100000000000, "One Hundred Billion"),
	Entry("ReadNumberInEnglish #9", 900000000001, "Nine Hundred Billion One"),
)

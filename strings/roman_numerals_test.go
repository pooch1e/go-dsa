package strings

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestIntToRoman tests solution(s) with the following signature and problem description:

	func IntToRoman(input int) string

Given a positive integer like return the equivalent inRoman numerals:

For example:

* Given 1, return I
* Given 4, return IV
* Given 5, return V
* Given 9, return IX
* Given 10, return X
* Given 40, return XL
* Given 50, return L
* Given 90, return XC
* Given 100, return C
* Given 400, return CD
* Given 500, return D
* Given 900, return CM
* Given 1000, return M
* Given 1917, return MCMXVII.
*/
var _ = DescribeTable("IntToRoman",
	func(number int, roman string) {
		if got := IntToRoman(number); got != roman {
			Expect(got).To(Equal(roman))
		}
	},
	Entry("IntToRoman #1", 0, ""),
	Entry("IntToRoman #2", 1, "I"),
	Entry("IntToRoman #3", 2, "II"),
	Entry("IntToRoman #4", 3, "III"),
	Entry("IntToRoman #5", 4, "IV"),
	Entry("IntToRoman #6", 5, "V"),
	Entry("IntToRoman #7", 6, "VI"),
	Entry("IntToRoman #8", 58, "LVIII"),
	Entry("IntToRoman #9", 60, "LX"),
	Entry("IntToRoman #10", 1000, "M"),
	Entry("IntToRoman #11", 1917, "MCMXVII"),
)

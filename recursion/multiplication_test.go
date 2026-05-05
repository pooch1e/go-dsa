package recursion

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestMultiplication tests solution(s) with the following signature and problem description:

	func Multiply(a, b int) int

Given two integers, return their product using recursion and without using the
multiplication operator.

For example given 2 and 3 return 6.
*/
var _ = DescribeTable("Multiplication",
	func(a int, b int, product int) {
		if got := Multiply(a, b); got != product {
			Expect(got).To(Equal(product))
		}
	},
	Entry("Multiplication #1", 0, 0, 0),
	Entry("Multiplication #2", 0, 1, 0),
	Entry("Multiplication #3", 0, 1, 0),
	Entry("Multiplication #4", 1, 1, 1),
	Entry("Multiplication #5", 1, 2, 2),
	Entry("Multiplication #6", 2, 2, 4),
	Entry("Multiplication #7", 3, 3, 9),
	Entry("Multiplication #8", 3, -3, -9),
	Entry("Multiplication #9", -3, -3, 9),
)

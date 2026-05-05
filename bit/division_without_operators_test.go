package bit

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestDivision tests solution(s) with the following signature and problem description:

	func Divide(x, y int) int

Given two integers, divide them without using the built-in `/` or `*` operators.

For example given 20 and 4 return 5.
*/
var _ = DescribeTable("Division",
	func(a int, b int) {
		got := Divide(a, b)
		want := a / b
		if got != want {
			Expect(got).To(Equal(want))
		}
	},
	Entry("Division #1", 0, 1),
	Entry("Division #2", 20, 2),
	Entry("Division #3", 20, 4),
	Entry("Division #4", 20, 3),
)

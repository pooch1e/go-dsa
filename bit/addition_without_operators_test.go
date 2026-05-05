package bit

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestAdditionWithoutArithmeticOperators tests solution(s) with the following signature and problem description:

	func Add(x, y int) int

Given two integers add them without using any arithmetic operators such as {+,-,/,*,++,--,+=,…}.

For example given 2 and 3 return 5.
*/
var _ = DescribeTable("AdditionWithoutArithmeticOperators",
	func(a int, b int) {
		got := AdditionWithoutArithmeticOperators(a, b)
		want := a + b
		if got != want {
			Expect(got).To(Equal(want))
		}
	},
	Entry("AdditionWithoutArithmeticOperators #1", 0, 1),
	Entry("AdditionWithoutArithmeticOperators #2", 20, 2),
	Entry("AdditionWithoutArithmeticOperators #3", 20, 4),
	Entry("AdditionWithoutArithmeticOperators #4", 20, 3),
)

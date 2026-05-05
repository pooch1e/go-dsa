package bit

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestIsPowerOfTwo tests solution(s) with the following signature and problem description:

	func IsPowerOfTwo(input int) bool

Using bit manipulation, return true if a given number is a power of 2 and false otherwise.

For example given 20 return false. Given 256 return true because 2 ^ 8 = 256.
*/
var _ = DescribeTable("IsPowerOfTwo",
	func(input int, isPowerOfTwo bool) {
		if got := IsPowerOfTwo(input); got != isPowerOfTwo {
			Expect(got).To(Equal(isPowerOfTwo))
		}
	},
	Entry("IsPowerOfTwo #1", 0, false),
	Entry("IsPowerOfTwo #2", 1, true),
	Entry("IsPowerOfTwo #3", 2, true),
	Entry("IsPowerOfTwo #4", 3, false),
	Entry("IsPowerOfTwo #5", 4, true),
	Entry("IsPowerOfTwo #6", 20, false),
	Entry("IsPowerOfTwo #7", 256, true),
)

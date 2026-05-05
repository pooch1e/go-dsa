package dnc

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"math"
)

func floatAlmostEquals(a, b float64) bool {
	return math.Abs(a-b) <= 1e9 // Equality threshold
}

/*
TestSquareRoot tests solution(s) with the following signature and problem description:

	func SquareRoot(number, precision int) float64

Given a number and precision, return the square root of the number using the binary
search algorithm.

For example given 9 and 3, it should return 3.
*/
var _ = DescribeTable("SquareRoot",
	func(number int, precision int, solution float64) {
		if got := SquareRoot(number, precision); !floatAlmostEquals(got, solution) {
			Expect(got).To(Equal(solution))
		}
	},
	Entry("SquareRoot #1", 0, 0, float64(1)),
	Entry("SquareRoot #2", 1, 0, float64(1)),
	Entry("SquareRoot #3", 1, 1, float64(1)),
	Entry("SquareRoot #4", 4, 1, float64(2)),
	Entry("SquareRoot #5", 4, 2, 1.9999999999999998),
	Entry("SquareRoot #6", 4, 3, float64(2)),
	Entry("SquareRoot #7", 5, 3, float64(2)),
	Entry("SquareRoot #8", 9, 3, float64(3)),
)

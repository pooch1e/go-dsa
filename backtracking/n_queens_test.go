package backtracking

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"reflect"
)

/*
TestNQueens tests solution(s) with the following signature and problem description:

	func NQueens(n int) []Chessboard

Given an integer n representing an n by n chessboard, return all possible
arrangements of placing n queens on the chessboard such that the queens do not
attack each other.
*/
var _ = DescribeTable("NQueens",
	func(n int, solutions []Chessboard) {
		if got := NQueens(n); !reflect.DeepEqual(solutions, got) {
			Expect(got).To(Equal(solutions))
		}
	},
	Entry("NQueens #1", 0, []Chessboard{}),
	Entry("NQueens #2", 1, []Chessboard{{{queen}}}),
	Entry("NQueens #3", 2, []Chessboard{}),
	Entry("NQueens #4", 3, []Chessboard{}),
	Entry("NQueens #5", 4, []Chessboard{{{empty, queen, empty, empty}, {empty, empty, empty, queen}, {queen, empty, empty, empty}, {empty, empty, queen, empty}}, {{empty, empty, queen, empty}, {queen, empty, empty, empty}, {empty, empty, empty, queen}, {empty, queen, empty, empty}}}),
)

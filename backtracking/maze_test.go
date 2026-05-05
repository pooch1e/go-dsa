package backtracking

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestMaze tests solution(s) with the following signature and problem description:

	func Maze(walls [][2]int, start, finish [2]int,m,n int) string

Given the coordinates of walls in a m x n maze in tuples of {row, col} format, and start
and finish coordinates in the same format, find a path from start to finish. return
a string of directions like `lrud` (left, right, up, down) to get a robot from
start to finish.

The robot can only move in the four left, right, down and up directions and not
through walls. The robot does not know where the finish line is so it has to
explore every possible cell in the order of directions given.
*/
var _ = DescribeTable("Maze",
	func(m int, n int, walls [][2]int, start [2]int, finish [2]int, moves string) {
		if got := Maze(m, n, walls, start, finish); moves != got {
			Expect(got).To(Equal(moves))
		}
	},
	Entry("Maze #1", 1, 1, [][2]int{}, [2]int{0, 0}, [2]int{0, 1}, ""),
	Entry("Maze #2", 5, 5, [][2]int{}, [2]int{0, 0}, [2]int{0, 1}, "r"),
	Entry("Maze #3", 10, 10, [][2]int{}, [2]int{0, 0}, [2]int{0, 1}, "r"),
	Entry("Maze #4", 5, 5, [][2]int{}, [2]int{0, 0}, [2]int{4, 4}, "rrrrdlllldrrrrdlllldrrrr"),
	Entry("Maze #5", 5, 5, [][2]int{{1, 1}, {1, 2}, {1, 3}, {2, 3}, {3, 3}, {3, 4}}, [2]int{0, 0}, [2]int{2, 4}, "rrrrdd"),
	Entry("Maze #6", 5, 5, [][2]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {1, 3}, {2, 1}, {2, 2}, {2, 3}, {3, 1}, {4, 1}}, [2]int{4, 0}, [2]int{1, 2}, "uuurr"),
	Entry("Maze #7", 5, 5, [][2]int{{1, 0}, {1, 1}, {1, 2}, {1, 3}, {3, 1}, {3, 2}, {3, 3}, {3, 4}}, [2]int{0, 0}, [2]int{4, 4}, "rrrrddllllddrrrr"),
	Entry("Maze #8", 5, 5, [][2]int{{1, 0}, {1, 1}, {1, 4}, {1, 3}, {3, 1}, {3, 2}, {3, 3}, {3, 4}}, [2]int{0, 0}, [2]int{4, 4}, "rrddllddrrrr"),
)

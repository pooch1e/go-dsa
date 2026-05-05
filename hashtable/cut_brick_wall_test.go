package hashtable

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestCutBrickWall tests solution(s) with the following signature and problem description:

	func CutBrickWall(wall [][]int) int

Given a two dimensional array of numbers representing the length of bricks of equal hight in each
row of a wall, return the position at which we could do the least amount of cuts in order
to have a rectangular wall.

For example given {1, 2, 2}, {3, 2}, {2, 3}, return 3.

	┌─────┬───────────┬───────────┐
	│  1  │     2     │     2     │
	├─────┴───────────┼───────────┤
	│        3        │     2     │
	├───────────┬─────┴───────────┤
	│     2     │         3       │
	└───────────┴─────────────────┘

	┌─────┬───────────┐
	│  1  │     2     │
	├─────┴───────────┤
	│        3        │
	├───────────┬─────┤
	│     2     │  1  │
	└───────────┴─────┘

We can have the above wall after cutting only one brick at position 2 row 3.
*/
var _ = DescribeTable("CutBrickWall",
	func(cutPosition int, bricks [][]int) {
		got := CutBrickWall(bricks)
		if got != cutPosition {
			Expect(got).To(Equal(cutPosition))
		}
	},
	Entry("CutBrickWall #1", 0, [][]int{}),
	Entry("CutBrickWall #2", 3, [][]int{{1, 2, 2}, {3, 2}, {2, 3}}),
	Entry("CutBrickWall #3", 5, [][]int{{5, 4, 1}, {2, 2, 1, 3, 2}, {3, 2, 4, 1}, {3, 4, 3}, {2, 3, 5}, {3, 2, 1, 4}}),
	Entry("CutBrickWall #4", 1, [][]int{{2, 5, 4, 1}, {1, 2, 2, 1, 3, 2}, {1, 3, 2, 4, 1}, {1, 3, 4, 3}}),
	Entry("CutBrickWall #5", 1, [][]int{{2, 5, 4, 1}, {1, 2, 2, 1, 3, 2}, {1, 3, 2, 4, 1}, {1, 3, 4, 3}}),
)

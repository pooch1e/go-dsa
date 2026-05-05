package dnc

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"slices"
)

/*
TestTowerOfHanoi tests solution(s) with the following signature and problem description:

	func TowerOfHanoi(n, start, end int) [][2]int

Given n, number of disks, and start and end tower, return the number of moves it takes to
move all disks from start to end tower. The disks are stacked on top of each other with the
lightest being on top and heaviest being in the bottom. A heavier disk cannot be placed
on a lighter disk. You can move one disk at a time.

For example given 2 disks, start of 1 and end of 3 return {1, 2}, {1, 3}, {2, 3}.
The first move is to move the top disk from tower 1 to tower 2.
The second move is to move the top disk from tower 1 to tower 3.
The final move is to move the top disk from 2 to 3.
By the end all disks are moved from tower 1 to tower 3.
*/
var _ = DescribeTable("TowerOfHanoi",
	func(n int, start int, end int, moves [][2]int) {
		if got := TowerOfHanoi(n, start, end); !slices.Equal(got, moves) {
			Expect(got).To(Equal(moves))
		}
	},
	Entry("TowerOfHanoi #1", 0, 1, 1, [][2]int{{1, 1}}),
	Entry("TowerOfHanoi #2", 1, 1, 1, [][2]int{{1, 1}}),
	Entry("TowerOfHanoi #3", 1, 1, 2, [][2]int{{1, 2}}),
	Entry("TowerOfHanoi #4", 2, 1, 3, [][2]int{{1, 2}, {1, 3}, {2, 3}}),
	Entry("TowerOfHanoi #5", 3, 1, 3, [][2]int{{1, 3}, {1, 2}, {3, 2}, {1, 3}, {2, 1}, {2, 3}, {1, 3}}),
	Entry("TowerOfHanoi #6", 4, 1, 3, [][2]int{{1, 2}, {1, 3}, {2, 3}, {1, 2}, {3, 1}, {3, 2}, {1, 2}, {1, 3}, {2, 3}, {2, 1}, {3, 1}, {2, 3}, {1, 2}, {1, 3}, {2, 3}}),
)

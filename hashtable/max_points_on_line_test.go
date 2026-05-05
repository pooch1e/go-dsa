package hashtable

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestMaxPointsOnALine tests solution(s) with the following signature and problem description:

	func MaxPointsOnALine(points [][]int) int

Given multiple coordinates of points like, return the maximum number of points that are on the same.

For example given {[1,1], [2,2], [3,3], [4,5]} return 3 because the points [1,1], [2,2], [3,3] are on
the same line.
*/
var _ = DescribeTable("MaxPointsOnALine",
	func(points [][]int, maxLines int) {
		got := MaxPointsOnALine(points)
		if got != maxLines {
			Expect(got).To(Equal(maxLines))
		}
	},
	Entry("MaxPointsOnALine #1", [][]int{}, 0),
	Entry("MaxPointsOnALine #2", [][]int{{2, 5}}, 1),
	Entry("MaxPointsOnALine #3", [][]int{{2, 5}, {4, 1}}, 2),
	Entry("MaxPointsOnALine #4", [][]int{{2, 5}, {2, 2}, {1, 1}, {7, 7}, {1, 5}}, 3),
	Entry("MaxPointsOnALine #5", [][]int{{2, 5}, {2, 2}, {2, 2}, {1, 1}, {7, 7}, {1, 5}}, 4),
	Entry("MaxPointsOnALine #6", [][]int{{-2, 1}, {-3, 2}, {-4, 3}, {-5, 4}, {-1, 1}, {-5, 3}}, 4),
	Entry("MaxPointsOnALine #7", [][]int{{2, 1}, {3, 2}, {5, 4}, {2, 3}, {3, 4}, {4, 5}}, 3),
)

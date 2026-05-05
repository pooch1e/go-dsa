package heap

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"reflect"
)

/*
TestKClosestPointToOrigin tests solution(s) with the following signature and problem description:

	func KClosestPointToOrigin(points [][]int, k int) [][]int

Given coordination of a point on an x,y axis and an integer k, return k closest points to the origin.

For example given {1,1}, {2,2}, {3,3}, {4,4} and k=2, return {1,1}, {2,2}.
*/
var _ = DescribeTable("KClosestPointToOrigin",
	func(points [][]int, k int, kthClosest [][]int) {
		got := KClosestPointToOrigin(points, k)
		if !reflect.DeepEqual(got, kthClosest) {
			Expect(got).To(Equal(kthClosest))
		}
	},
	Entry("KClosestPointToOrigin #1", [][]int{{}}, 1, [][]int{{}}),
	Entry("KClosestPointToOrigin #2", [][]int{{10, 10}}, 1, [][]int{{10, 10}}),
	Entry("KClosestPointToOrigin #3", [][]int{{10, 10}, {-10, -10}, {11, 11}, {12, 12}}, 1, [][]int{{10, 10}}),
	Entry("KClosestPointToOrigin #4", [][]int{{10, 10}, {-10, -10}, {11, 11}, {12, 12}}, 2, [][]int{{10, 10}, {-10, -10}}),
	Entry("KClosestPointToOrigin #5", [][]int{{35, 11}, {14, -12}, {50, 100}, {2, 19}, {-17, -18}, {1, 1}, {27, 22}}, 1, [][]int{{1, 1}}),
)

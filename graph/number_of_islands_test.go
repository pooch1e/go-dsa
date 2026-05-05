package graph

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestNumberOfIslands tests solution(s) with the following signature and problem description:

	func NumberOfIslands(grid [][]int) int

Given a grid in which 0 represents water and 1 represents a piece of land, return the
number of islands. An Island is one or more piece of land that could be connected to other
pieces of land on left, right, top or bottom.
*/
var _ = DescribeTable("NumberOfIslands",
	func(grid [][]int, islandCount int) {
		if got := NumberOfIslands(grid); got != islandCount {
			Expect(got).To(Equal(islandCount))
		}
	},
	Entry("NumberOfIslands #1", [][]int{}, 0),
	Entry("NumberOfIslands #2", [][]int{{1, 0, 0}, {0, 0, 0}}, 1),
	Entry("NumberOfIslands #3", [][]int{{1, 0, 1}, {0, 0, 0}}, 2),
	Entry("NumberOfIslands #4", [][]int{{1, 0, 1}, {0, 0, 0}, {0, 1, 0}}, 3),
	Entry("NumberOfIslands #5", [][]int{{1, 0, 1}, {0, 0, 0}, {0, 1, 1}}, 3),
	Entry("NumberOfIslands #6", [][]int{{1, 0, 1}, {0, 0, 1}, {0, 1, 1}}, 2),
)

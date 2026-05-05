package graph

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestIsDAG tests solution(s) with the following signature and problem description:

	Vertex struct {
			Val int
			Edges []*Vertex
	}
	func IsDAG(graph []*Vertex) bool

Given a directed graph determine if it's a DAG or not. A directed acyclic graph (DAG)
is a directed graph that has no cycles.
*/
var _ = DescribeTable("IsDAG",
	func(graph [][]int, isDAG bool) {
		if got := IsDAG(toGraph(graph)); got != isDAG {
			Expect(got).To(Equal(isDAG))
		}
	},
	Entry("IsDAG #1", [][]int{{2}, {1, 3}, {}}, false),
	Entry("IsDAG #2", [][]int{{2, 3}, {3}, {4}, {}}, true),
	Entry("IsDAG #3", [][]int{{2, 3}, {4, 5}, {}, {}, {}}, true),
	Entry("IsDAG #4", [][]int{{2, 3, 4}, {3}, {4}, {5}, {}}, true),
	Entry("IsDAG #5", [][]int{{2, 3, 4}, {3, 5}, {5}, {3, 5}, {}}, true),
	Entry("IsDAG #6", [][]int{{2, 3, 4}, {3, 5}, {5}, {3, 5}, {5, 3}}, false),
	Entry("IsDAG #7", readmeGraphs["Figure_1_A"], true),
	Entry("IsDAG #8", readmeGraphs["Figure_1_B"], false),
	Entry("IsDAG #9", readmeGraphs["Figure_1_C"], true),
)

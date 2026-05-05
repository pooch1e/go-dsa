package graph

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"errors"
	"slices"
)

var readmeGraphs = map[string][][]int{
	"Figure_1_A": {{2, 3, 4}, {3}, {5}, {3}, {}},
	"Figure_1_B": {{2}, {3}, {4}, {1}},
	"Figure_1_C": {{2}, {4}, {2}, {5}, {}},
}

/*
TestTopologicalSort tests solution(s) with the following signature and problem description:

	type VertexWithIngress struct {
		Val any
		Edges []*VertexWithIngress
		Ingress int
	}
	func TopologicalSort(graph []*VertexWithIngress) ([]int, error)

In a DAG the topological order returns elements such that if there's a path from v(i) to
v(j), then v(i) appears before v(j) in the ordering. Given a graph like (B) of _Figure 1_
output the graph in topological order like {1,2,3,4}.
*/
var _ = DescribeTable("TopologicalSort",
	func(graph [][]int, topologicalSort []int, expectedErr error) {
		orderedGraph, err := TopologicalSort(toGraphWithIngress(graph))
		got := make([]int, 0, len(orderedGraph))
		for i := range orderedGraph {
			got = append(got, orderedGraph[i].Val.(int))
		}
		if err != nil {
			if expectedErr == nil {
				Expect(err).ToNot(HaveOccurred())
			}
			if !errors.Is(err, expectedErr) {
				Expect(expectedErr, err).ToNot(HaveOccurred())
			}
		}
		if err == nil && expectedErr != nil {
			Expect(expectedErr).ToNot(HaveOccurred())
		}
		if !slices.Equal(got, topologicalSort) {
			Expect(got).To(Equal(topologicalSort))
		}
	},
	Entry("TopologicalSort #1", [][]int{}, []int{}, nil),
	Entry("TopologicalSort #2", [][]int{{}}, []int{1}, nil),
	Entry("TopologicalSort #3", [][]int{{}, {}}, []int{1, 2}, nil),
	Entry("TopologicalSort #4", [][]int{{}, {}, {4}, {}}, []int{1, 2, 3, 4}, nil),
	Entry("TopologicalSort #5", [][]int{{}, {1}, {}, {3}}, []int{2, 4, 1, 3}, nil),
	Entry("TopologicalSort #6", [][]int{{}, {1}, {}, {3}}, []int{2, 4, 1, 3}, nil),
	Entry("TopologicalSort #7", [][]int{{2, 5}, {3, 4}, {}, {5}, {}}, []int{1, 2, 3, 4, 5}, nil),
	Entry("TopologicalSort #8", [][]int{{2}, {3}, {1}}, []int{}, ErrNotADAG),
	Entry("TopologicalSort #9", readmeGraphs["Figure_1_A"], []int{1, 2, 4, 3, 5}, nil),
	Entry("TopologicalSort #10", readmeGraphs["Figure_1_B"], []int{}, ErrNotADAG),
	Entry("TopologicalSort #11", readmeGraphs["Figure_1_C"], []int{1, 3, 2, 4, 5}, nil),
)

func toGraphWithIngress(graph [][]int) []*VertexWithIngress {
	graphMap := make(map[int]*VertexWithIngress)

	for i := range graph {
		graphMap[i+1] = &VertexWithIngress{Val: i + 1, Edges: []*VertexWithIngress{}}
	}

	for i, v := range graph {
		graphMap[i+1].Edges = make([]*VertexWithIngress, len(v))
		for j, e := range v {
			graphMap[i+1].Edges[j] = graphMap[e]
		}
	}

	output := make([]*VertexWithIngress, len(graphMap))
	for i := range len(graph) {
		output[i] = graphMap[i+1]
	}
	return output
}

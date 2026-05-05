package graph

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"errors"
	"strings"
)

const (
	lineBreak = "\n"
	delimiter = ","
)

/*
TestDependencyOrder tests solution(s) with the following signature and problem description:

	func DependencyOrder(graph []*VertexWithIngress)  []*VertexWithIngress

Given a multiline input, each line starting with  an element and followed by a comma separated
list of dependent elements, output the elements in the order in which the dependencies are met.

In real world this is similar to courses and their prerequisite where a prerequisite must be
taken before a course can be taken. This is also similar to the dependencies of a software,
where there are dependencies of dependencies and a software can only built when all of its
dependencies are built.

For example the following input:
a,b,c,d
b,c
d,e,f

Means that a depends on b, c, d; b depends on c; d depends on e, f.
The output should be f,e,c,d,b,a. Note that the order of the elements in the output is not
important, but the order of the dependencies is important.
*/
var _ = DescribeTable("DependencyOrder",
	func(input string, order string, expectedErr error) {
		orderedGraph, err := DependencyOrder(toDependencyGraph(input))
		got := ""
		for i := range orderedGraph {
			got += orderedGraph[i].Val.(string) + delimiter
		}
		if len(orderedGraph) > 0 {
			got = strings.TrimSuffix(got, delimiter)
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
		if got != order {
			Expect(got).To(Equal(order))
		}
	},
	Entry("DependencyOrder #1", "", "", nil),
	Entry("DependencyOrder #2", "a,b", "b,a", nil),
	Entry("DependencyOrder #3", "a,b,c\nc,a", "", ErrNotADAG),
	Entry("DependencyOrder #4", "a,b,c,d\nb,c\nd,e,f", "f,e,c,d,b,a", nil),
	Entry("DependencyOrder #5", "a,f\nb,f\nc,f\nd,f\ne,f", "f,e,d,c,b,a", nil),
	Entry("DependencyOrder #6", "a,f\nb,f\nc,f\nd,f\ne,f\nf,g\ng,h", "h,g,f,e,d,c,b,a", nil),
	Entry("DependencyOrder #7", "a,f\nb,f\nc,f\nd,f\ne,f\nf,g\ng,h\nd,h", "h,g,f,e,d,c,b,a", nil),
)

func toDependencyGraph(input string) []*VertexWithIngress {
	graph := []*VertexWithIngress{}
	graphMap := make(map[string]*VertexWithIngress)
	for _, line := range strings.Split(input, lineBreak) {
		firstElement := ""
		for _, element := range strings.Split(line, delimiter) {
			if _, ok := graphMap[element]; !ok {
				vertex := &VertexWithIngress{Val: element, Edges: []*VertexWithIngress{}}
				graphMap[element] = vertex
				graph = append(graph, vertex)
			}

			if firstElement == "" {
				firstElement = element
				continue
			}
			graphMap[firstElement].Edges = append(graphMap[firstElement].Edges, graphMap[element])
		}
	}
	return graph
}

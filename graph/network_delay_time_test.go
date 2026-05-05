package graph

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestNetworkDelayTime tests solution(s) with the following signature and problem description:

	func NetworkDelayTime(n, k int, edges [][3]int) int

Given `n`, the number of nodes in a network, a list of unidirectional travel times in
`{source, destination, delay}` format, and `k`, an origin node number, return the time it will
take for a message sent from k to be received by all other nodes.
*/
var _ = DescribeTable("NetworkDelayTime",
	func(edges [][3]int, numberOfNodes int, messageFrom int, shortestPath int) {
		if got := NetworkDelayTime(numberOfNodes, messageFrom, edges); got != shortestPath {
			Expect(got).To(Equal(shortestPath))
		}
	},
	Entry("NetworkDelayTime #1", [][3]int{{0, 1, 1}}, 2, 1, -1),
	Entry("NetworkDelayTime #2", [][3]int{{0, 1, 1}, {1, 2, 2}, {2, 3, 2}, {3, 1, 3}}, 4, 0, 5),
	Entry("NetworkDelayTime #3", [][3]int{{0, 1, 1}, {1, 2, 2}, {2, 3, 2}, {3, 1, 3}}, 4, 1, -1),
	Entry("NetworkDelayTime #4", [][3]int{{0, 1, 2}, {0, 2, 3}, {2, 4, 2}, {1, 3, 2}, {3, 4, 3}, {4, 5, 2}}, 5, 0, 7),
	Entry("NetworkDelayTime #5", [][3]int{{0, 1, 2}, {0, 2, 3}, {2, 4, 2}, {1, 3, 2}, {3, 4, 3}, {4, 5, 2}, {5, 0, 1}}, 5, 5, 5),
)

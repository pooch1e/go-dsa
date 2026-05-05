package graph

type edgeMinHeap [][]int

const (
	vertexSource, vertexDestination, vertexDelay = 0, 1, 2
	edgeDestination, edgeCost                    = 0, 1
)

// NetworkDelayTime solves the problem in O(n*Log n) time and O(n) space.
func NetworkDelayTime(n, k int, edges [][3]int) int {
	panic("not implemented")

}

// verticesAndEdges takes edges, and returns
// a map of vertex ID as keys and the outgoing edges in the [vertexDestination, cost] form
// a minimum heap of edges in the [vertexDestination, cost].
func verticesAndEdges(edges [][3]int, k int) (map[int][][2]int, edgeMinHeap) {
	panic("not implemented")

}

func (e edgeMinHeap) Len() int           { panic("not implemented") }
func (e edgeMinHeap) Less(i, j int) bool { panic("not implemented") }
func (e edgeMinHeap) Swap(i, j int)      { panic("not implemented") }
func (e *edgeMinHeap) Push(x any)        { panic("not implemented") }
func (e *edgeMinHeap) Pop() any {
	panic("not implemented")

}

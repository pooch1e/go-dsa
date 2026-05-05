package graph

type (
	dijkstraVertex struct {
		val        int
		distance   int
		discovered bool
		edges      []*weightedEdge
	}
	weightedEdge struct {
		edge   *dijkstraVertex
		weight int
	}
	verticesHeap []*dijkstraVertex
)

// Dijkstra solves the problem in O(ElogV) time and O(V) space.
func Dijkstra(graph []*dijkstraVertex, source *dijkstraVertex) {
	panic("not implemented")

}

func (p verticesHeap) Len() int           { panic("not implemented") }
func (p verticesHeap) Less(i, j int) bool { panic("not implemented") }
func (p verticesHeap) Swap(i, j int)      { panic("not implemented") }
func (p *verticesHeap) Push(x any)        { panic("not implemented") }
func (p *verticesHeap) Pop() any {
	panic("not implemented")

}

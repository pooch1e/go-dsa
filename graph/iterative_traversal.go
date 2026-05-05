package graph

type (
	// Vertex is a vertex in a Graph that has a value and can be connected to more vertices.
	Vertex struct {
		// Val is the value of the vertex
		Val int

		// The edges that this Vertex is connected to
		Edges []*Vertex
	}
	queue     struct{ collection []*Vertex }
	stack     struct{ collection []*Vertex }
	container interface {
		Push(n *Vertex)
		Pop() *Vertex
		Len() int
	}
)

// IterativeTraversal solves the problem in O(n) time and O(n) space.
func IterativeTraversal(graph []*Vertex) ([]int, []int) {
	panic("not implemented")

}

// traverseAllNodes performs BFS or DFS if a graph and a container that is
// either respectively a stack or queue passed to it.
func traverseAllNodes(c container, graph []*Vertex) []int {
	panic("not implemented")

}

func newVertex(v int) *Vertex { panic("not implemented") }

func (q *queue) Len() int       { panic("not implemented") }
func (q *queue) Push(n *Vertex) { panic("not implemented") }
func (q *queue) Pop() *Vertex {
	panic("not implemented")

}

func (s *stack) Len() int       { panic("not implemented") }
func (s *stack) Push(n *Vertex) { panic("not implemented") }
func (s *stack) Pop() *Vertex {
	panic("not implemented")

}

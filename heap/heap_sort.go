package heap

type (
	// Vertex is a node in a heap.
	Vertex struct {
		// Val is the value of the vertex.
		Val int
		// Left is the left child of the vertex.
		Left *Vertex
		// Right is the right child of the vertex.
		Right *Vertex
	}
	// MinHeap is a heap where the root is always the minimum value.
	MinHeap struct {
		Data []*Vertex
	}
)

// Sort solves the problem in O(n*Log n) time and O(n) space.
func Sort(list []int) []int {
	panic("not implemented")

}

// NewMinHeap Returns a new Min Heap.
func NewMinHeap() *MinHeap {
	panic("not implemented")

}

// Push inserts a new value into the heap.
func (m *MinHeap) Push(value int) {
	panic("not implemented")

}

// Pop removes the root value from the heap.
func (m *MinHeap) Pop() int {
	panic("not implemented")

}

// Len returns the number of elements in the heap.
func (m *MinHeap) Len() int {
	panic("not implemented")

}

// percolateUp swaps the vertex up the heap to maintain the heap property.
func (m *MinHeap) percolateUp(index int) {
	panic("not implemented")

}

// percolateDown moves the vertex down the heap to maintain the heap property.
func (m *MinHeap) percolateDown(index int) {
	panic("not implemented")

}

// swap swaps the positions of two vertices in the heap.
func (m *MinHeap) swap(i, j int) {
	panic("not implemented")

}

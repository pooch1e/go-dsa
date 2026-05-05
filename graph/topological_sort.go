package graph

import (
	"errors"
)

// VertexWithIngress is a Vertex with the count of vertices that connect to it.
type VertexWithIngress struct {
	// Val is the value of the vertex
	Val any

	// The edges that this Vertex is connected to
	Edges []*VertexWithIngress

	// Ingress is the number of vertices that connect to this vertex
	Ingress int
}

// ErrNotADAG occurs when a graph is not a Direct Acyclic Graph - DAG.
var ErrNotADAG = errors.New("graph is not a Direct Acyclic Graph - DAG")

// TopologicalSort solves the problem in O(n) time and O(n) space.
func TopologicalSort(graph []*VertexWithIngress) ([]*VertexWithIngress, error) {
	panic("not implemented")

}

package queue

import (
	"github.com/spring1843/go-dsa/tree"
)

type queue struct {
	data []*tree.BinaryTreeNode
}

// IsTreeSymmetrical solves the problem in O(n) time and O(n) space.
func IsTreeSymmetrical(root *tree.BinaryTreeNode) (bool, error) {
	panic("not implemented")

}

func (q *queue) len() int                         { panic("not implemented") }
func (q *queue) enqueue(val *tree.BinaryTreeNode) { panic("not implemented") }
func (q *queue) dequeue() *tree.BinaryTreeNode {
	panic("not implemented")

}

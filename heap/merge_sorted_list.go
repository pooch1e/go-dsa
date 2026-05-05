package heap

import (
	"github.com/spring1843/go-dsa/linkedlist"
)

type (
	priorityQueue []*linkedlist.Node
)

// MergeSortedLists solves the problem in O(n*Log k) time and O(k) space.
func MergeSortedLists(lists []*linkedlist.Node) *linkedlist.Node {
	panic("not implemented")

}

func (pq priorityQueue) Len() int { panic("not implemented") }

func (pq priorityQueue) Less(i, j int) bool {
	panic("not implemented")

}

func (pq *priorityQueue) Pop() any {
	panic("not implemented")

}

func (pq *priorityQueue) Push(x any) {
	panic("not implemented")

}

func (pq priorityQueue) Swap(i, j int) {
	panic("not implemented")

}

package queue

import "math"

const emptyValue = math.MinInt64

type CircularQueue struct {
	data []int
	size int
	head int
	tail int
}

func NewCircularQueue(size int) *CircularQueue {
	panic("not implemented")

}

func (queue *CircularQueue) enqueue(n int) {
	panic("not implemented")

}

func (queue *CircularQueue) dequeue() (int, error) {
	panic("not implemented")

}

func (queue *CircularQueue) nextCircularSlicePosition(value int) int {
	panic("not implemented")

}

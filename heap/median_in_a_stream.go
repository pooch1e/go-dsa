package heap

type (
	minHeap      []int
	maxHeap      []int
	medianKeeper struct {
		max *maxHeap
		min *minHeap
	}
)

func newMedianKeeper() medianKeeper {
	panic("not implemented")

}

// addNumber solves the problem in O(Log n) time and O(n) space.
func (m *medianKeeper) addNumber(num int) {
	panic("not implemented")

}

func (m *medianKeeper) len() int {
	panic("not implemented")

}

func (m *medianKeeper) median() float64 {
	panic("not implemented")

}

func (m minHeap) Len() int           { panic("not implemented") }
func (m minHeap) Less(i, j int) bool { panic("not implemented") }
func (m minHeap) Swap(i, j int)      { panic("not implemented") }
func (m *minHeap) Push(x any)        { panic("not implemented") }

func (m *minHeap) Pop() any {
	panic("not implemented")

}

func (m maxHeap) Len() int           { panic("not implemented") }
func (m maxHeap) Less(i, j int) bool { panic("not implemented") }
func (m maxHeap) Swap(i, j int)      { panic("not implemented") }
func (m *maxHeap) Push(x any)        { panic("not implemented") }

func (m *maxHeap) Pop() any {
	panic("not implemented")

}

package heap

type (
	point struct {
		coordinates []int
		distance    float64
	}
	pointsHeap []*point
)

// KClosestPointToOrigin solves the problem in O(n*Log k) time and O(k) space.
func KClosestPointToOrigin(points [][]int, k int) [][]int {
	panic("not implemented")

}

func newPoint(a []int) *point {
	panic("not implemented")

}

func (p pointsHeap) Len() int           { panic("not implemented") }
func (p pointsHeap) Less(i, j int) bool { panic("not implemented") }
func (p pointsHeap) Swap(i, j int)      { panic("not implemented") }
func (p *pointsHeap) Push(x any)        { panic("not implemented") }

func (p *pointsHeap) Pop() any {
	panic("not implemented")

}

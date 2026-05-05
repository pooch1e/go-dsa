package graph

type vertex struct {
	word     string
	edges    []*vertex
	distance int
}

var graph = make(map[string]*vertex)

// WordLadder solves the problem in O(n) time and O(n) space.
func WordLadder(start, end string, dic []string) int {
	panic("not implemented")

}

func bfsMinTransformation(beginWord string, endWord string) int {
	panic("not implemented")

}

func isDifferentByOneLetter(a, b string) bool {
	panic("not implemented")

}

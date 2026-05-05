package backtracking

var (
	directions      = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	directionLetter = map[int]string{0: "r", 1: "l", 2: "d", 3: "u"}
)

// Maze solves the problem in O(mn) time and O(mn) space complexity.
func Maze(m, n int, walls [][2]int, start, finish [2]int) string {
	panic("not implemented")

}

func mazeRecursive(x, y int, path string, wallMap map[[2]int]bool, visited [][]bool, finish [2]int) string {
	panic("not implemented")

}

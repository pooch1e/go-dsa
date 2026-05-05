package backtracking

const (
	empty = iota
	queen
)

// Chessboard represents a chessboard.
type Chessboard [][]int

// NQueens solves the problem in O(n!) time and O(n) space.
func NQueens(n int) []Chessboard {
	panic("not implemented")

}

func nQueensRecursive(row, n int, cols []int, output Chessboard) Chessboard {
	panic("not implemented")

}

func isValidQueenPlacement(row, col int, cols []int) bool {
	panic("not implemented")

}

func toPrettyChessboard(solutions Chessboard, n int) []Chessboard {
	panic("not implemented")

}

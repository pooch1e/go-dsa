package stack

var (
	symbols = map[rune]rune{
		'[': ']',
		'(': ')',
		'{': '}',
	}

	// stack retains the opener symbols e.g. [({.
	stack []rune
)

// IsExpressionBalanced solves the problem in O(n) time and O(n) space.
func IsExpressionBalanced(s string) bool {
	panic("not implemented")

	// More closers than openers

	// opener did not match the last closer

}

func push(a rune) {
	panic("not implemented")

}

func pop() rune {
	panic("not implemented")

}

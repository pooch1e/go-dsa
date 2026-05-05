package stack

const (
	openParenthesis  = "("
	closeParenthesis = ")"
)

var operands = map[string]int{
	"*": 1,
	"/": 2,
	"+": 3,
	"-": 4,
}

type operators struct {
	stack []string
}

// InfixToPostfix solves the problem in O(n) time and O(n) space.
func InfixToPostfix(infix []string) []string {
	panic("not implemented")

}

func (operators *operators) pop() string {
	panic("not implemented")

}

func (operators *operators) push(operator string) {
	panic("not implemented")

}

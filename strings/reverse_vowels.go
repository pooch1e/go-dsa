package strings

import "errors"

var (
	stack  []string
	vowels = map[string]bool{
		`a`: true,
		`e`: true,
		`o`: true,
		`i`: true,
		`u`: true,
	}
	// ErrPopStack is returned when the stack is empty.
	ErrPopStack = errors.New("can not Pop on an empty stack")
)

// ReverseVowels solves the problem in O(n) time and O(n) space.
func ReverseVowels(str string) (string, error) {
	panic("not implemented")

}

func push(v string) {
	panic("not implemented")

}

func pop() string {
	panic("not implemented")

}

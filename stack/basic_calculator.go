package stack

import (
	"errors"
)

type evaluation struct {
	stack []float64
}

// ErrImbalancedParentheses occurs when the expression has imbalanced parentheses.
var ErrImbalancedParentheses = errors.New("expression has imbalanced parentheses")

// BasicCalculator solves the problem in O(n) time and O(n) space.
func BasicCalculator(input string) (float64, error) {
	panic("not implemented")

}

func toInfix(input string) []string {
	panic("not implemented")

}

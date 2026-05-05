package graph

import (
	"errors"
)

const (
	lineSeparator   = "\n"
	numberSeparator = ","
)

var errInvalidInteger = errors.New("invalid integer")

// HeadCount solves the problem in O(n) time and O(n) space.
func HeadCount(data string, employeeID int) int {
	panic("not implemented")

}

func headCountBFS(graph map[int][]int, employeeID int) int {
	panic("not implemented")

}

// toGraphOfEmployees creates a map of employees and their direct reports.
func toGraphOfEmployees(data string) (map[int][]int, error) {
	panic("not implemented")

}

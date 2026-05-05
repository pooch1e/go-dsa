package graph

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestWordLadder tests solution(s) with the following signature and problem description:

	func WordLadder(start, end string, dic []string) int

Given a start word like `pop` and an end word like `car`, a dictionary of same length words
like  `{"top","cop","cap","car"}` return the minimum number of transformations like 4 to get
from start to end where each transformation between two words can happen when they are
different by only one letter.
*/
var _ = DescribeTable("WordLadder",
	func(start string, end string, dic []string, minTransformations int) {
		if got := WordLadder(start, end, dic); got != minTransformations {
			Expect(got).To(Equal(minTransformations))
		}
	},
	Entry("WordLadder #1", "foo", "bar", []string{}, 0),
	Entry("WordLadder #2", "foo", "bar", []string{"baz"}, 0),
	Entry("WordLadder #3", "a", "c", []string{"a", "b", "c"}, 1),
	Entry("WordLadder #4", "pop", "cop", []string{"tap", "top", "cop"}, 1),
	Entry("WordLadder #5", "car", "tor", []string{"cap", "tap", "top", "tar", "tor"}, 3),
	Entry("WordLadder #6", "pop", "car", []string{"top", "cop", "cap", "car"}, 4),
	Entry("WordLadder #7", "pot", "cop", []string{"rot", "rat", "fat", "fax", "tax", "tap", "top", "cop"}, 8),
)

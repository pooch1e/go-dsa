package backtracking

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"reflect"
)

/*
TestPermutations tests solution(s) with the following signature and problem description:

	func Permutations(input []int) [][]int

Given a list of integers like {1,2}, produce all possible combinations like {1,2},{2,1}.
*/
var _ = DescribeTable("Permutations",
	func(nums []int, permutations [][]int) {
		if got := Permutations(nums); !reflect.DeepEqual(permutations, got) {
			Expect(got).To(Equal(permutations))
		}
	},
	Entry("Permutations #1", []int{}, [][]int{}),
	Entry("Permutations #2", []int{1}, [][]int{{1}}),
	Entry("Permutations #3", []int{1, 2}, [][]int{{1, 2}, {2, 1}}),
	Entry("Permutations #4", []int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}}),
	Entry("Permutations #5", []int{1, 2, 3, 4}, [][]int{{1, 2, 3, 4}, {1, 2, 4, 3}, {1, 3, 2, 4}, {1, 3, 4, 2}, {1, 4, 3, 2}, {1, 4, 2, 3}, {2, 1, 3, 4}, {2, 1, 4, 3}, {2, 3, 1, 4}, {2, 3, 4, 1}, {2, 4, 3, 1}, {2, 4, 1, 3}, {3, 2, 1, 4}, {3, 2, 4, 1}, {3, 1, 2, 4}, {3, 1, 4, 2}, {3, 4, 1, 2}, {3, 4, 2, 1}, {4, 2, 3, 1}, {4, 2, 1, 3}, {4, 3, 2, 1}, {4, 3, 1, 2}, {4, 1, 3, 2}, {4, 1, 2, 3}}),
)

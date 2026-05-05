package array

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestFindDuplicate tests solution(s) with the following signature and problem description:

	FindDuplicate(list []int) int

Given an unsorted slice of n positive integers like {3,2,1,4,5,4,…,n} where each number is smaller
than n and there is at most one duplicate, return the duplicate value like 4.
*/
var _ = DescribeTable("FindDuplicate",
	func(list []int, duplicate int) {
		if got := FindDuplicate(list); got != duplicate {
			Expect(got).To(Equal(duplicate))
		}
	},
	Entry("FindDuplicate #1", []int{}, -1),
	Entry("FindDuplicate #2", []int{1, 2, 2}, 2),
	Entry("FindDuplicate #3", []int{1, 2, 3}, -1),
	Entry("FindDuplicate #4", []int{1, 1, 2, 3}, 1),
	Entry("FindDuplicate #5", []int{1, 2, 2, 3}, 2),
	Entry("FindDuplicate #6", []int{1, 2, 3, 2, 4, 5}, 2),
	Entry("FindDuplicate #7", []int{3, 2, 1, 4, 5, 4}, 4),
)

package dnc

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"slices"
)

/*
TestMergeSort tests solution(s) with the following signature and problem description:

	func MergeSort(list []int) []int

Given a slice of integers, return a sorted slice using Merge Sort.

Merge sort is a divide and conquer algorithm that works by dividing the
input array into two halves, recursively sorting each half, and then merging
the two sorted halves back together. The merge step is where the actual
sorting takes place.

For example given {3,1,2}, return {1,2,3}. The merge sort algorithm would first
divide the array into two halves: {3} and {1,2}. It would then it would divide {1,2}
into {1} and {2}. Finally, it would merge the two halves back together
into a single sorted array: {1,2,3}.
*/
var _ = DescribeTable("MergeSort",
	func(list []int, sorted []int) {
		if got := MergeSort(list); !slices.Equal(got, sorted) {
			Expect(got).To(Equal(sorted))
		}
	},
	Entry("MergeSort #1", []int{}, []int{}),
	Entry("MergeSort #2", []int{1, 2}, []int{1, 2}),
	Entry("MergeSort #3", []int{2, 1}, []int{1, 2}),
	Entry("MergeSort #4", []int{1, 2, 3}, []int{1, 2, 3}),
	Entry("MergeSort #5", []int{3, 2, 1}, []int{1, 2, 3}),
	Entry("MergeSort #6", []int{1, 3, 2}, []int{1, 2, 3}),
	Entry("MergeSort #7", []int{-1, 3, 2, 0, 4}, []int{-1, 0, 2, 3, 4}),
)

package array

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"slices"
)

/*
TestBubbleSort tests solution(s) with the following signature and problem description:

	BubbleSort(input []int)

Sort a slice of unsorted integers using the Bubble Sort algorithm.

The algorithm should be in-place, meaning it should not create a new array and it should
work by swapping elements in the array until it is sorted.

To use Bubble Sort to increasingly sort a slice:

1- Go through each item of the list.
2- If item at n[i] is larger than item at n[i+1], then swap the two elements and repeat the same process for every remaining element(s).
3- If a swap has happened go back to 1, otherwise the list is sorted.

The name Bubble Sort comes from the observation that smaller values bubble up to the top of the list and larger values sink lower until
the list is sorted.

To sort {2,3,1} using Bubble Sort:

1- Compare 2 and 3. Do not swap 2 and 3 because 2 <3.
2- Swap 3 and 1 because 1 < 3, the list becomes {2,1,3}.
3- Swap 1 and 2 because 1 < 2, the list becomes {1,2,3}.
4- Do not swap 2 and 3 because 2 < 3.
5- Do not swap 1 and 2 because 1 < 2.
6- Do not swap 2 and 3 because 2 < 3.
7- The list is sorted, return {1,2,3}.
*/
var _ = DescribeTable("BubbleSort",
	func(input []int, sorted []int) {
		BubbleSort(input)
		if !slices.Equal(input, sorted) {
			Expect(input).To(Equal(sorted))
		}
	},
	Entry("BubbleSort #1", []int{}, []int{}),
	Entry("BubbleSort #2", []int{1}, []int{1}),
	Entry("BubbleSort #3", []int{1, 2}, []int{1, 2}),
	Entry("BubbleSort #4", []int{2, 1}, []int{1, 2}),
	Entry("BubbleSort #5", []int{2, 3, 1}, []int{1, 2, 3}),
	Entry("BubbleSort #6", []int{4, 2, 3, 1, 5}, []int{1, 2, 3, 4, 5}),
)

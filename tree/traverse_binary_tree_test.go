package tree

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"slices"
)

/*
TestTraverseBinaryTree tests solution(s) with the following signature and problem description:

	func TraverseBinaryTree(root *BinaryTreeNode) ([]int, []int, []int)

Given a binary tree like the above return the tree in-order, pre-order, and post-order traversals.

	     4
	    / \
	   /   \
	  2     6
	 / \   / \
	1   3 5   7

For example given the above tree as "4,2,6,1,3,5,7" return:

	in-order traversal as [1,2,3,4,5,6,7],
	pre-order traversal as [4,2,1,3,6,5,7],
	post-order traversal as [1,3,2,5,7,6,4].
*/
var _ = DescribeTable("TraverseBinaryTree",
	func(tree string, in []int, pre []int, post []int) {
		gotIn, gotPre, gotPost := TraverseBinaryTree(Deserialize(tree))
		if !slices.Equal(gotIn, in) {
			Expect(gotIn).To(Equal(in))
		}
		if !slices.Equal(gotPre, pre) {
			Expect(gotPre).To(Equal(pre))
		}
		if !slices.Equal(gotPost, post) {
			Expect(gotPost).To(Equal(post))
		}
	},
	Entry("TraverseBinaryTree #1", "", []int{}, []int{}, []int{}),
	Entry("TraverseBinaryTree #2", "1", []int{1}, []int{1}, []int{1}),
	Entry("TraverseBinaryTree #3", "4,2,6,1,3,5,7", []int{1, 2, 3, 4, 5, 6, 7}, []int{4, 2, 1, 3, 6, 5, 7}, []int{1, 3, 2, 5, 7, 6, 4}),
	Entry("TraverseBinaryTree #4", "3,1,2", []int{1, 3, 2}, []int{3, 1, 2}, []int{1, 2, 3}),
	Entry("TraverseBinaryTree #5", "4,2,6,nil,3,5", []int{2, 3, 4, 5, 6}, []int{4, 2, 3, 6, 5}, []int{3, 2, 5, 6, 4}),
)

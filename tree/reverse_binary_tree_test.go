package tree

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestReverseBinaryTree tests solution(s) with the following signature and problem description:

	func ReverseBinaryTree(root *BinaryTreeNode) *BinaryTreeNode

Given a binary tree, reverse it.

	     1
	   /   \
	  2     3
	 / \   / \
	4   5 6   7

For example given the above tree, return:

	     1
	   /   \
	  3     2
	 / \   / \
	7   6 5   4

Each tree is represented as a string by "1,2,3,4,5,6,7" and "1,3,2,7,6,5,4" respectively.
*/
var _ = DescribeTable("ReverseBinaryTree",
	func(tree string, reversed string) {
		if got := Serialize(ReverseBinaryTree(Deserialize(tree))); got != reversed {
			Expect(got).To(Equal(reversed))
		}
	},
	Entry("ReverseBinaryTree #1", "", ""),
	Entry("ReverseBinaryTree #2", "1,2,3", "1,3,2"),
	Entry("ReverseBinaryTree #3", "1,2,nil", "1,nil,2"),
	Entry("ReverseBinaryTree #4", "1,2,3,4,5,6,7", "1,3,2,7,6,5,4"),
	Entry("ReverseBinaryTree #5", "1,2,3,4", "1,3,2,nil,nil,nil,4"),
)

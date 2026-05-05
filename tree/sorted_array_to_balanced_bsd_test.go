package tree

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestBalancedBinarySearchTree tests solution(s) with the following signature and problem description:

	type BinaryTreeNode struct {
		Val int
		Left *BinaryTreeNode
		Right *BinaryTreeNode
	}
	func BalancedBinarySearchTree(sorted []int) *BinaryTreeNode

Given a sorted slice of integers like return a string representing the integers as a Balanced Binary Tree (BST).

	     3
	    / \
	   /   \
	  2     4
	 /       \
	1         5

For example given return a *BinaryTreeNode that is the root element of the above tree.
*/
var _ = DescribeTable("BalancedBinarySearchTree",
	func(numbers []int, tree string) {
		if got := Serialize(BalancedBinarySearchTree(numbers)); got != tree {
			Expect(got).To(Equal(tree))
		}
	},
	Entry("BalancedBinarySearchTree #1", []int{}, ""),
	Entry("BalancedBinarySearchTree #2", []int{1}, "1"),
	Entry("BalancedBinarySearchTree #3", []int{1, 2}, "2,1"),
	Entry("BalancedBinarySearchTree #4", []int{1, 2, 3}, "2,1,3"),
	Entry("BalancedBinarySearchTree #5", []int{1, 2, 3, 4}, "3,2,4,1"),
	Entry("BalancedBinarySearchTree #6", []int{1, 2, 3, 4, 5}, "3,2,5,1,nil,4"),
	Entry("BalancedBinarySearchTree #7", []int{1, 2, 3, 4, 6}, "3,2,6,1,nil,4"),
	Entry("BalancedBinarySearchTree #8", []int{1, 2, 3, 4, 6, 7}, "4,2,7,1,3,6"),
	Entry("BalancedBinarySearchTree #9", []int{1, 2, 3, 4, 6, 7, 8}, "4,2,7,1,3,6,8"),
)

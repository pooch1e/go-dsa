package tree

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestSerializeAndUnserializeBinaryTree tests solution(s) with the following signature and problem description:

			type BinaryTreeNode struct {
				Val int
				Left *BinaryTreeNode
				Right *BinaryTreeNode
			}

			func Serialize(root *BinaryTreeNode) string
			func Deserialize(s string) *BinaryTreeNode

	     4
	    / \
	   /   \
	  2     6
	 / \   / \
	    3 5

Given a binary tree, write a serialize and deserialize function that turns it into a string representation
and back.

For example given `4,2,6,nil,3,5,nil` representing the above tree, deserialize it to a *BinaryTreeNode,
and given the *BinaryTreeNode serialize it back to `4,2,6,nil,3,5,nil`.
*/
var _ = DescribeTable("SerializeAndUnserializeBinaryTree",
	func(test string) {
		if got := Serialize(Deserialize(test)); got != test {
			Expect(got).To(Equal(test))
		}
	},
	Entry("SerializeAndUnserializeBinaryTree #1", ""),
	Entry("SerializeAndUnserializeBinaryTree #2", "1"),
	Entry("SerializeAndUnserializeBinaryTree #3", "1,2,3"),
	Entry("SerializeAndUnserializeBinaryTree #4", "1,2,nil,4"),
	Entry("SerializeAndUnserializeBinaryTree #5", "1,2,3,4,nil,5,6"),
	Entry("SerializeAndUnserializeBinaryTree #6", "1,2,nil,4,nil,5,6"),
)

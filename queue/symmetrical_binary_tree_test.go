package queue

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/spring1843/go-dsa/tree"
)

/*
TestIsTreeSymmetrical tests solution(s) with the following signature and problem description:

	func IsTreeSymmetrical(root *tree.BinaryTreeNode) (bool, error)

Given a binary tree return true of it is symmetric and false otherwise. A tree is symmetric if you
can draw a vertical line through the root and then the left subtree is the mirror image of the right subtree.

	  Symmetric       Not Symmetric
	      2                2
	    /   \            /   \
	   /     \          /     \
	  4       4        3       4
	 / \     / \      / \     / \
	5   6   6   5    5   6   6   5

For example given "2,4,4,5,6,6,5", shown in the symmetric tree above return true.
Given "2,3,4,5,6,6,5", shown in the not symmetric tree above return false.
*/
var _ = DescribeTable("IsTreeSymmetrical",
	func(treeStr string, isSymmetric bool) {
		got, err := IsTreeSymmetrical(tree.Deserialize(treeStr))
		if err != nil {
			Expect(err).ToNot(HaveOccurred())
		}
		if got != isSymmetric {
			Expect(got).To(Equal(isSymmetric))
		}
	},
	Entry("IsTreeSymmetrical #1", "", false),
	Entry("IsTreeSymmetrical #2", "1", true),
	Entry("IsTreeSymmetrical #3", "1,2,2", true),
	Entry("IsTreeSymmetrical #4", "1,2,3", false),
	Entry("IsTreeSymmetrical #5", "1,2,2,3,nil,nil,3", true),
	Entry("IsTreeSymmetrical #6", "1,2,nil,4", false),
	Entry("IsTreeSymmetrical #7", "1,2,3,4,nil,5,6", false),
	Entry("IsTreeSymmetrical #8", "1,2,nil,4,nil,5,6", false),
	Entry("IsTreeSymmetrical #9", "2,4,4,5,6,5,6", false),
	Entry("IsTreeSymmetrical #10", "2,4,4,5,6,6,5", true),
)

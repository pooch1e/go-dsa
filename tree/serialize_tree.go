package tree

const (
	delimiter = ","
	nilNode   = "nil"
)

// BinaryTreeNode is a node in a Binary Tree of integers.
type BinaryTreeNode struct {
	// Val is an integer value of a node in a Binary Tree
	Val int

	// Left is the left node in a binary tree
	Left *BinaryTreeNode

	// Right is the right node in a binary tree
	Right *BinaryTreeNode
}

// Serialize serializes a given binary tree into a string.
func Serialize(root *BinaryTreeNode) string {
	panic("not implemented")

}

// Deserialize a string into a binary tree in O(n) time and O(n) space.
func Deserialize(s string) *BinaryTreeNode {
	panic("not implemented")

}

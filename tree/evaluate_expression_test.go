package tree

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"strings"
)

/*
TestEvaluateBinaryExpressionTree tests solution(s) with the following signature and problem description:

	func EvaluateBinaryExpressionTree(node *stringBinaryTreeNode) (float64, error)

Given an expression binary tree that contains integers and the four basic arithmetic operations, return
the resulting evaluation of the expression.

		   *
		  / \
		 /   \
		*     +
	   / \   / \
	  5   5 2   2

For example given "*,*,+,5,5,2,2" representing the above tree return 100 because (5 * 5) * (2 + 2) = 100.
*/
var _ = DescribeTable("EvaluateBinaryExpressionTree",
	func(tree string, result float64, expectedErr bool) {
		got, err := EvaluateBinaryExpressionTree(unserializeStringBinaryTree(tree))
		if !expectedErr && err != nil {
			Expect(err).ToNot(HaveOccurred())
		}
		if expectedErr && err == nil {
			Fail("did not get expected error")
		}
		if got != result {
			Expect(got).To(Equal(result))
		}
	},
	Entry("EvaluateBinaryExpressionTree #1", "", 0.0, false),
	Entry("EvaluateBinaryExpressionTree #2", "*,6,2", 12.0, false),
	Entry("EvaluateBinaryExpressionTree #3", "/,6,2", 3.0, false),
	Entry("EvaluateBinaryExpressionTree #4", "-,6,2", 4.0, false),
	Entry("EvaluateBinaryExpressionTree #5", "+,*,6,+,2,nil,nil,3,4", 20.0, false),
	Entry("EvaluateBinaryExpressionTree #6", "*,*,+,5,5,2,2", 100.0, false),
	Entry("EvaluateBinaryExpressionTree #7", "*,A,2", -1, true),
	Entry("EvaluateBinaryExpressionTree #8", "*,2,B", -1, true),
	Entry("EvaluateBinaryExpressionTree #9", "A,1,2", -1, true),
)

var _ = DescribeTable("SerializeAndUnserializeStringBinaryTree",
	func(test string) {
		if got := serializeStringBinaryTree(unserializeStringBinaryTree(test)); got != test {
			Expect(got).To(Equal(test))
		}
	},
	Entry("SerializeAndUnserializeStringBinaryTree #1", ""),
	Entry("SerializeAndUnserializeStringBinaryTree #2", "1"),
	Entry("SerializeAndUnserializeStringBinaryTree #3", "1,2,3"),
	Entry("SerializeAndUnserializeStringBinaryTree #4", "1,2,nil,4"),
	Entry("SerializeAndUnserializeStringBinaryTree #5", "1,2,3,4,nil,5,6"),
	Entry("SerializeAndUnserializeStringBinaryTree #6", "1,2,nil,4,nil,5,6"),
	Entry("SerializeAndUnserializeStringBinaryTree #7", "a,b,nil,c,nil,d,e"),
	Entry("SerializeAndUnserializeStringBinaryTree #8", "+,b,nil,a,nil,c,e"),
)

func serializeStringBinaryTree(root *stringBinaryTreeNode) string {
	if root == nil {
		return ""
	}

	result := ""
	queue := []*stringBinaryTreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = result + nilNode + delimiter
		} else {
			result = result + node.val + delimiter
			queue = append(queue, node.left)
			queue = append(queue, node.right)
		}
	}

	for strings.HasSuffix(result, nilNode+delimiter) {
		result = strings.TrimSuffix(result, nilNode+delimiter)
	}
	for strings.HasSuffix(result, delimiter) {
		result = strings.TrimSuffix(result, delimiter)
	}

	return result
}

func unserializeStringBinaryTree(s string) *stringBinaryTreeNode {
	if s == "" {
		return nil
	}

	nodes := []*stringBinaryTreeNode{}

	for _, node := range strings.Split(s, delimiter) {
		if node == nilNode {
			nodes = append(nodes, nil)
		} else {
			nodes = append(nodes, &stringBinaryTreeNode{val: node})
		}
	}

	i := 1
	for _, node := range nodes {
		if node == nil {
			continue
		}

		if i >= len(nodes) {
			continue
		}
		if nodes[i] != nil {
			node.left = nodes[i]
		}
		i++

		if i >= len(nodes) {
			continue
		}
		if nodes[i] != nil {
			node.right = nodes[i]
		}
		i++
	}

	return nodes[0]
}

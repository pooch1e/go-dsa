package linkedlist

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestReverseLinkedList tests solution(s) with the following signature and problem description:

	type Node struct {
		Val int
		Next *Node
	}

	func ReverseLinkedList(head *Node) *Node

Reverse a linked list in-place.

For example given 1->2->3, return 3->2->1.
*/
var _ = DescribeTable("ReverseLinkedList",
	func(list string, reversed string) {
		got := Serialize(ReverseLinkedList(Deserialize(list)))
		if got != reversed {
			Expect(got).To(Equal(reversed))
		}
	},
	Entry("ReverseLinkedList #1", "", ""),
	Entry("ReverseLinkedList #2", "1", "1"),
	Entry("ReverseLinkedList #3", "1->2", "2->1"),
	Entry("ReverseLinkedList #4", "1->2->3", "3->2->1"),
)

package linkedlist

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"fmt"
)

/*
TestSerializeAndDeserializeLinkedList tests solution(s) with the following signature and problem description:

	type Node struct {
		Val int
		Next *Node
	}

	func Serialize(node *Node) string
	func Deserialize(stringRepresentation string) *Node

Write a function that turns a linked list of integers into a string representation (Serialize) and  a function
that turns that string representation to an actual linked list (Deserialize).

For example consider the following example of a linked list containing numbers 1,2,3:

node1 := &Node{Val: 1}
node2 := &Node{Val: 2}
node3 := &Node{Val: 3}
node1.Next = node2
node2.Next = node3

A string representation of this linked-list should look like 1->2->3. This means that
the node with value 1 is connected to node with value 2, and node with value 2 is connected to node with value 3.
*/
var _ = DescribeTable("SerializeAndDeserializeLinkedList",
	func(test string) {
		got := Serialize(Deserialize(test))
		if got != test {
			Expect(got).To(Equal(test))
		}
	},
	Entry("SerializeAndDeserializeLinkedList #1", ""),
	Entry("SerializeAndDeserializeLinkedList #2", "1"),
	Entry("SerializeAndDeserializeLinkedList #3", "1->2"),
	Entry("SerializeAndDeserializeLinkedList #4", "1->2->3->4->2->1"),
)

var _ = Describe("ATOI", func() {
	It("works correctly", func() {
		if got := atoi("A"); got != -1 {
			Fail(fmt.Sprintf("want %d, got %d", -1, got))
		}
	})
})

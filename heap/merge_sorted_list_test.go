package heap

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/spring1843/go-dsa/linkedlist"
)

/*
TestMergeSortedLists tests solution(s) with the following signature and problem description:

	func MergeSortedLists(lists []*linkedlist.Node) *linkedlist.Node

Given multiple sorted linked lists, join them together into one sorted linked list.

For example given {1->2, 1->3->4, 4->5} return 1->1->2->3->4->4->5.
*/
var _ = DescribeTable("MergeSortedLists",
	func(sortedLinkedLists []string, merged string) {
		nodes := []*linkedlist.Node{}
		for _, sortedLinkedList := range sortedLinkedLists {
			nodes = append(nodes, linkedlist.Deserialize(sortedLinkedList))
		}
		got := linkedlist.Serialize(MergeSortedLists(nodes))
		if got != merged {
			Expect(got).To(Equal(merged))
		}
	},
	Entry("MergeSortedLists #1", []string{""}, ""),
	Entry("MergeSortedLists #2", []string{"1"}, "1"),
	Entry("MergeSortedLists #3", []string{"1", ""}, "1"),
	Entry("MergeSortedLists #4", []string{"1", "1->2"}, "1->1->2"),
	Entry("MergeSortedLists #5", []string{"1", "1->2", "3->4->5"}, "1->1->2->3->4->5"),
	Entry("MergeSortedLists #6", []string{"1", "1->2", "1->2->3", "1->3->5", "1->2->4"}, "1->1->1->1->1->2->2->2->3->3->4->5"),
)

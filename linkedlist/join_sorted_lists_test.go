package linkedlist

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestJoinTwoSortedLinkedLists tests solution(s) with the following signature and problem description:

	func JoinTwoSortedLinkedLists(l1, l2 *Node) *Node

Given two sorted linked lists of integers, merge them into one sorted linked list.

For example given 1->4->6 and 2->3->5->7, return 1->2->3->4->5->6->7.
*/
var _ = DescribeTable("JoinTwoSortedLinkedLists",
	func(list1 string, list2 string, joined string) {
		got := Serialize(JoinTwoSortedLinkedLists(Deserialize(list1), Deserialize(list2)))
		if got != joined {
			Expect(got).To(Equal(joined))
		}
	},
	Entry("JoinTwoSortedLinkedLists #1", "", "", ""),
	Entry("JoinTwoSortedLinkedLists #2", "1", "", "1"),
	Entry("JoinTwoSortedLinkedLists #3", "", "1", "1"),
	Entry("JoinTwoSortedLinkedLists #4", "1", "2", "1->2"),
	Entry("JoinTwoSortedLinkedLists #5", "1", "2->3", "1->2->3"),
	Entry("JoinTwoSortedLinkedLists #6", "1->3", "2", "1->2->3"),
	Entry("JoinTwoSortedLinkedLists #7", "1->4->6", "2->3->5->7", "1->2->3->4->5->6->7"),
	Entry("JoinTwoSortedLinkedLists #8", "1->4->5->6", "2->3->7", "1->2->3->4->5->6->7"),
	Entry("JoinTwoSortedLinkedLists #9", "1->2", "1->2->3", "1->1->2->2->3"),
	Entry("JoinTwoSortedLinkedLists #10", "1->4", "2->3->7", "1->2->3->4->7"),
)

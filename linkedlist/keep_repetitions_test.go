package linkedlist

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestKeepRepetitions tests solution(s) with the following signature and problem description:

	func KeepRepetitions(head *Node) *Node

Given a linked list of sorted integers, create a copy of the list that contains one example of
each repeated item.

For example given 1->1->4->4->6->6->7, return 1->4->6 because
1,4,6 are items that are repeated in this list and 7 is not repeated.
*/
var _ = DescribeTable("KeepRepetitions",
	func(list string, solution string) {
		got := Serialize(KeepRepetitions(Deserialize(list)))
		if got != solution {
			Expect(got).To(Equal(solution))
		}
	},
	Entry("KeepRepetitions #1", "", ""),
	Entry("KeepRepetitions #2", "1", ""),
	Entry("KeepRepetitions #3", "1->1", "1"),
	Entry("KeepRepetitions #4", "1->4->6", ""),
	Entry("KeepRepetitions #5", "1->4->4->4->6", "4"),
	Entry("KeepRepetitions #6", "1->1->4->4->4->6", "1->4"),
	Entry("KeepRepetitions #7", "1->1->4->4->4->6->6", "1->4->6"),
	Entry("KeepRepetitions #8", "1->1->4->4->6->6->7", "1->4->6"),
	Entry("KeepRepetitions #9", "1->1->4->4->4->6->7->7", "1->4->7"),
)

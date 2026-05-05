package queue

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestQueueUsingStacks tests solution(s) with the following signature and problem description:

	type UsingStacks struct {
		stack1 Stack
		stack2 Stack
	}
	(usingStacks *UsingStacks) enqueue(n int)
	(usingStacks *UsingStacks) dequeue() int

Implement a queue of integers using two stacks. The queue should support enqueue and dequeue operations.
*/
var _ = DescribeTable("QueueUsingStacks",
	func(enqueue []int, dequeueTimes int, expectedLastDequeuedItem int, expectErr bool) {
		queue := new(UsingStacks)
		for _, n := range enqueue {
			queue.enqueue(n)
		}
		var n int
		var err error
		for j := 1; j <= dequeueTimes; j++ {
			n = queue.dequeue()
			if !expectErr && err != nil {
				Expect(err).ToNot(HaveOccurred())
			}
		}
		if n != expectedLastDequeuedItem {
			Expect(n).To(Equal(expectedLastDequeuedItem))
		}
	},
	Entry("QueueUsingStacks #1", []int{1, 2, 3, 4}, 1, 1, false),
	Entry("QueueUsingStacks #2", []int{1, 2, 3, 4}, 2, 2, false),
	Entry("QueueUsingStacks #3", []int{1, 2, 3, 4}, 3, 3, false),
	Entry("QueueUsingStacks #4", []int{1, 2, 2, 3}, 2, 2, false),
)

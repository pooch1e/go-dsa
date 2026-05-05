package linkedlist

import (
	. "github.com/onsi/ginkgo/v2"

	"fmt"
)

/*
TestLRU tests solution(s) with the following signature and problem description:

	func get(key int) int
	func put(key int, value int)

Implement a least recently used cache with integer keys and values, where the least
recently used key is evicted upon insertion when the cache is at full capacity.

For example if we have a cache with capacity 2, and we perform:

	put(0, 1) // put key 0 with value 1
	put(1, 2)
	get(1) // get value with key 1
	put(2, 3)

When putting 2,3, the cache is at capacity and it needs to evict. Since we used key 1 by
getting it, the [1,2] key value pair which is the least used value will be evicted.
What remains in the cache
is [1:2] and [2:3].
*/
var _ = DescribeTable("LRU",
	func(capacity int, puts []int, gets []int) {
		lru := newLRU(capacity)
		for j := 0; j < len(puts); j += 2 {
			lru.put(puts[j], puts[j+1])
		}
		for j := 0; j < len(gets); j += 2 {
			got := lru.get(gets[j])
			want := gets[j+1]
			if got != want {
				Fail(fmt.Sprintf("Failed test  on get #%d, want %d got %d", j, want, got))
			}
		}
	},
	Entry("LRU #1", 1, []int{}, []int{1, -1}),
	Entry("LRU #2", 1, []int{2, 1, 1, 1, 2, 3, 4, 1}, []int{2, -1, 4, 1}),
	Entry("LRU #3", 2, []int{2, 1, 1, 1, 2, 3, 4, 1}, []int{1, -1, 2, 3, 4, 1}),
	Entry("LRU #4", 3, []int{2, 1, 1, 1, 2, 3, 4, 1}, []int{1, 1, 2, 3, 4, 1}),
)

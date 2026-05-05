package linkedlist

import (
	"container/list"
)

type (
	lruCache struct {
		list     *list.List
		elements map[int]*list.Element
		cap      int
	}
	element struct {
		key   int
		value int
	}
)

func newLRU(capacity int) *lruCache {
	panic("not implemented")

}

// get solves the problem in O(1) time and O(n) space.
func (cache *lruCache) get(key int) int {
	panic("not implemented")

}

// put solves the problem in O(1) time and O(n) space.
func (cache *lruCache) put(key int, value int) {
	panic("not implemented")

}

// String represents the cache as a string.
func (cache *lruCache) String() string {
	panic("not implemented")

}

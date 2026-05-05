package dnc

var rateLimitEvents []int64

// nanoSecond represents one second in nanoseconds.
const nanoSecond = 1e9

// IsAllowed solves the problem in O(log n) time and O(n) space.
func IsAllowed(limitPerSecond int) bool {
	panic("not implemented")

}

// removeOldRateLimitEvents uses binary search to remove events older than 1 second.
func removeOldRateLimitEvents(now int64) {
	panic("not implemented")

}

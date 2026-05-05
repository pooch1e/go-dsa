package dnc

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"slices"
	"time"
)

/*
TestRateLimiter tests solution(s) with the following signature and problem description:

	func IsAllowed(limitPerSecond int) bool

Given a number of allowed requests calls per second (calls/time) write an IsAllowed
function which returns false if the request should be rate limited because it exceeds the
limit and true if the request should be allowed.

For example given limit of 2 requests per second, if 3 requests are made in the first second
two calls to IsAllowed would return true and the third would return false.
*/
var _ = DescribeTable("RateLimiter",
	func(limitPerSecond int, firstCallTimes int, sleep int, secondCallTimes int, want []bool) {
		rateLimitEvents = make([]int64, 0)
		got := make([]bool, 0)
		for range firstCallTimes {
			got = append(got, IsAllowed(limitPerSecond))
		}
		time.Sleep(time.Duration(sleep) * time.Second)
		for range secondCallTimes {
			got = append(got, IsAllowed(limitPerSecond))
		}
		if !slices.Equal(got, want) {
			Expect(got).To(Equal(want))
		}
	},
	Entry("RateLimiter #1", 0, 0, 0, 0, []bool{}),
	Entry("RateLimiter #2", 0, 1, 0, 0, []bool{false}),
	Entry("RateLimiter #3", 0, 1, 0, 1, []bool{false, false}),
	Entry("RateLimiter #4", 1, 2, 0, 2, []bool{true, false, false, false}),
	Entry("RateLimiter #5", 2, 5, 0, 5, []bool{true, true, false, false, false, false, false, false, false, false}),
	Entry("RateLimiter #6", 3, 5, 1, 5, []bool{true, true, true, false, false, true, true, true, false, false}),
	Entry("RateLimiter #7", 10, 100, 1, 100, []bool{true, true, true, true, true, true, true, true, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}),
)

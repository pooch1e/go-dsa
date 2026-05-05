package greedy

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestMaxStockProfit tests solution(s) with the following signature and problem description:

	func MaxStockProfit(prices []int) int

Given a list of stock prices for a given stock over time like {50, 10, 4, 100, 1, 101, 5, 10}
return the maximum profit that can be made by buying and selling a single unit of this stock.
Like 101 - 1 = 100.
*/
var _ = DescribeTable("MaxStockProfit",
	func(prices []int, maxProfit int) {
		if got := MaxStockProfit(prices); got != maxProfit {
			Expect(got).To(Equal(maxProfit))
		}
	},
	Entry("MaxStockProfit #1", []int{}, 0),
	Entry("MaxStockProfit #2", []int{1}, 0),
	Entry("MaxStockProfit #3", []int{1, 3}, 2),
	Entry("MaxStockProfit #4", []int{1, 3, 1}, 2),
	Entry("MaxStockProfit #5", []int{1, 3, 1, 2, 3, 8, 2, 12}, 11),
	Entry("MaxStockProfit #6", []int{50, 10, 4, 100, 1, 101, 5, 10}, 100),
)

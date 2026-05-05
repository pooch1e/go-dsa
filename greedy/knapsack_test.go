package greedy

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestKnapsack tests solution(s) with the following signature and problem description:

	func Knapsack(items []KnapsackItem, capacity int) int

Given weight capacity of a knapsack like 5, a list of divisible items (such as pieces
of metal) with values and weights like `{Value: 6, Weight: 2}, {Value: 10, Weight: 2},
{Value: 12, Weight: 3}` return the maximum value of items that can be placed in the
knapsack like `22`.
*/
var _ = DescribeTable("Knapsack",
	func(items []KnapsackItem, capacity int, maxProfit int) {
		if got := Knapsack(items, capacity); got != maxProfit {
			Expect(got).To(Equal(maxProfit))
		}
	},
	Entry("Knapsack #1", []KnapsackItem{}, 0, 0),
	Entry("Knapsack #2", []KnapsackItem{{Value: 1, Weight: 1}}, 1, 1),
	Entry("Knapsack #3", []KnapsackItem{{Value: 1, Weight: 1}, {Value: 10, Weight: 5}}, 5, 10),
	Entry("Knapsack #4", []KnapsackItem{{Value: 1, Weight: 1}, {Value: 10, Weight: 5}}, 2, 4),
	Entry("Knapsack #5", []KnapsackItem{{Value: 6, Weight: 2}, {Value: 10, Weight: 2}, {Value: 12, Weight: 3}}, 5, 22),
	Entry("Knapsack #6", []KnapsackItem{{Value: 7, Weight: 3}, {Value: 5, Weight: 2}, {Value: 2, Weight: 1}, {Value: 10, Weight: 4}}, 6, 14),
	Entry("Knapsack #7", []KnapsackItem{{Value: 50, Weight: 10}, {Value: 55, Weight: 20}, {Value: 110, Weight: 30}}, 50, 187),
)

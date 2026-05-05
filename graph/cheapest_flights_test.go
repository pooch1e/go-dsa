package graph

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"math"
)

/*
TestCheapestFlights tests solution(s) with the following signature and problem description:

	func CheapestFlights(flights [][]int, cityCount, source, destination, maxStops int) int

Given a list of flights each with a source and destination and a price, a maximum number of stops,
source, and destination cities, return the cheapest costs not exceeding the maximum number of stops
to reach from source city to destination.
*/
var _ = DescribeTable("CheapestFlights",
	func(flight [][]int, cityCount int, source int, destination int, maxStops int, cheapestCost int) {
		got := CheapestFlights(flight, cityCount, source, destination, maxStops)
		if got != cheapestCost {
			Expect(got).To(Equal(cheapestCost))
		}
	},
	Entry("CheapestFlights #1", [][]int{{0, 1, math.MaxInt64}}, 3, 0, 2, 2, -1),
	Entry("CheapestFlights #2", [][]int{{0, 1, 10}, {1, 2, 10}, {0, 2, 5}}, 3, 0, 2, 2, 5),
	Entry("CheapestFlights #3", [][]int{{0, 1, 10}, {1, 2, 10}, {0, 2, 5}}, 3, 0, 2, 2, 5),
	Entry("CheapestFlights #4", [][]int{{0, 1, 10}, {1, 2, 10}, {0, 2, 10}}, 3, 0, 2, 2, 10),
	Entry("CheapestFlights #5", [][]int{{0, 1, 10}, {1, 2, 10}, {0, 2, 20}}, 3, 0, 2, 1, 20),
	Entry("CheapestFlights #6", [][]int{{0, 1, 10}, {1, 2, 10}, {0, 2, 20}}, 3, 0, 2, 1, 20),
)

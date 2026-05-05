package graph

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
TestEmployeeHeadCount tests solution(s) with the following signature and problem description:

	HeadCount(employeeID int) int

Headcount of a person in an organization is 1 + the number of all their reports
(direct and indirect). Given a list of employee IDs and their direct reports in each line like
`1,4,5`, `5,7`, `4`, `7`. Where 1 has two direct reports (4 and 5) and one indirect report (7); 5 has one report (7);
4 and 7 have no reports; Return the headcount of a given employee ID. For example head
count of 7 is 1, and headcount of 1 is 4.
*/
var _ = DescribeTable("EmployeeHeadCount",
	func(employeeData string, headCounts map[int]int) {
		for employeeID, headCount := range headCounts {
			got := HeadCount(employeeData, employeeID)
			if got != headCount {
				Expect(headCount, got).To(Equal(employeeID))
			}
		}
	},
	Entry("EmployeeHeadCount #1", "A", map[int]int{1: -1}),
	Entry("EmployeeHeadCount #2", "1", map[int]int{1: 1}),
	Entry("EmployeeHeadCount #3", "A,B\nC,D,E,F\nG\nH\nI", map[int]int{1: -1}),
	Entry("EmployeeHeadCount #4", "1,2\n3,4,5,6\n4\n5\n6", map[int]int{1: 2}),
	Entry("EmployeeHeadCount #5", "1,2\n3,4,5,6\n4\n5\n6", map[int]int{3: 4}),
	Entry("EmployeeHeadCount #6", "1,2\n3,4,5,6\n4\n5\n6,7\n7", map[int]int{3: 5}),
	Entry("EmployeeHeadCount #7", "1,4,5\n5,7\n4\n7", map[int]int{7: 1, 4: 1, 5: 2, 1: 4}),
	Entry("EmployeeHeadCount #8", "23,29,27\n25\n26,25\n27\n28,32\n29,30\n30\n31\n32", map[int]int{23: 4, 29: 2, 32: 1, 30: 1}),
	Entry("EmployeeHeadCount #9", "23,29,27,28\n25\n26,25\n27\n28,32\n29,30\n30\n31\n32", map[int]int{23: 6, 29: 2, 32: 1, 30: 1}),
)

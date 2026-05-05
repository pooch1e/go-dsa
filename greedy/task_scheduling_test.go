package greedy

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"slices"
)

/*
TestScheduleEvents tests solution(s) with the following signature and problem description:

	func ScheduleEvents(events []Event) []Event

Given a list of named tasks with their start and end timing like `{A, 1, 3}, {B, 2, 3}, {C, 3, 4}`
(Task A starts at 1 and ends at 3). return a schedule that includes as many events as
possible like `{A 1 3}, {C 3 4}`.
*/
var _ = DescribeTable("ScheduleEvents",
	func(events []Event, schedule []Event) {
		if got := ScheduleEvents(events); !slices.Equal(got, schedule) {
			Expect(got).To(Equal(schedule))
		}
	},
	Entry("ScheduleEvents #1", []Event{}, nil),
	Entry("ScheduleEvents #2", []Event{{Name: "A", StartTime: 1, EndTime: 2}}, []Event{{Name: "A", StartTime: 1, EndTime: 2}}),
	Entry("ScheduleEvents #3", []Event{{Name: "A", StartTime: 1, EndTime: 3}, {Name: "B", StartTime: 2, EndTime: 3}, {Name: "C", StartTime: 3, EndTime: 4}}, []Event{{Name: "A", StartTime: 1, EndTime: 3}, {Name: "C", StartTime: 3, EndTime: 4}}),
	Entry("ScheduleEvents #4", []Event{{Name: "A", StartTime: 1, EndTime: 2}, {Name: "B", StartTime: 1, EndTime: 2}}, []Event{{Name: "A", StartTime: 1, EndTime: 2}}),
	Entry("ScheduleEvents #5", []Event{{Name: "A", StartTime: 1, EndTime: 2}, {Name: "B", StartTime: 3, EndTime: 4}}, []Event{{Name: "A", StartTime: 1, EndTime: 2}, {Name: "B", StartTime: 3, EndTime: 4}}),
	Entry("ScheduleEvents #6", []Event{{Name: "A", StartTime: 1, EndTime: 2}, {Name: "B", StartTime: 1, EndTime: 3}, {Name: "C", StartTime: 5, EndTime: 6}}, []Event{{Name: "A", StartTime: 1, EndTime: 2}, {Name: "C", StartTime: 5, EndTime: 6}}),
	Entry("ScheduleEvents #7", []Event{{Name: "A", StartTime: 1, EndTime: 3}, {Name: "B", StartTime: 1, EndTime: 3}, {Name: "C", StartTime: 3, EndTime: 9}, {Name: "D", StartTime: 5, EndTime: 10}}, []Event{{Name: "A", StartTime: 1, EndTime: 3}, {Name: "C", StartTime: 3, EndTime: 9}}),
)

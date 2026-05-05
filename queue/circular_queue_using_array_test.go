package queue

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type testRound struct {
	enqueueStart int
	enqueueEnd   int
	dequeueStart int
	dequeueEnd   int
	expectedErr  error
}

/*
TestCircularQueue tests solution(s) with the following signature and problem description:

	func (queue *CircularQueue) enqueue(n int) error
	func (queue *CircularQueue) dequeue() (int, error)

Given a size, implement a circular queue using a slice.

A circular queue also called a ring buffer is different from a normal queue in that
the last element is connected to the first element.

For example given a circular array of size 4 if we enqueue integers {1,2,3,4,5,6} and then dequeue 2 times
we should get 3 and 4.

	        , - ~ - ,                 , - ~ - ,                 , - ~ - ,                 , - ~ - ,
	    , '    ...    ' ,         , '    ...    ' ,         , '    ...    ' ,         , '    ...    ' ,
	  ,        ...        ,     ,        ...        ,     ,        ...        ,     ,        ...        ,
	 ,    1    ...         ,   ,    1    ...   2     ,   ,    1    ...   2     ,   ,    1    ...   2     ,
	,          ...          , ,          ...          , ,          ...          , ,          ...          ,
	,  ...................  , ,  ...................  , ,  ...................  , ,  ...................  ,
	,  ...................  , ,  ...................  , ,  ...................  , ,  ...................  ,
	 ,         ...         ,   ,         ...         ,   ,         ...         ,   ,         ...         ,
	  ,        ...        ,     ,        ...        ,     ,        ...   3    ,     ,   4    ...   3    ,
	    ,      ...     , '        ,      ...     , '        ,      ...     , '        ,      ...     , '
	      ' - ,...,  '              ' - ,...,  '              ' - ,...,  '              ' - ,...,  '

	        , - ~ - ,                 , - ~ - ,                 , - ~ - ,                 , - ~ - ,
	    , '    ...    ' ,         , '    ...    ' ,         , '    ...    ' ,         , '    ...    ' ,
	  ,        ...        ,     ,        ...        ,     ,        ...        ,     ,        ...        ,
	 ,     5   ...   2     ,   ,     5   ...   6     ,   ,     5   ...   6     ,   ,    5    ...    6    ,
	,          ...          , ,          ...          , ,          ...          , ,          ...          ,
	,  ...................  , ,  ...................  , ,  ...................  , ,  ...................  ,
	,  ...................  , ,  ...................  , ,  ...................  , ,  ...................  ,
	 ,         ...         ,   ,         ...         ,   ,         ...         ,   ,         ...         ,
	  ,    4   ...   3    ,     ,    4   ...   3    ,     ,    4   ...        ,     ,        ...        ,
	    ,      ...     , '        ,      ...     , '        ,      ...     , '        ,      ...     , '
	      ' - ,...,  '              ' - ,...,  '              ' - ,...,  '              ' - ,...,  '

As shown in the above example the circular queue does not run out of capacity and still allows FIFO operations.
*/
var _ = DescribeTable("CircularQueue",
	func(size int, testRounds []testRound) {
		queue := NewCircularQueue(size)
		for j, testRound := range testRounds {
			for i := testRound.enqueueStart; i <= testRound.enqueueEnd; i++ {
				queue.enqueue(i)
			}
			for want := testRound.dequeueStart; want <= testRound.dequeueEnd; want++ {
				got, err := queue.dequeue()
				if err != nil {
					if err != testRound.expectedErr {
						Expect(j, err).ToNot(HaveOccurred())
					}
					break
				}
				if got != want {
					Expect(want, got).To(Equal(j))
				}
			}
		}
	},
	Entry("CircularQueue #1", 1, []testRound{{1, 6, 6, 6, nil}}),
	Entry("CircularQueue #2", 2, []testRound{{1, 6, 5, 5, nil}}),
	Entry("CircularQueue #3", 4, []testRound{{1, 6, 3, 6, nil}}),
	Entry("CircularQueue #4", 4, []testRound{{1, 6, 3, 6, nil}, {1, 6, 3, 6, nil}}),
	Entry("CircularQueue #5", 4, []testRound{{1, 6, 3, 7, ErrQueueEmpty}}),
)

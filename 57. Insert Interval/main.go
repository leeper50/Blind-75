// Result - Passed
// Runtime - Beats 100%
// Memory - Beats 97.47%

package main

import (
	"fmt"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	var mergedInterval [][]int
	i := 0

	// While below newInterval's start, append the current interval
	for ; i < len(intervals) && intervals[i][1] < newInterval[0]; i++ {
		mergedInterval = append(mergedInterval, intervals[i])
	}

	// If newInterval is within an interval, adjust newInterval until it is not
	for ; i < len(intervals) && intervals[i][0] <= newInterval[1]; i++ {
		newInterval[0] = min(intervals[i][0], newInterval[0])
		newInterval[1] = max(intervals[i][1], newInterval[1])
	}
	mergedInterval = append(mergedInterval, newInterval)

	// Append all remaining intervals
	for i < len(intervals) {
		mergedInterval = append(mergedInterval, intervals[i])
		i++
	}

	return mergedInterval
}

func main() {
	var data [][]int
	data = [][]int{{1, 3}, {6, 9}}
	fmt.Printf("%v == [[1 5] [6 9]]\n", insert(data, []int{2, 5}))
	data = [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	fmt.Printf("%v == [[1 2] [3 10] [12 16]]\n", insert(data, []int{4, 8}))
	data = [][]int{{1, 1}, {6, 9}}
	fmt.Printf("%v == [[1 1] [2 5] [6 9]\n", insert(data, []int{2, 5}))
	data = [][]int{{3, 5}, {6, 9}}
	fmt.Printf("%v == [[1 2] [3 5] [6 9]\n", insert(data, []int{1, 2}))
}

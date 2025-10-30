// Result - Passed
// Runtime - Beats 100%
// Memory - Beats 97.47%

package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	// Sort intervals based on first item
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Store first interval
	result := [][]int{intervals[0]}

	// Loop through intervals
	for i := 1; i < len(intervals); i++ {
		lastMerged := result[len(result)-1]
		current := intervals[i]
		// Check start of current against end of last merged
		if current[0] <= lastMerged[1] {
			// Check end of current against end of last merged
			if current[1] > lastMerged[1] {
				lastMerged[1] = current[1]
			}
		} else {
			// If current start is beyond last merged end, add to merged
			result = append(result, current)
		}
	}
	return result
}

func main() {
	var data [][]int
	data = [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Printf("%v\n", merge(data))
	data = [][]int{{1, 4}, {4, 5}}
	fmt.Printf("%v\n", merge(data))
	data = [][]int{{4, 7}, {1, 4}}
	fmt.Printf("%v\n", merge(data))
	data = [][]int{{1, 3}, {2, 4}, {2, 6}, {8, 9}, {8, 10}, {9, 11}, {15, 18}, {16, 17}}
	fmt.Printf("%v\n", merge(data))
	data = [][]int{{1, 4}, {2, 3}}
	fmt.Printf("%v\n", merge(data))
}

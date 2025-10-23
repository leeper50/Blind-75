// Result - Passed
// Runtime - Beats 81.11%
// Memory - Beats 78.13%

package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	result := [][]int{}
	// Sort array
	sort.Ints(nums)
	for i := range nums {
		// Loop until nums[i] is a new value
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// Make 2 pointers at start and end of nums
		j, k := i+1, len(nums)-1
		// Loop until they equal
		for j < k {
			total := nums[i] + nums[j] + nums[k]
			// If total is too small, move j forward
			// If total is too big, move k backward
			if total < 0 {
				j++
			} else if total > 0 {
				k--
			} else {
				// Total sums to 0, add result
				// move j forward and keep doing it until nums[j] is a new value
				// move k backward and keep doing it until nums[k] is a new value
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				for nums[j] == nums[j-1] && j < k {
					j++
				}
				for nums[k-1] == nums[k] && j < k {
					k--
				}
			}
		}
	}
	return result
}

func main() {
	result := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Printf("%v\t%s\n", result, "[[-1 -1 2] [-1 0 1]]")
	result = threeSum([]int{-1, -2, 3})
	fmt.Printf("%v\t%s\n", result, "[[-2 1 3]]")
	result = threeSum([]int{0, 0, 0})
	fmt.Printf("%v\t%s\n", result, "[[0 0 0]]")
	result = threeSum([]int{0, 0, 0, 0})
	fmt.Printf("%v\t%s\n", result, "[[0 0 0]]")
}

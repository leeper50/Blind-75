// Result - Passed
// Runtime - Beats 100%
// Memory - Beats 52.98%

package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var backtrack func([]int, int, []int)

	backtrack = func(candidates []int, target int, selected []int) {
		// If current total equals original target, append to result and return
		if target == 0 {
			result = append(result, append([]int{}, selected...))
			return
		}
		// Loop through candidates until value is greater than the target
		// Backtrack:
		// Pass candidates starting at index i to check the same value (allowed duplicates)
		// Setting new target as original target minus current value
		// Keep track of summed values by appending current value to selected
		for i, val := range candidates {
			if val <= target {
				backtrack(candidates[i:], target-val, append(selected, val))
			}
		}
	}

	backtrack(candidates, target, []int{})
	return result
}

func main() {
	fmt.Printf("%v\n", combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Printf("%v\n", combinationSum([]int{2, 3, 5, 8}, 8))
}

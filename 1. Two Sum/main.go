// Result - Passed
// Runtime - Beats 100%
// Memory - Beats 62.25%

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	tempSet := make(map[int]int, len(nums)/2)
	for i := range nums {
		complement := target - nums[i]
		if idx, present := tempSet[complement]; present {
			return []int{idx, i}
		}
		tempSet[nums[i]] = i
	}
	return []int{0, 0}
}

func main() {
	fmt.Printf("%v [0 1]\n", twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Printf("%v [2 3]\n", twoSum([]int{2, 7, 11, 15}, 26))
	fmt.Printf("%v [1 2]\n", twoSum([]int{3, 2, 4}, 6))
	fmt.Printf("%v [0 1]\n", twoSum([]int{3, 3}, 6))
	fmt.Printf("%v [0 1]\n", twoSum([]int{3, 3}, 6))
	fmt.Printf("%v [0 1]\n", twoSum([]int{3, -3}, 0))
	fmt.Printf("%v [0 1]\n", twoSum([]int{-3, 3}, 0))
}

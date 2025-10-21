// Result - Passed
// Runtime - Beats 66.87%
// Memory - Beats 97.38%

package main

import (
	"sort"
)

func containsDuplicateMap(nums []int) bool {
	mySet := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := mySet[v]; ok {
			return true
		}
		mySet[v] = struct{}{}
	}
	return false
}

func containsDuplicateSort(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

func main() {
	println(containsDuplicateMap([]int{1, 2, 3, 1}) == true)
	println(containsDuplicateMap([]int{1, 2, 3, 4}) == false)
	println(containsDuplicateMap([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}) == true)
	println(containsDuplicateMap([]int{1}) == false)

	println(containsDuplicateSort([]int{1, 2, 3, 1}) == true)
	println(containsDuplicateSort([]int{1, 2, 3, 4}) == false)
	println(containsDuplicateSort([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}) == true)
	println(containsDuplicateSort([]int{1}) == false)
}

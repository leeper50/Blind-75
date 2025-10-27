// Result - Passed
// Runtime - Beats 100%
// Memory - Beats 79.95%

package main

func search(nums []int, target int) int {
	// Run binary search
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		// Return index if equal
		if nums[mid] == target {
			return mid
		}
		// Check that indices left to mid are sorted
		if nums[left] <= nums[mid] {
			// Run binary search on left half
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
			// If indices left to mid are not sorted
		} else {
			// Run binary search on right half
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	// Return index not found value
	return -1
}

func main() {
	println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0), 4)
	println(search([]int{4, 5, 6, 7, 8, 9, 0, 1, 2}, 3), -1)
	println(search([]int{1}, 0), -1)
	println(search([]int{1, 3}, 0), -1)
	println(search([]int{1, 3}, 1), 0)
	println(search([]int{1, 3}, 3), 1)
	println(search([]int{5, 1, 3}, 3), 2)
}

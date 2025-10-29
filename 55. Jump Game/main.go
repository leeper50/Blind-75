// Result - Passed
// Runtime - Beats 100%
// Memory - Beats 99.91%

package main

func canJump(nums []int) bool {
	jumps := nums[0]
	for i := 1; i < len(nums); i++ {
		if jumps == 0 {
			return false
		}
		if nums[i] > jumps-1 {
			jumps = nums[i]
		} else {
			jumps--
		}
	}
	return true
}

func main() {
	println(canJump([]int{0}) == true)
	println(canJump([]int{2, 3, 1, 1, 4}) == true)
	println(canJump([]int{2, 5, 0, 0}) == true)
	println(canJump([]int{3, 2, 1, 0, 4}) == false)
	println(canJump([]int{1, 2, 4, 2, 1, 0, 4, 0, 1}) == true)
	println(canJump([]int{1, 2, 0, 0, 2}) == false)
	println(canJump([]int{3, 0, 0, 0, 2, 0}) == false)
	println(canJump([]int{2, 0, 0}) == true)
}

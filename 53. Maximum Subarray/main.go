// Result - Passed
// Runtime - Beats 100%
// Memory - Beats 62.62%

package main

func maxSubArray(nums []int) int {
	curSum, bestSum := 0, nums[0]
	for _, num := range nums {
		curSum += num
		bestSum = max(bestSum, curSum)
		if curSum < 0 {
			curSum = 0
		}
	}
	return bestSum
}

func main() {
	println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}), "6")
	println(maxSubArray([]int{1}), "1")
	println(maxSubArray([]int{5, 4, -1, 7, 8}), "23")
	println(maxSubArray([]int{-5, -4, -3}), "-3")
}

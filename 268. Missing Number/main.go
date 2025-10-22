// Result - Passed
// Runtime - Beats 100%
// Memory - Beats 73.89%

package main

func missingNumber(nums []int) int {
	n := len(nums)
	total := n * (n + 1) / 2
	for _, v := range nums {
		total = total - v
	}
	return total
}

func main() {
	println("Result:", missingNumber([]int{3, 0, 1}), 2)
	println("Result:", missingNumber([]int{0, 1}), 2)
	println("Result:", missingNumber([]int{2, 1}), 0)
	println("Result:", missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}), 8)
	println("Result:", missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 8, 1}), 0)
}

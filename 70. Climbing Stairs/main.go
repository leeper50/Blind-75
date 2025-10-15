// Result - Passed
// Runtime - Beats 100%
// Memory - Beats 99.82%

package main

func climbStairs(n int) int {
	result, next, secondNext := 0, 0, 0
	for i := 1; i <= n; i++ {
		if i == 1 {
			result = 1
		} else if i == 2 {
			result = 2
		} else {
			result = secondNext + next
		}
		next = secondNext
		secondNext = result
	}
	return result
}

func main() {
	println(climbStairs(1), "1")
	println(climbStairs(2), "2")
	println(climbStairs(3), "3")
	println(climbStairs(4), "5")
	println(climbStairs(5), "8")
}

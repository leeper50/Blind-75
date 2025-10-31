// Result - Passed
// Runtime - Beats 100%
// Memory - Beats 96.57%

package main

func uniquePaths(m int, n int) int {
	// Combinatorial mathematics
	// m-1 choices for down
	m--
	// n-1 choices for right
	n--
	// Each unique path is a unique string of Down and Right
	// Use binomial to calculate all possible choices
	var binomial func(int, int) int
	binomial = func(n, k int) int {
		if k == 0 {
			return 1
		}
		return (n - k + 1) * binomial(n, k-1) / k
	}
	return binomial(m+n, min(m, n))
}

func main() {
	println(uniquePaths(3, 7), 28)
	println(uniquePaths(2, 2), 2)
	println(uniquePaths(3, 2), 3)
	println(uniquePaths(3, 3), 6)
	println(uniquePaths(4, 4), 20)
	println(uniquePaths(80, 2), 80)
}

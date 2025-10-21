// Result - Passed
// Runtime - Beats 100%
// Memory - Beats 18.05%

package main

func hammingWeight(n int) int {
	result := 0
	for range 32 {
		result += n & 1
		n /= 2
	}
	return result
}

func main() {
	println(hammingWeight(11) == 3)
	println(hammingWeight(128) == 1)
	println(hammingWeight(2147483645) == 30)
}

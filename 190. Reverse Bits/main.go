// Result - Passed
// Runtime - Beats 100%
// Memory - Beats 33.97%

package main

func reverseBits(n int) int {
	var result int = 0
	for range 32 {
		result *= 2
		result += n & 1
		n /= 2
	}
	return result
}

func main() {
	println(reverseBits(43261596) == 964176192)
	println(reverseBits(0) == 0)
	println(reverseBits(2147483644) == 1073741822)
}

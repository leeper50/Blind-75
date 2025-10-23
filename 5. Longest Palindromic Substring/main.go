// Result - Passed
// Runtime - Beats 73.93%
// Memory - Beats 28.47%

package main

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	start, end := 0, 0

	expand := func(left, right int) (int, int) {
		for left >= 0 && right < len(runes) && runes[left] == runes[right] {
			left--
			right++
		}
		return left + 1, right - 1
	}
	for i := range runes {
		l1, r1 := expand(i, i)
		l2, r2 := expand(i, i+1)

		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}
	return string(runes[start : end+1])
}

func main() {
	println(longestPalindrome("babad"), "bab")
	println(longestPalindrome("cbbd"), "bb")
	println(longestPalindrome("a"), "a")
	println(longestPalindrome("ac"), "a")
}

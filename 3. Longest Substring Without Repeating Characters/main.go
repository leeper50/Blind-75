// Result - Passed
// Runtime - Beats 84.39%
// Memory - Beats 48.63%

package main

func lengthOfLongestSubstring(s string) int {
	// use pointer to keep track of start of subtring
	result, start := 0, 0
	runeMap := make(map[rune]int)
	// iterate through string
	for i, r := range s {
		// use constant lookup table to check for char
		// if conflict, move pointer up to next char, and keep searching
		if _, ok := runeMap[r]; ok && runeMap[r] >= start {
			start = runeMap[r] + 1
		}
		runeMap[r] = i
		// keep max of result and diff between i and pointer to start of string
		if result < i-start+1 {
			result = i - start + 1
		}
	}
	return result
}

func main() {
	println(lengthOfLongestSubstring("abcabcbb") == 3)
	println(lengthOfLongestSubstring("bbbbb") == 1)
	println(lengthOfLongestSubstring("pwwkew") == 3)
	println(lengthOfLongestSubstring("dvdf") == 3)
	println(lengthOfLongestSubstring("abba") == 2)
	println(lengthOfLongestSubstring("") == 0)
	println(lengthOfLongestSubstring("a") == 1)
	println(lengthOfLongestSubstring("b") == 1)
	println(lengthOfLongestSubstring("ab") == 2)
}

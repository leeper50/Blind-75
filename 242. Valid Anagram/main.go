// Result - Passed
// Runtime - Beats 100%
// Memory - Beats 73.93%

package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	freq := make([]int, 26)
	for _, r := range s {
		freq[r-'a']++
	}
	for _, r := range t {
		freq[r-'a']--
		if freq[r-'a'] < 0 {
			return false
		}
	}
	return true
}

func main() {
	println(isAnagram("ef", "d") == false)
	println(isAnagram("anagram", "nagaram") == true)
	println(isAnagram("rat", "car") == false)
}

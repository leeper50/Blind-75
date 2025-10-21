// Result - Passed
// Runtime - Beats 100%
// Memory - Beats 94.14%

package main

func isPalindrome(s string) bool {
	isAlphaNumeric := func(r rune) bool {
		if ('a' <= r && r <= 'z') ||
			('A' <= r && r <= 'Z') ||
			('0' <= r && r <= '9') {
			return true
		}
		return false
	}
	right := len(s) - 1
	for left := 0; left < right; {
		lc, rc := rune(s[left]), rune(s[right])
		if !isAlphaNumeric(lc) {
			left++
			continue
		}
		if !isAlphaNumeric(rc) {
			right--
			continue
		}
		// Check number range
		if (lc < 'A' || rc < 'A') && lc != rc {
			return false
		}
		// Check match including casing
		if lc != rc && lc+32 != rc && lc != rc+32 {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	println(isPalindrome("A man, a plan, a canal: Panama"), true)
	println(isPalindrome("a man, a plan, a canal: panama"), true)
	println(isPalindrome("race a car"), false)
	println(isPalindrome("race car"), true)
	println(isPalindrome(" "), true)
	println(isPalindrome(" e "), true)
	println(isPalindrome(" ee "), true)
	println(isPalindrome(" fe "), false)
	println(isPalindrome(" e1e "), true)
	println(isPalindrome("  e "), true)
	println(isPalindrome("."), true)
	println(isPalindrome("0P"), false)
	println(rune('0'))
	println(rune('P'))
	println(rune('9'))
}

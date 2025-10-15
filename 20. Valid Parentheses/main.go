// Result - Passed
// Runtime - Beats 100%
// Memory - Beats 92.47%

package main

func isValid(s string) bool {
	// Use stack to store open values
	// Pop open value with matching closed value
	// If doesn't match return false
	// If stack contains values at end return false

	// Use switch case with lookup map for more condensed code
	myStack := make([]rune, 0, len(s))
	pairMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, c := range s {
		switch c {
		case '(', '[', '{':
			myStack = append(myStack, c)
		case ')', ']', '}':
			if len(myStack) != 0 && pairMap[c] == myStack[len(myStack)-1] {
				myStack = myStack[0 : len(myStack)-1]
			} else {
				return false
			}
		}
	}
	if len(myStack) != 0 {
		return false
	}
	return true
}

func main() {
	println(isValid("()"), "true")
	println(isValid("()[]{}"), "true")
	println(isValid("(]"), "false")
	println(isValid("([])"), "true")
	println(isValid("([)]"), "false")
	println(isValid("("), "false")
	println(isValid("]"), "false")
	println(isValid("((("), "false")
}

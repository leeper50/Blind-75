// Result - Passed
// Runtime - Beats 93.74%
// Memory - Beats 51.34%

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	// Use two pointers
	// Slow moves 1 step per iteration
	// Fast moves 2 steps per iteration
	// If no cycle, fast will reach end (nil)
	// If cycle, fast will become equal to slow when it loops
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func main() {
	print("odd erm")
}

// Result - Passed
// Runtime - Beats 100%
// Memory - Beats 50.04%

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	result, p1 := head, head
	count := 0
	for head != nil {
		if count > n {
			p1 = p1.Next
		}
		count++
		head = head.Next
	}

	if count != n {
		p1.Next = p1.Next.Next
	} else {
		result = p1.Next
	}
	return result
}

func main() {
	val := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	result := removeNthFromEnd(val, 2)
	for result != nil {
		print(result.Val, " ")
		result = result.Next
	}
	println("\n1 2 3 5\n")

	val = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	result = removeNthFromEnd(val, 4)
	for result != nil {
		print(result.Val, " ")
		result = result.Next
	}
	println("\n1 3 4 5\n")

	val = &ListNode{1, &ListNode{2, nil}}
	result = removeNthFromEnd(val, 1)
	for result != nil {
		print(result.Val, " ")
		result = result.Next
	}
	println("\n1\n")

	val = &ListNode{1, &ListNode{2, nil}}
	result = removeNthFromEnd(val, 2)
	for result != nil {
		print(result.Val, " ")
		result = result.Next
	}
	println("\n2\n")

	println()
	val = &ListNode{1, nil}
	result = removeNthFromEnd(val, 1)
	for result != nil {
		print(result.Val)
		result = result.Next
	}
	println("\nnothing")
}

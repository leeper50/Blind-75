// Result - Passed
// Runtime - Beats 100%
// Memory - Beats 40.85%

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Use two pointers
	// Make pointer to dummy node
	// Make iterator pointer as current
	// Loop until either list is empty

	// If less than, set current.next to list1, set list1 to list1.next
	// If greater than, set current.next to list2, set list2 to list2.next
	// Move iterator pointer along

	// Append remaining list elements when loop finished
	// Return dummy.next
	dummy := &ListNode{}
	current := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return dummy.Next // head of sorted list
}

func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	l3 := mergeTwoLists(l1, l2)
	for l3 != nil {
		print(l3.Val, " ")
		l3 = l3.Next
	}
	println("\n[1 1 2 3 4 4]")
}

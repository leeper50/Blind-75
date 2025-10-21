// Result - Passed
// Runtime - Beats 100%
// Memory - Beats 58.33%

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseListIter(head *ListNode) *ListNode {
	var prev, curr, next *ListNode
	curr = head
	for curr != nil {
		next = curr.Next // Store next node
		curr.Next = prev // Reverse pointer
		// Move nodes forward
		prev = curr // Move previous to current
		curr = next // Move current to next element
	}
	return prev
}

func reverseListRec(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	revHead := reverseListRec(head.Next)
	head.Next.Next = head
	head.Next = nil
	return revHead
}

func main() {
	var result, myList *ListNode
	myList = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	result = reverseListIter(myList)
	println(result)
	myList = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	result = reverseListRec(myList)
	println(result)
}

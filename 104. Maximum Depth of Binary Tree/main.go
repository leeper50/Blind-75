// Result - Passed
// Runtime - Beats 100%
// Memory - Beats 92.44%

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	// Check if node is nil, return 0 if so
	// Check if node has children, call function on children
	// Keep largest level when going through tree and add 1
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func printValDFS(root *TreeNode) {
	// Check if node is nil, return 0 if so
	// Print value of node
	// Call function on node's children
	// This continues until all nodes have been processed
	if root == nil {
		return
	}
	println(root.Val)
	printValDFS(root.Left)
	printValDFS(root.Right)
}

func printValBFS(root *TreeNode) {
	// Check if root is nil, return if so
	// Create a queue
	// Use queue to contain all nodes on a level
	// i.e. root's children get put in queue
	// then append each node's append children to queue
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	for len(q) != 0 {
		for size := len(q); size != 0; size-- {
			curNode := q[0]
			q = q[1:]
			println(curNode.Val)
			if curNode.Left != nil {
				q = append(q, curNode.Left)
			}
			if curNode.Right != nil {
				q = append(q, curNode.Right)
			}
		}
	}
}

func main() {
	t := &TreeNode{3, &TreeNode{9, &TreeNode{35, nil, nil}, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	println(maxDepth(t), 3)
	println()
	printValDFS(t) // Go from left to right
	println()
	printValBFS(t) // Go level by level (left to right)
	t = &TreeNode{1, nil, &TreeNode{2, nil, nil}}
	println(maxDepth(t), 2)
	t = &TreeNode{}
	println(maxDepth(t), 1)
	t = &TreeNode{1, nil, nil}
	println(maxDepth(t), 1)
}

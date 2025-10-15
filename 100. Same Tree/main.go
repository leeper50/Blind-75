// Result - Passed
// Runtime - Beats 100%
// Memory - Beats 12.44%

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Check if trees are empty, return true
	// If one tree is empty but the other isn't return false
	// If they aren't empty but the values dont match return false
	// Loop over 1-2-3 until done
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || (p.Val != q.Val) {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	t1 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	t2 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	println(isSameTree(t1, t2), "true")
	t1 = &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	t2 = &TreeNode{1, nil, &TreeNode{2, nil, nil}}
	println(isSameTree(t1, t2), "false")
	t1 = &TreeNode{1, nil, &TreeNode{2, nil, nil}}
	t2 = &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	println(isSameTree(t1, t2), "false")
	t1 = &TreeNode{1, nil, nil}
	t2 = &TreeNode{1, nil, nil}
	println(isSameTree(t1, t2), "true")
	t1 = &TreeNode{}
	t2 = &TreeNode{}
	println(isSameTree(t1, t2), "true")
}

// Result - Passed
// Runtime - Beats 100%
// Memory - Beats 43.17%

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	if isIdentical(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isIdentical(r *TreeNode, s *TreeNode) bool {
	if r == nil && s == nil {
		return true
	}
	if r == nil || s == nil || (r.Val != s.Val) {
		return false
	}
	return isIdentical(r.Left, s.Left) && isIdentical(r.Right, s.Right)
}

func main() {
	bigTree := &TreeNode{3, &TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}}, &TreeNode{5, nil, nil}}
	smallTree := &TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}}
	println(isSubtree(bigTree, smallTree), true)
}

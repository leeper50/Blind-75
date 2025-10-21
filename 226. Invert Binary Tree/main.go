// Result - Passed
// Runtime - Beats 100%
// Memory - Beats 99.33%

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func main() {
	data := &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}
	invertTree(data)
	invertTree(nil)
}

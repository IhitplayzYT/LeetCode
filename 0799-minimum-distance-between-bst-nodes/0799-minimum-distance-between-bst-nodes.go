/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var mn, prev int

func minDiffInBST(root *TreeNode) int {
	mn = math.MaxInt
	prev = -1

	inorder(root)
	return mn
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}

	inorder(root.Left)
	if prev != -1 {
		mn = min(mn, root.Val-prev)
	}

	prev = root.Val

	inorder(root.Right)
}
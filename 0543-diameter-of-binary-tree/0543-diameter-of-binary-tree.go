/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     l *TreeNode
 *     r *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    mx := 0
    var height func(*TreeNode) int

    height = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        l,r := height(node.Left),height(node.Right)
        mx = max(mx, l+r)
        return 1 + max(l, r)
    }
    height(root)
    return mx
}
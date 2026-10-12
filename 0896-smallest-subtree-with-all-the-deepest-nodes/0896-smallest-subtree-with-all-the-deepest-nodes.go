func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
    node, _ := dfs(root)
    return node
}

func dfs(node *TreeNode) (*TreeNode, int) {
        if node == nil {
            return nil, 0
        }

        leftNode, leftDepth := dfs(node.Left)
        rightNode, rightDepth := dfs(node.Right)

        if leftDepth > rightDepth {
            return leftNode, leftDepth + 1
        }

        if rightDepth > leftDepth {
            return rightNode, rightDepth + 1
        }

        return node, leftDepth + 1
    }
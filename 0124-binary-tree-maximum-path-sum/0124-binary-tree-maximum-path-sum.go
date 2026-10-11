func maxPathSum(root *TreeNode) int {
    ret := root.Val
    recurse(root, &ret)
    return ret
}

func recurse(node *TreeNode, mx *int) int {
    if node == nil { return 0 }
    lmax := max(0, recurse(node.Left, mx))
    rmax := max(0, recurse(node.Right, mx))
    *mx = max(*mx, node.Val+lmax+rmax)
    return node.Val + max(lmax, rmax)
}
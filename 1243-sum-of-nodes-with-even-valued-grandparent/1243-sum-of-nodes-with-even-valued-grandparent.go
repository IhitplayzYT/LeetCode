func sumEvenGrandparent(root *TreeNode) int {
    ret := 0
    recurse(root, nil, nil, &ret)
    return ret
}

func recurse(node, parent, gp *TreeNode, ret *int) {
    if node == nil { return }
    if gp != nil && gp.Val%2 == 0 { *ret += node.Val }
    recurse(node.Left, node, parent, ret)
    recurse(node.Right, node, parent, ret)
}
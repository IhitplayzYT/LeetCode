func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	p_pth,q_pth := get_bst_path(root,p.Val),get_bst_path(root,q.Val)
	lp,lq := len(p_pth),len(q_pth)
	var ret *TreeNode = nil
	for i,j:=0,0;i < lp && j < lq && p_pth[i].Val == q_pth[j].Val;{
		ret = p_pth[i]
		i += 1
		j += 1 
	}
	return ret	
}

func get_bst_path(root *TreeNode,tgt int) []*TreeNode{
    pth := make([]*TreeNode,0)
	    for root != nil {
        pth = append(pth, root)

        if root.Val == tgt {
            return pth
        } else if root.Val > tgt {
            root = root.Left
        } else {
            root = root.Right
        }
    }
    return pth
}


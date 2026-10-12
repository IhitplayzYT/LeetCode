func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	p_pth,q_pth := get_path(root,p.Val),get_path(root,q.Val)
	lp,lq := len(p_pth),len(q_pth)
	var ret *TreeNode = nil
	for i,j:=0,0;i < lp && j < lq && p_pth[i].Val == q_pth[j].Val;{
		ret = p_pth[i]
		i += 1
		j += 1 
	}
	return ret	
}

func get_path(node * TreeNode,tgt int) []*TreeNode{
  path := []*TreeNode{}

    var dfs func(*TreeNode) bool
    dfs = func(node *TreeNode) bool {
        if node == nil {
            return false
        }

        path = append(path, node)

        if node.Val == tgt {
            return true
        }

        if dfs(node.Left) || dfs(node.Right) {
            return true
        }

        path = path[:len(path)-1]
        return false
    }

    if !dfs(node) {
        return nil
    }

    return path
}


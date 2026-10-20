func cloneGraph(node *Node) *Node {
	if node == nil { return nil }
	copies := make([]*Node, 101)
	dfs(node, copies)
	return copies[node.Val]
}

func dfs(nd *Node, copies []*Node) {
	nn := new(Node)
	nn.Val = nd.Val
	copies[nd.Val] = nn
	for _, n := range nd.Neighbors {
		if copies[n.Val] == nil { dfs(n, copies) }
		nn.Neighbors = append(nn.Neighbors, copies[n.Val])
	}
}
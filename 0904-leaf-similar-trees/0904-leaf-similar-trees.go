/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
   	list1 := make([]int, 0)
	list2 := make([]int, 0)
	recursion(root1, &list1)
	recursion(root2, &list2)
    l1,l2 := len(list1),len(list2)
    i,j:=0,0
    for ; i < l1 && j < l2;{
        if list1[i] != list2[j] { return false}
        i++
        j++
    } 
    return i == l1 && j == l2
}


func recursion(node *TreeNode, ll *[]int) {
	if node == nil { return }
	if node.Left == nil && node.Right == nil {
		*ll = append(*ll, node.Val)
		return
	}
	recursion(node.Left, ll)
	recursion(node.Right, ll)
}
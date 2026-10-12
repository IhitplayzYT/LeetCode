/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
    res:= []int{}
    add(root1, &res)
    add(root2, &res)
    sort.Ints(res)
    return res
}

func add(r *TreeNode, arr *[]int){
    if r==nil{
        return
    }
    *arr= append(*arr,r.Val)
    add(r.Left,arr)
    add(r.Right,arr)
}
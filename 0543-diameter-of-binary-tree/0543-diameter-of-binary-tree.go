/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    mx := 0
    recurse(root,&mx)
    return mx
}

func recurse(root *TreeNode,mx *int){
    if root == nil {return}
    *mx = max(*mx,diameter(root))
    recurse(root.Left,mx)
    recurse(root.Right,mx)
}


func diameter(node *TreeNode) int{
    if node == nil {return 0}
    return height(node.Left) + height(node.Right)
}

func height(node * TreeNode) int {
    if node == nil {return 0}
    return 1 + max(height(node.Left),height(node.Right))
}





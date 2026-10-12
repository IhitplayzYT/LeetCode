/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type FindElements struct {
    cache map[int]int
}


func Constructor(root *TreeNode) FindElements {
    ret := FindElements{cache:make(map[int]int,0)}
    build(root,0,&ret.cache)
    return ret
}


func (this *FindElements) Find(target int) bool {
    return this.cache[target] != 0 
}

func build(root*TreeNode,idx int,mp * map[int]int){
    if root == nil {return}
    root.Val = idx
    (*mp)[idx] += 1
    build(root.Left,2 * idx + 1,mp)
    build(root.Right,2 * idx + 2,mp)
}



/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
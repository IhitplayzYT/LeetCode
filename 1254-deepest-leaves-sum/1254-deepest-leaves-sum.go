/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, error) {
	var zero T
	if q.IsEmpty() { return zero, errors.New("queue is empty")}
	item := q.items[0]
	q.items = q.items[1:] 
	return item, nil
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func deepestLeavesSum(root *TreeNode) int {
    if root == nil {return 0}
    que := Queue[*TreeNode]{}
    que.Enqueue(root)
    var ret int = 0
    for ;len(que.items) != 0;{
        l := len(que.items)
        if len(que.items) != 0{
            ret = 0
            for _,v := range que.items{ 
                ret += v.Val;
            }
        }
        for ;l != 0;l--{
            z,_ := que.Dequeue()
            if z.Left != nil{
                que.Enqueue(z.Left)
            }
            if z.Right != nil{
                que.Enqueue(z.Right)
            }
        }

    }
    return ret 
}
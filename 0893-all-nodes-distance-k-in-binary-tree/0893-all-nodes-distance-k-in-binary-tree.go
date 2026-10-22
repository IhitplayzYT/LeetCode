/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type IntQueue struct {
	items []int
}

// Enqueue adds an element to the back of the queue
func (q *IntQueue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the front element of the queue
func (q *IntQueue) Dequeue() (int, error) {
	if len(q.items) == 0 {
		return 0, errors.New("queue is empty")
	}
	
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

// Size returns the current number of elements
func (q *IntQueue) Size() int {
	return len(q.items)
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
    adj := make(map[int][]int,0)
    build_adjacency(root,nil,&adj)
    que := &IntQueue{}
    que.Enqueue(target.Val)
    seen := map[int]bool{target.Val: true}

    for ;k != 0;k--{
        l := que.Size()
        for ;l!=0;l--{
            k,_ := que.Dequeue()
            for _,v := range adj[k]{
                if !seen[v] {
                    seen[v] = true
                    que.Enqueue(v)
                }
            }
        }
    }
    return que.items
}

func build_adjacency(n * TreeNode,p * TreeNode,adj *map[int][]int){
    if n == nil {return}
    _,ok := (*adj)[n.Val]
    if !ok{ (*adj)[n.Val] = make([]int,0) }
    if p != nil {(*adj)[n.Val] = append((*adj)[n.Val], p.Val)}
    if n.Left != nil {(*adj)[n.Val] = append((*adj)[n.Val], n.Left.Val)}
    if n.Right != nil {(*adj)[n.Val] = append((*adj)[n.Val], n.Right.Val)}
    build_adjacency(n.Left,n,adj)
    build_adjacency(n.Right,n,adj)
}
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

import (
	"errors"
	"fmt"
)

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

func levelOrderBottom(root *TreeNode) [][]int {
    if root == nil {
		return [][]int{}
	}
    ret := make([][]int,0)
    que := &Queue[*TreeNode]{}
	que.Enqueue(root)
	for ;!que.IsEmpty();{
		buff := []int{}		
		l := len(que.items)
		for ;l != 0;l--{
			k,_ := que.Dequeue()
			if k.Left != nil {que.Enqueue(k.Left)}
			if k.Right != nil {que.Enqueue(k.Right)}
			buff = append(buff, k.Val)
		}
		ret = append(ret, buff)
	}
    slices.Reverse(ret)
	return ret
}
type BSTIterator struct {
    stack []*TreeNode;
	cur *TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{stack: make([]*TreeNode, 0),cur:root}
}


func (self *BSTIterator) Next() int {
   if self.cur != nil || len(self.stack) > 0{
		for ;self.cur != nil;{
			self.stack = append(self.stack, self.cur)
			self.cur = self.cur.Left
		}
		self.cur = self.stack[len(self.stack)-1]
		self.stack = self.stack[:len(self.stack)-1]
		ret := self.cur.Val
		self.cur = self.cur.Right
		return ret
   }
   return -1
}


func (self *BSTIterator) HasNext() bool {
	return self.cur != nil || len(self.stack) > 0
}
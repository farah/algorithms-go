package hey

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	stack := make([]*TreeNode, 0)
	o := BSTIterator{stack: stack}
	o.LeftmostInorder(root)
	return o
}

func (this *BSTIterator) LeftmostInorder(root *TreeNode) {
	for root != nil {
		this.stack = append(this.stack, root)
		root = root.Left
	}
}

func (this *BSTIterator) Next() int {
	topmostNode := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if topmostNode.Right != nil {
		this.LeftmostInorder(topmostNode.Right)
	}
	return topmostNode.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

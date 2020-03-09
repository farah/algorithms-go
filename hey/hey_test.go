package hey

import (
	"testing"
)

/* TreeNodeA
             10
           /   \
          8     15
         /     / \
	      6     12   17
	     /
	    5
*/

// 5, 6, 8, 10, 12, 15, 17

func TestLeftView(t *testing.T) {
	root := &TreeNode{Val: 10, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 8, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 6, Left: nil, Right: nil}
	root.Left.Left.Left = &TreeNode{Val: 5, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 15, Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 17, Left: nil, Right: nil}
	root.Right.Left = &TreeNode{Val: 12, Left: nil, Right: nil}

	obj := Constructor(root)
	param_1 := obj.Next()
	param_2 := obj.Next()
	param_3 := obj.Next()

}

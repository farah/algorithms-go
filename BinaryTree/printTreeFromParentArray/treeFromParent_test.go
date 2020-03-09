package spiralOrder

import (
	"testing"
)

/*
              0
           /     \
          1       2
         /       / \
	      3       4   5
	             / \
              6   7

*/
func TestCreateTreeFromParent(t *testing.T) {
	root := &Node{Value: 0, Left: nil, Right: nil}
	root.Left = &Node{Value: 1, Left: nil, Right: nil}
	root.Right = &Node{Value: 2, Left: nil, Right: nil}
	root.Left.Left = &Node{Value: 3, Left: nil, Right: nil}
	root.Right.Left = &Node{Value: 4, Left: nil, Right: nil}
	root.Right.Right = &Node{Value: 5, Left: nil, Right: nil}
	root.Right.Left.Left = &Node{Value: 6, Left: nil, Right: nil}
	root.Right.Left.Right = &Node{Value: 7, Left: nil, Right: nil}

	parent := []int{-1, 0, 0, 1, 2, 2, 4, 4}

	result := createTreeFromParent(parent)

	_ = result
}

/*
             45
           /   \
          6     3
         /     / \
	      72     2   6
	     /
	    16


	root := &Node{Value: 45, Left: nil, Right: nil}
	root.Left = &Node{Value: 6, Left: nil, Right: nil}
	root.Left.Left = &Node{Value: 72, Left: nil, Right: nil}
	root.Right = &Node{Value: 3, Left: nil, Right: nil}
	root.Right.Right = &Node{Value: 6, Left: nil, Right: nil}
	root.Right.Left = &Node{Value: 2, Left: nil, Right: nil}
	root.Right.Left.Left = &Node{Value: 16, Left: nil, Right: nil}

*/

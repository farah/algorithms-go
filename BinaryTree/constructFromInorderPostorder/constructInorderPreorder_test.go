package identicalTree

import (
	"testing"
)

/* Consider below tree
          1
        /   \
       /     \
      2       3
     /       / \
    /       /   \
   4       5     6
          / \
         /   \
        7     8
*/

func TestIdenticalTree(t *testing.T) {
	root := &Node{Value: 1, Left: nil, Right: nil}
	root.Left = &Node{Value: 2, Left: nil, Right: nil}
	root.Left.Left = &Node{Value: 4, Left: nil, Right: nil}
	root.Right = &Node{Value: 3, Left: nil, Right: nil}
	root.Right.Right = &Node{Value: 6, Left: nil, Right: nil}
	root.Right.Left = &Node{Value: 5, Left: nil, Right: nil}
	root.Right.Left.Left = &Node{Value: 7, Left: nil, Right: nil}
	root.Right.Left.Right = &Node{Value: 8, Left: nil, Right: nil}

	result := constructInorderPreorder()

	if result != nil {
		t.Errorf("Expected true, but it was %v instead.", result)
	}

}

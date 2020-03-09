package spiralOrder

import (
	"strings"
	"testing"
)

/*
              A
           /     \
          B       C
         / \     / \
	      D   E   F   6
	     / \     / \   \
      H   I   J   K   L
     / \     / \     / \
    M   N   P   Q   R   S
*/
func TestLeftView(t *testing.T) {
	root := &Node{Value: "A", Left: nil, Right: nil}
	root.Left = &Node{Value: "B", Left: nil, Right: nil}
	root.Right = &Node{Value: "C", Left: nil, Right: nil}
	root.Left.Left = &Node{Value: "D", Left: nil, Right: nil}
	root.Left.Right = &Node{Value: "E", Left: nil, Right: nil}
	root.Right.Left = &Node{Value: "F", Left: nil, Right: nil}
	root.Right.Right = &Node{Value: "G", Left: nil, Right: nil}
	root.Left.Left.Left = &Node{Value: "H", Left: nil, Right: nil}
	root.Left.Left.Right = &Node{Value: "I", Left: nil, Right: nil}
	root.Right.Left.Left = &Node{Value: "J", Left: nil, Right: nil}
	root.Right.Left.Right = &Node{Value: "K", Left: nil, Right: nil}
	root.Right.Right.Right = &Node{Value: "L", Left: nil, Right: nil}
	root.Left.Left.Left.Left = &Node{Value: "M", Left: nil, Right: nil}
	root.Left.Left.Left.Right = &Node{Value: "N", Left: nil, Right: nil}
	root.Right.Left.Left.Left = &Node{Value: "P", Left: nil, Right: nil}
	root.Right.Left.Left.Right = &Node{Value: "Q", Left: nil, Right: nil}
	root.Right.Right.Right.Left = &Node{Value: "R", Left: nil, Right: nil}
	root.Right.Right.Right.Right = &Node{Value: "S", Left: nil, Right: nil}

	result := spiralOrder(root)
	expect := []string{
		"A", "B", "C", "G", "F", "E",
		"D", "H", "I", "J", "K",
		"L", "S", "R", "Q", "P", "N", "M",
	}
	a := strings.Join(result, "")
	b := strings.Join(expect, "")
	if a != b {
		t.Errorf("Expected true, but it was %v instead.", false)
	}

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

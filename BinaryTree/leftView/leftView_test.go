package leftView

import (
	"reflect"
	"testing"
)

/* NodeA
             1
           /   \
          2     3
         /     / \
	      4     5   6
	     /
	    7
*/
func TestLeftView(t *testing.T) {
	root := &Node{Value: 1, Left: nil, Right: nil}
	root.Left = &Node{Value: 2, Left: nil, Right: nil}
	root.Left.Left = &Node{Value: 4, Left: nil, Right: nil}
	root.Right = &Node{Value: 3, Left: nil, Right: nil}
	root.Right.Right = &Node{Value: 6, Left: nil, Right: nil}
	root.Right.Left = &Node{Value: 5, Left: nil, Right: nil}
	root.Right.Left.Left = &Node{Value: 7, Left: nil, Right: nil}

	result := leftViewMain(root)

	expect := map[int][]int{
		1: {1},
		2: {2},
		3: {4},
		4: {7},
	}

	if !reflect.DeepEqual(result, expect) {
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

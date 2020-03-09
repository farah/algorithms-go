package identicalTree

import (
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

/* NodeB
             1
           /   \
          2     3
         /       \
	      4         6
	     /
	    7
*/
func TestIdenticalTree(t *testing.T) {
	root := &Node{Value: 1, Left: nil, Right: nil}
	root.Left = &Node{Value: 2, Left: nil, Right: nil}
	root.Left.Left = &Node{Value: 4, Left: nil, Right: nil}
	root.Right = &Node{Value: 3, Left: nil, Right: nil}
	root.Right.Right = &Node{Value: 6, Left: nil, Right: nil}
	root.Right.Left = &Node{Value: 5, Left: nil, Right: nil}
	root.Right.Left.Left = &Node{Value: 7, Left: nil, Right: nil}

	nodeA := root
	nodeB := root

	result := identicalTree(nodeA, nodeB)

	if result != true {
		t.Errorf("Expected true, but it was %v instead.", result)
	}

}

func TestNotIdenticalTree(t *testing.T) {
	//NodeA
	root := &Node{Value: 1, Left: nil, Right: nil}
	root.Left = &Node{Value: 2, Left: nil, Right: nil}
	root.Left.Left = &Node{Value: 4, Left: nil, Right: nil}
	root.Right = &Node{Value: 3, Left: nil, Right: nil}
	root.Right.Right = &Node{Value: 6, Left: nil, Right: nil}
	root.Right.Left = &Node{Value: 5, Left: nil, Right: nil}
	root.Right.Left.Left = &Node{Value: 7, Left: nil, Right: nil}

	nodeA := root

	//NodeB

	root = &Node{Value: 1, Left: nil, Right: nil}
	root.Left = &Node{Value: 2, Left: nil, Right: nil}
	root.Left.Left = &Node{Value: 4, Left: nil, Right: nil}
	root.Right = &Node{Value: 3, Left: nil, Right: nil}
	root.Right.Right = &Node{Value: 6, Left: nil, Right: nil}
	root.Right.Left = &Node{Value: 5, Left: nil, Right: nil}
	root.Left.Left.Left = &Node{Value: 7, Left: nil, Right: nil}

	nodeB := root

	result := identicalTree(nodeA, nodeB)

	if result != false {
		t.Errorf("Expected true, but it was %v instead.", result)
	}

}

package levelOrderTraversalPackage

import (
	"reflect"
	"testing"
)

/*
             1
           /   \
          2     3
         /     / \
	      4     5   6
	     /
	    7
*/

func TestLevelOrderTraversal(t *testing.T) {
	root := &Node{Value: 1, Left: nil, Right: nil}
	root.Left = &Node{Value: 2, Left: nil, Right: nil}
	root.Left.Left = &Node{Value: 4, Left: nil, Right: nil}
	root.Right = &Node{Value: 3, Left: nil, Right: nil}
	root.Right.Right = &Node{Value: 6, Left: nil, Right: nil}
	root.Right.Left = &Node{Value: 5, Left: nil, Right: nil}
	root.Right.Left.Left = &Node{Value: 7, Left: nil, Right: nil}

	result := levelOrderTraversal(root)

	if result != nil {
		t.Errorf("Expected true, but it was %v instead.", result)
	}

}

func TestLevelOrderTraversalDictionary(t *testing.T) {
	root := &Node{Value: 1, Left: nil, Right: nil}
	root.Left = &Node{Value: 2, Left: nil, Right: nil}
	root.Left.Left = &Node{Value: 4, Left: nil, Right: nil}
	root.Right = &Node{Value: 3, Left: nil, Right: nil}
	root.Right.Right = &Node{Value: 6, Left: nil, Right: nil}
	root.Right.Left = &Node{Value: 5, Left: nil, Right: nil}
	root.Right.Left.Left = &Node{Value: 7, Left: nil, Right: nil}

	result := levelOrderMain(root)

	expect := map[int][]int{
		1: {1},
		2: {2, 3},
		3: {4, 5, 6},
		4: {7},
	}

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Expected true, but it was %v instead.", false)
	}

}

func TestLevelOrderTraversalRecursive(t *testing.T) {
	root := &Node{Value: 1, Left: nil, Right: nil}
	root.Left = &Node{Value: 2, Left: nil, Right: nil}
	root.Left.Left = &Node{Value: 4, Left: nil, Right: nil}
	root.Right = &Node{Value: 3, Left: nil, Right: nil}
	root.Right.Right = &Node{Value: 6, Left: nil, Right: nil}
	root.Right.Left = &Node{Value: 5, Left: nil, Right: nil}
	root.Right.Left.Left = &Node{Value: 7, Left: nil, Right: nil}

	levelOrderRecursive(root)

}

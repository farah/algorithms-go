package spiralOrder

import (
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

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

func createTreeFromParent(parent []int) (list []string) {
	return nil
}

func printTree(root *Node) *Node {
	if root == nil {
		return root
	}
	printTree(root.Left)
	fmt.Println(root.Value)
	printTree(root.Right)
	return root
}

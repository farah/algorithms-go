package spiralOrder

import (
	"fmt"
)

type Node struct {
	Value string
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
func printTree(root *Node) *Node {
	if root == nil {
		return root
	}

	printTree(root.Left)

	fmt.Println(root.Value)

	printTree(root.Right)

	return root
}

func spiralOrder(root *Node) (list []string) {
	var stackA []*Node
	var stackB []*Node
	var curr *Node

	stackA = append(stackA, root)

	for len(stackA) > 0 || len(stackB) > 0 {
		for len(stackA) > 0 {
			curr, stackA = stackA[len(stackA)-1], stackA[:len(stackA)-1]
			list = append(list, curr.Value)

			if curr.Right != nil {
				stackB = append(stackB, curr.Right)
			}

			if curr.Left != nil {
				stackB = append(stackB, curr.Left)
			}

		}

		for len(stackB) > 0 {
			curr, stackB = stackB[len(stackB)-1], stackB[:len(stackB)-1]

			list = append(list, curr.Value)

			if curr.Left != nil {
				stackA = append(stackA, curr.Left)
			}

			if curr.Right != nil {
				stackA = append(stackA, curr.Right)
			}
		}
	}
	return list
}

func heightOfTree(root *Node) int {
	if root == nil {
		return 0
	}
	lDepth := heightOfTree(root.Left)
	rDepth := heightOfTree(root.Right)

	if lDepth > rDepth {
		return (lDepth + 1)
	}
	return rDepth + 1
}

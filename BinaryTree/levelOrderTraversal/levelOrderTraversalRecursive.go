package levelOrderTraversalPackage

import (
	"fmt"
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

func levelOrderRecursive(root *Node) {
	height := heightOfTree(root)
	for i := 1; i <= height; i++ {
		traverse(root, i)
	}
}

func traverse(root *Node, level int) {
	if root != nil {
		if level == 1 {
			//  printf("%d ", root->info);
			fmt.Printf("value: %d", root.Value)
		} else if level > 1 {
			traverse(root.Left, level-1)
			traverse(root.Right, level-1)
		}
	}
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

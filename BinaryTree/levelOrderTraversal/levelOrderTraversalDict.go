package levelOrderTraversalPackage

/* NodeA
             1
           /   \
          2     3
         /     / \
	      4     5   6
	     /
	    7
*/

func levelOrderMain(root *Node) map[int][]int {
	dict := make(map[int][]int)
	levelOrder(root, 1, &dict)
	return dict
}

func levelOrder(root *Node, level int, dict *map[int][]int) {
	if root == nil {
		return
	}

	(*dict)[level] = append((*dict)[level], root.Value)
	levelOrder(root.Left, level+1, dict)
	levelOrder(root.Right, level+1, dict)
}

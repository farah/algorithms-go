package leftView

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

/* NodeA
             1
           /   \
          2     3
         /     / \
	      4     5   6
	     /
	    7
*/

func leftViewMain(root *Node) map[int][]int {
	dict := make(map[int][]int)
	leftView(root, 1, &dict)

	for key, list := range dict {
		dict[key] = []int{list[0]}
	}
	return dict
}

func leftView(root *Node, level int, dict *map[int][]int) {
	if root == nil {
		return
	}

	(*dict)[level] = append((*dict)[level], root.Value)
	leftView(root.Left, level+1, dict)
	leftView(root.Right, level+1, dict)
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

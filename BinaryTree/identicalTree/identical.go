package identicalTree

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

/* NodeB
             1
           /   \
          2     3
         /       \
	      4         6
	     /
	    7
*/

func identicalTree(nodeA *Node, nodeB *Node) bool {
	if nodeA == nil && nodeB == nil {
		return true
	}
	if nodeA != nil && nodeB != nil {
		if nodeA.Value == nodeB.Value && identicalTree(nodeA.Left, nodeB.Left) && identicalTree(nodeA.Right, nodeB.Right) {
			return true
		}
	}

	return false
}

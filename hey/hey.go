package hey

//Solution from https://www.youtube.com/watch?v=JdmAYw5h3G8

// Case1: If the node has a right subtree, find the least value
// node from that that right subtree.
// Case2: If the node does not have right subtree, search
// that 'p' from the root. The node from where we take the last left
// is the answer

type Node struct {
	Root  *Node
	Left  *Node
	Right *Node
	Data  int
}

func (node *Node) Next(root *Node, n *Node) *Node {
	//Case 1
	if n.Right != nil {
		return minValue(n.Right)
	}

	var succ *Node

	//Case2
	for root != nil {
		if n.Data < root.Data {
			succ = root
			root = root.Left
		} else if n.Data > root.Data {
			root = root.Right
		} else {
			break
		}

	}

	return succ
}

func minValue(node *Node) *Node {
	current := node

	for current.Left != nil {
		current = current.Left
	}
	return current
}

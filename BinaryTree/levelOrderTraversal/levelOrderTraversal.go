package levelOrderTraversalPackage

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

func levelOrderTraversal(node *Node) [][]int {
	var curr *Node
	var level []int
	var matrix [][]int
	var queue []*Node

	levelCount := 0
	queue = append(queue, node)

	for len(queue) > 0 {
		size := len(queue)
		level = []int{}

		for i := 0; i < size; i++ {
			curr = queue[0]
			queue = queue[1:]
			level = append(level, curr.Value)

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		matrix = append(matrix, level)
		levelCount++
	}
	return matrix
}

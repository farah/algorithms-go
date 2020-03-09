package identicalTree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func construct(start int, end int, preorder []int, pIndex *int, dict map[int]int) *Node {

	if start > end {
		return nil
	}

	root := &Node{Value: preorder[*pIndex]}

	*pIndex = *pIndex + 1

	index := dict[root.Value]

	root.Left = construct(start, index-1, preorder, pIndex, dict)

	root.Right = construct(index+1, end, preorder, pIndex, dict)

	return root
}

func constructInorderPreorder() *Node {

	inorder := []int{4, 2, 1, 7, 5, 8, 3, 6}
	preorder := []int{1, 2, 4, 3, 5, 7, 8, 6}

	n := len(inorder)
	dict := make(map[int]int, 0)

	for i := 0; i < n; i++ {
		dict[inorder[i]] = i
	}

	pIndex := 0

	root := construct(0, n-1, preorder, &pIndex, dict)
	return root

}

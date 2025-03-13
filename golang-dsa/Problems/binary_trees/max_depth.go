package binarytrees


func MaxDepth[d comparable](tree *BinaryTree[d]) int64 {
	if tree == nil {
		return 0
	}

	leftDepth := MaxDepth(tree.left)
	rightDepth := MaxDepth(tree.right)
	
	depth := leftDepth
	if leftDepth < rightDepth {
		depth = rightDepth
	}

	return depth + 1
}
package binarytrees

import "math"


func IsBalanced[d comparable](tree *BinaryTree[d]) bool {
	return isBalanced(tree) != -1
}

func isBalanced[d comparable](tree *BinaryTree[d]) int64 {
	if tree == nil {
		return 0
	}

	leftDepth := isBalanced(tree.left)
	if leftDepth == -1 {
		return -1
	}
	rightDepth := isBalanced(tree.right)
	if rightDepth == -1 {
		return -1
	}
	
	if math.Abs(float64(leftDepth)- float64(rightDepth)) > 1 {
		return -1 
	}

	depth := leftDepth
	if leftDepth < rightDepth {
		depth = rightDepth
	}

	return depth + 1
}
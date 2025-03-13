package binarytrees



func Diameter[d comparable](tree *BinaryTree[d]) (int64) {
	_, diameter := maxDiameter(tree)
	return diameter
}

func maxDiameter[d comparable](tree *BinaryTree[d]) (depth int64, diameter int64) {
	if tree == nil {
		return 0, 0
	}

	leftDepth, leftDiameter := maxDiameter(tree.left)
	rightDepth, rightDiameter := maxDiameter(tree.right)

	depth = leftDepth
	if leftDepth < rightDepth {
		depth = rightDepth
	}

	diameter = leftDiameter
	if leftDiameter < rightDiameter {
		diameter = rightDiameter
	}

	currentDiameter := leftDepth + rightDepth

	if currentDiameter > diameter{
		diameter = currentDiameter
	}

	return depth + 1, diameter
}
package binarytrees

type BinaryTree[d comparable] struct {
	Data  d
	left  *BinaryTree[d]
	right *BinaryTree[d]
}

func NewTree[d comparable](data d) *BinaryTree[d] {
    return &BinaryTree[d]{
        Data:  data,
        left:  nil,
        right: nil,
    }
}

// Left returns the left child of the tree
func (t *BinaryTree[d]) Left() *BinaryTree[d] {
    return t.left
}

// Right returns the right child of the tree
func (t *BinaryTree[d]) Right() *BinaryTree[d] {
    return t.right
}

// SetLeft sets the left child of the tree
func (t *BinaryTree[d]) SetLeft(left *BinaryTree[d]) {
    t.left = left
}

// SetRight sets the right child of the tree
func (t *BinaryTree[d]) SetRight(right *BinaryTree[d]) {
    t.right = right
}
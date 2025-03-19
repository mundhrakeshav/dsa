package binarysearchtree

import (
	"golang.org/x/exp/constraints"
)

type BST[T constraints.Ordered] struct {
	Root *Node[T]
}

type Node[T constraints.Ordered] struct {
	Left, Right *Node[T]
	Value       T
}

// Public functions

func NewBST[T constraints.Ordered]() *BST[T] {
	return &BST[T]{}
}

func (b *BST[T]) Find(t T) *Node[T] {
	return b.Root.find(t)
}

func (b *BST[T]) FindMaxNode() *Node[T] {
	if b.Root == nil {
		return b.Root
	}
	return b.Root.findMaxNode()
}

func (b *BST[T]) InOrder() []T {
	return b.Root.inOrder(make([]T, 0))
}

func (b *BST[T]) InOrderNodes() []*Node[T] {
	return b.Root.inOrderNodes(make([]*Node[T], 0))
}

func (b *BST[T]) InsertRecursive(v T) {
	if b.Root == nil {
		b.Root = &Node[T]{Value: v}
		return
	}
	b.Root.insertRecursive(v)
}

func (b *BST[T]) Delete(v T) {
	b.Root = b.Root.deleteNode(v)
}

// Private functions
func (n *Node[T]) deleteNode(v T) *Node[T] {
	if n == nil {
		return nil
	}

	switch {
	case v < n.Value:
		n.Left = n.Left.deleteNode(v)
	case v > n.Value:
		n.Right = n.Right.deleteNode(v)
	default:

		if n.Left == nil && n.Right == nil {
			return nil
		}

		if n.Left == nil {
			return n.Right
		}
		if n.Right == nil {
			return n.Left
		}

		successor := n.Right.findMinNode()
		n.Value = successor.Value
		n.Right = n.Right.deleteNode(successor.Value)
	}

	return n
}

func (b *Node[T]) find(t T) *Node[T] {
	switch {
	case b == nil:
		return nil
	case b.Value == t:
		return b
	case b.Value > t:
		return b.Left.find(t)
	default:
		return b.Right.find(t)
	}
}

func (n *Node[T]) findMinNode() *Node[T] {
	if n.Left == nil {
		return n
	}
	return n.Left.findMinNode()
}

func (b *Node[T]) findMaxNode() *Node[T] {
	if b.Right == nil {
		return b
	}
	return b.Right.findMaxNode()
}

func (b *Node[T]) inOrder(t []T) []T {
	if b == nil {
		return t
	}
	t = b.Left.inOrder(t)
	t = append(t, b.Value)
	t = b.Right.inOrder(t)
	return t
}

func (b *Node[T]) inOrderNodes(t []*Node[T]) []*Node[T] {
	if b == nil {
		return t
	}
	t = b.Left.inOrderNodes(t)
	t = append(t, b)
	t = b.Right.inOrderNodes(t)
	return t
}

func (n *Node[T]) insertRecursive(v T) {
	if v <= n.Value {
		if n.Left == nil {
			n.Left = &Node[T]{Value: v}
			return
		}
		n.Left.insertRecursive(v)
		return
	}
	if n.Right == nil {
		n.Right = &Node[T]{Value: v}
		return
	}
	n.Right.insertRecursive(v)
}

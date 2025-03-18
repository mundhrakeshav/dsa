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

func NewBST[T constraints.Ordered]() *BST[T] {
	return &BST[T]{}
}

func (b *BST[T]) InsertRecursive(v T) {
	if b.Root == nil {
		b.Root = &Node[T]{Value: v}
		return
	}
	b.Root.insertRecursive(v)
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

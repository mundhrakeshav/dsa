package binarytrees_test

import (
	"testing"

	binarytrees "github.com/mundhrakeshav/DSA/golang-dsa/Problems/binary_trees"
)

func TestMaxDepth(t *testing.T) {
    tests := []struct {
        name     string
        buildTree func() *binarytrees.BinaryTree[int]
        expected int64
    }{
        {
            name: "nil tree",
            buildTree: func() *binarytrees.BinaryTree[int] {
                return nil
            },
            expected: 0,
        },
        {
            name: "single node",
            buildTree: func() *binarytrees.BinaryTree[int] {
                return binarytrees.NewTree(1)
            },
            expected: 1,
        },
        {
            name: "balanced tree",
            buildTree: func() *binarytrees.BinaryTree[int] {
                root := binarytrees.NewTree(1)
                root.SetLeft(binarytrees.NewTree(2))
                root.SetRight(binarytrees.NewTree(3))
                root.Left().SetLeft(binarytrees.NewTree(4))
                root.Left().SetRight(binarytrees.NewTree(5))
                root.Right().SetLeft(binarytrees.NewTree(6))
                root.Right().SetRight(binarytrees.NewTree(7))
                return root
            },
            expected: 3,
        },
        {
            name: "left-heavy tree",
            buildTree: func() *binarytrees.BinaryTree[int] {
                root := binarytrees.NewTree(1)
                root.SetLeft(binarytrees.NewTree(2))
                root.Left().SetLeft(binarytrees.NewTree(3))
                root.Left().Left().SetLeft(binarytrees.NewTree(4))
                return root
            },
            expected: 4,
        },
        {
            name: "right-heavy tree",
            buildTree: func() *binarytrees.BinaryTree[int] {
                root := binarytrees.NewTree(1)
                root.SetRight(binarytrees.NewTree(2))
                root.Right().SetRight(binarytrees.NewTree(3))
                root.Right().Right().SetRight(binarytrees.NewTree(4))
                return root
            },
            expected: 4,
        },
        {
            name: "complex tree",
            buildTree: func() *binarytrees.BinaryTree[int] {
                root := binarytrees.NewTree(1)
                root.SetLeft(binarytrees.NewTree(2))
                root.SetRight(binarytrees.NewTree(3))
                root.Left().SetLeft(binarytrees.NewTree(4))
                root.Right().SetRight(binarytrees.NewTree(5))
                root.Right().Right().SetRight(binarytrees.NewTree(6))
                return root
            },
            expected: 4,
        },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            tree := test.buildTree()
            got := binarytrees.MaxDepth(tree)
            if got != test.expected {
                t.Errorf("MaxDepth() = %v, want %v", got, test.expected)
            }
        })
    }
}
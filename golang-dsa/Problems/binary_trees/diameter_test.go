package binarytrees_test

import (
	"testing"

	binarytrees "github.com/mundhrakeshav/DSA/golang-dsa/Problems/binary_trees"
)

func TestDiameter(t *testing.T) {
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
            expected: 0,
        },
        {
            name: "line - right",
            buildTree: func() *binarytrees.BinaryTree[int] {
                root := binarytrees.NewTree(1)
                root.SetRight(binarytrees.NewTree(2))
                root.Right().SetRight(binarytrees.NewTree(3))
                return root
            },
            expected: 2,
        },
        {
            name: "line - left",
            buildTree: func() *binarytrees.BinaryTree[int] {
                root := binarytrees.NewTree(1)
                root.SetLeft(binarytrees.NewTree(2))
                root.Left().SetLeft(binarytrees.NewTree(3))
                return root
            },
            expected: 2,
        },
        {
            name: "balanced tree",
            buildTree: func() *binarytrees.BinaryTree[int] {
                root := binarytrees.NewTree(1)
                root.SetLeft(binarytrees.NewTree(2))
                root.SetRight(binarytrees.NewTree(3))
                return root
            },
            expected: 2,
        },
        {
            name: "diameter passes through root",
            buildTree: func() *binarytrees.BinaryTree[int] {
                root := binarytrees.NewTree(1)
                root.SetLeft(binarytrees.NewTree(2))
                root.SetRight(binarytrees.NewTree(3))
                root.Left().SetLeft(binarytrees.NewTree(4))
                root.Left().Left().SetLeft(binarytrees.NewTree(6))
                root.Right().SetRight(binarytrees.NewTree(5))
                root.Right().Right().SetRight(binarytrees.NewTree(7))
                return root
            },
            expected: 6,
        },
        {
            name: "diameter doesn't pass through root",
            buildTree: func() *binarytrees.BinaryTree[int] {
                root := binarytrees.NewTree(1)
                root.SetLeft(binarytrees.NewTree(2))
                root.SetRight(binarytrees.NewTree(3))
                root.Left().SetLeft(binarytrees.NewTree(4))
                root.Left().SetRight(binarytrees.NewTree(5))
                root.Left().Right().SetLeft(binarytrees.NewTree(7))
                root.Left().Right().SetRight(binarytrees.NewTree(8))
                return root
            },
            expected: 4,
        },
        {
            name: "complex tree structure",
            buildTree: func() *binarytrees.BinaryTree[int] {
                /*
                        1
                       / \
                      2   3
                     / \   \
                    4   5   6
                       / \
                      7   8
                     /
                    9
                */
                root := binarytrees.NewTree(1)
                root.SetLeft(binarytrees.NewTree(2))
                root.SetRight(binarytrees.NewTree(3))
                root.Left().SetLeft(binarytrees.NewTree(4))
                root.Left().SetRight(binarytrees.NewTree(5))
                root.Right().SetRight(binarytrees.NewTree(6))
                root.Left().Right().SetLeft(binarytrees.NewTree(7))
                root.Left().Right().SetRight(binarytrees.NewTree(8))
                root.Left().Right().Left().SetLeft(binarytrees.NewTree(9))
                return root
            },
            expected: 6,
        },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            tree := test.buildTree()
            result := binarytrees.Diameter(tree)
            if result != test.expected {
                t.Errorf("Diameter() = %v, want %v", result, test.expected)
            }
        })
    }
}
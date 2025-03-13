package binarytrees_test

import (
	"testing"

	binarytrees "github.com/mundhrakeshav/DSA/golang-dsa/Problems/binary_trees"
)

func TestIsBalanced(t *testing.T) {
    tests := []struct {
        name     string
        buildTree func() *binarytrees.BinaryTree[int]
        expected bool
    }{
        {
            name: "nil tree",
            buildTree: func() *binarytrees.BinaryTree[int] {
                return nil
            },
            expected: true,
        },
        {
            name: "single node",
            buildTree: func() *binarytrees.BinaryTree[int] {
                return binarytrees.NewTree(1)
            },
            expected: true,
        },
        {
            name: "balanced tree - perfect binary tree",
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
            expected: true,
        },
        {
            name: "balanced tree - left side one level deeper",
            buildTree: func() *binarytrees.BinaryTree[int] {
                root := binarytrees.NewTree(1)
                root.SetLeft(binarytrees.NewTree(2))
                root.SetRight(binarytrees.NewTree(3))
                root.Left().SetLeft(binarytrees.NewTree(4))
                return root
            },
            expected: true,
        },
        {
            name: "balanced tree - right side one level deeper",
            buildTree: func() *binarytrees.BinaryTree[int] {
                root := binarytrees.NewTree(1)
                root.SetLeft(binarytrees.NewTree(2))
                root.SetRight(binarytrees.NewTree(3))
                root.Right().SetRight(binarytrees.NewTree(4))
                return root
            },
            expected: true,
        },
        {
            name: "unbalanced tree - left side two levels deeper",
            buildTree: func() *binarytrees.BinaryTree[int] {
                root := binarytrees.NewTree(1)
                root.SetLeft(binarytrees.NewTree(2))
                root.SetRight(binarytrees.NewTree(3))
                root.Left().SetLeft(binarytrees.NewTree(4))
                root.Left().Left().SetLeft(binarytrees.NewTree(5))
                return root
            },
            expected: false,
        },
        {
            name: "unbalanced tree - right side two levels deeper",
            buildTree: func() *binarytrees.BinaryTree[int] {
                root := binarytrees.NewTree(1)
                root.SetLeft(binarytrees.NewTree(2))
                root.SetRight(binarytrees.NewTree(3))
                root.Right().SetRight(binarytrees.NewTree(4))
                root.Right().Right().SetRight(binarytrees.NewTree(5))
                return root
            },
            expected: false,
        },
        {
            name: "balanced tree - complex structure",
            buildTree: func() *binarytrees.BinaryTree[int] {
                root := binarytrees.NewTree(1)
                root.SetLeft(binarytrees.NewTree(2))
                root.SetRight(binarytrees.NewTree(3))
                root.Left().SetLeft(binarytrees.NewTree(4))
                root.Left().SetRight(binarytrees.NewTree(5))
                root.Right().SetLeft(binarytrees.NewTree(6))
                return root
            },
            expected: true,
        },
        {
            name: "unbalanced tree - balanced at root but unbalanced at subtree",
            buildTree: func() *binarytrees.BinaryTree[int] {
                root := binarytrees.NewTree(1)
                root.SetLeft(binarytrees.NewTree(2))
                root.SetRight(binarytrees.NewTree(3))
                root.Left().SetLeft(binarytrees.NewTree(4))
                root.Left().Left().SetLeft(binarytrees.NewTree(5))
                root.Right().SetRight(binarytrees.NewTree(6))
                return root
            },
            expected: false,
        },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            tree := test.buildTree()
            result := binarytrees.IsBalanced(tree)
            if result != test.expected {
                t.Errorf("IsBalanced() = %v, want %v", result, test.expected)
            }
        })
    }
}
package binarysearchtree_test

import (
	"testing"

	binarysearchtree "github.com/mundhrakeshav/DSA/golang-dsa/Problems/binary_search_tree"
	"golang.org/x/exp/constraints"
)

func TestInsertIntoEmptyTree(t *testing.T) {
	bst := binarysearchtree.NewBST[int]()
	bst.InsertRecursive(10)
	if bst.Root == nil {
		t.Fatal("Expected root to be non-nil after insertion")
	}
	if bst.Root.Value != 10 {
		t.Fatalf("Expected root value to be 10, got %v", bst.Root.Value)
	}
}

func TestLeftInsertion(t *testing.T) {
	bst := binarysearchtree.NewBST[int]()
	bst.InsertRecursive(10)
	bst.InsertRecursive(5) // 5 <= 10, should go left
	if bst.Root.Left == nil {
		t.Fatal("Expected left child to be non-nil")
	}
	if bst.Root.Left.Value != 5 {
		t.Fatalf("Expected left child value to be 5, got %v", bst.Root.Left.Value)
	}
}

func TestRightInsertion(t *testing.T) {
	bst := binarysearchtree.NewBST[int]()
	bst.InsertRecursive(10)
	bst.InsertRecursive(15) // 15 > 10, should go right
	if bst.Root.Right == nil {
		t.Fatal("Expected right child to be non-nil")
	}
	if bst.Root.Right.Value != 15 {
		t.Fatalf("Expected right child value to be 15, got %v", bst.Root.Right.Value)
	}
}

func TestDuplicateInsertion(t *testing.T) {
	bst := binarysearchtree.NewBST[int]()
	bst.InsertRecursive(10)
	bst.InsertRecursive(10) // duplicate; according to our logic, goes to left subtree
	if bst.Root.Left == nil {
		t.Fatal("Expected left child for duplicate insertion")
	}
	if bst.Root.Left.Value != 10 {
		t.Fatalf("Expected duplicate value 10 in left child, got %v", bst.Root.Left.Value)
	}
}

func TestMultipleInsertions(t *testing.T) {
	bst := binarysearchtree.NewBST[int]()
	// Insert a mix of values including duplicates
	values := []int{10, 5, 15, 3, 7, 12, 18, 5}
	for _, v := range values {
		bst.InsertRecursive(v)
	}

	// inOrder should yield a sorted slice.
	result := inOrder(bst.Root)
	expected := []int{3, 5, 5, 7, 10, 12, 15, 18}
	if len(result) != len(expected) {
		t.Fatalf("Expected %d nodes in in-order traversal, got %d", len(expected), len(result))
	}
	for i, v := range expected {
		if result[i] != v {
			t.Fatalf("Expected value %d at index %d, got %d", v, i, result[i])
		}
	}
}

func TestNegativeNumbers(t *testing.T) {
	bst := binarysearchtree.NewBST[int]()
	// Test with negative numbers
	values := []int{-10, -20, -5, -15}
	for _, v := range values {
		bst.InsertRecursive(v)
	}
	result := inOrder(bst.Root)
	expected := []int{-20, -15, -10, -5}
	if len(result) != len(expected) {
		t.Fatalf("Expected %d nodes in in-order traversal, got %d", len(expected), len(result))
	}
	for i, v := range expected {
		if result[i] != v {
			t.Fatalf("Expected value %d at index %d, got %d", v, i, result[i])
		}
	}
}

func TestStringBST(t *testing.T) {
	bst := binarysearchtree.NewBST[string]()
	values := []string{"m", "a", "z", "b"}
	for _, v := range values {
		bst.InsertRecursive(v)
	}
	result := inOrder(bst.Root)
	expected := []string{"a", "b", "m", "z"}
	if len(result) != len(expected) {
		t.Fatalf("Expected %d nodes in in-order traversal, got %d", len(expected), len(result))
	}
	for i, v := range expected {
		if result[i] != v {
			t.Fatalf("Expected value %s at index %d, got %s", v, i, result[i])
		}
	}
}

// inOrder returns the in-order traversal of the BST as a slice.
// This helper works for any type T that satisfies constraints.Ordered.
func inOrder[T constraints.Ordered](node *binarysearchtree.Node[T]) []T {
	var result []T
	if node == nil {
		return result
	}
	result = append(result, inOrder(node.Left)...)
	result = append(result, node.Value)
	result = append(result, inOrder(node.Right)...)
	return result
}

package binarysearchtree_test

import (
	"reflect"
	"testing"

	binarysearchtree "github.com/mundhrakeshav/DSA/golang-dsa/Problems/binary_search_tree"
)

func TestInsertRecursive(t *testing.T) {
	t.Run("Insert into empty tree", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		bst.InsertRecursive(10)
		if bst.Root == nil {
			t.Fatal("Expected root to be non-nil after insertion")
		}
		if bst.Root.Value != 10 {
			t.Fatalf("Expected root value to be 10, got %v", bst.Root.Value)
		}
	})

	t.Run("Left insertion", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		bst.InsertRecursive(10)
		bst.InsertRecursive(5) // 5 <= 10, should go left
		if bst.Root.Left == nil {
			t.Fatal("Expected left child to be non-nil")
		}
		if bst.Root.Left.Value != 5 {
			t.Fatalf("Expected left child value to be 5, got %v", bst.Root.Left.Value)
		}
	})

	t.Run("Right insertion", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		bst.InsertRecursive(10)
		bst.InsertRecursive(15) // 15 > 10, should go right
		if bst.Root.Right == nil {
			t.Fatal("Expected right child to be non-nil")
		}
		if bst.Root.Right.Value != 15 {
			t.Fatalf("Expected right child value to be 15, got %v", bst.Root.Right.Value)
		}
	})

	t.Run("Duplicate insertion", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		bst.InsertRecursive(10)
		bst.InsertRecursive(10) // duplicate; according to our logic, goes to left subtree
		if bst.Root.Left == nil {
			t.Fatal("Expected left child for duplicate insertion")
		}
		if bst.Root.Left.Value != 10 {
			t.Fatalf("Expected duplicate value 10 in left child, got %v", bst.Root.Left.Value)
		}
	})

	t.Run("Multiple insertions", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		// Insert a mix of values including duplicates
		values := []int{10, 5, 15, 3, 7, 12, 18, 5}
		for _, v := range values {
			bst.InsertRecursive(v)
		}

		// inOrder should yield a sorted slice.
		result := bst.InOrder()
		expected := []int{3, 5, 5, 7, 10, 12, 15, 18}
		if !reflect.DeepEqual(result, expected) {
			t.Fatalf("Expected in-order traversal %v, got %v", expected, result)
		}
	})

	t.Run("Negative numbers", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		// Test with negative numbers
		values := []int{-10, -20, -5, -15}
		for _, v := range values {
			bst.InsertRecursive(v)
		}
		result := bst.InOrder()
		expected := []int{-20, -15, -10, -5}
		if !reflect.DeepEqual(result, expected) {
			t.Fatalf("Expected in-order traversal %v, got %v", expected, result)
		}
	})

	t.Run("String BST", func(t *testing.T) {
		bst := binarysearchtree.NewBST[string]()
		values := []string{"m", "a", "z", "b"}
		for _, v := range values {
			bst.InsertRecursive(v)
		}
		result := bst.InOrder()
		expected := []string{"a", "b", "m", "z"}
		if !reflect.DeepEqual(result, expected) {
			t.Fatalf("Expected in-order traversal %v, got %v", expected, result)
		}
	})
}

func TestInOrder(t *testing.T) {
	t.Run("Empty tree", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		got := bst.InOrder()
		want := []int{}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("InOrder() = %v, want %v", got, want)
		}
	})

	t.Run("Single node tree", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		bst.InsertRecursive(5)
		got := bst.InOrder()
		want := []int{5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("InOrder() = %v, want %v", got, want)
		}
	})

	t.Run("Balanced tree", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		values := []int{5, 3, 7, 2, 4, 6, 8}
		for _, v := range values {
			bst.InsertRecursive(v)
		}
		got := bst.InOrder()
		want := []int{2, 3, 4, 5, 6, 7, 8}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("InOrder() = %v, want %v", got, want)
		}
	})

	t.Run("Right heavy tree", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		values := []int{1, 2, 3, 4, 5}
		for _, v := range values {
			bst.InsertRecursive(v)
		}
		got := bst.InOrder()
		want := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("InOrder() = %v, want %v", got, want)
		}
	})

	t.Run("Left heavy tree", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		values := []int{5, 4, 3, 2, 1}
		for _, v := range values {
			bst.InsertRecursive(v)
		}
		got := bst.InOrder()
		want := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("InOrder() = %v, want %v", got, want)
		}
	})

	t.Run("Tree with duplicates", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		values := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
		for _, v := range values {
			bst.InsertRecursive(v)
		}
		got := bst.InOrder()
		want := []int{1, 1, 2, 3, 4, 5, 5, 6, 9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("InOrder() = %v, want %v", got, want)
		}
	})

	t.Run("String type tree", func(t *testing.T) {
		bst := binarysearchtree.NewBST[string]()
		values := []string{"banana", "apple", "cherry", "date", "blueberry"}
		for _, v := range values {
			bst.InsertRecursive(v)
		}
		got := bst.InOrder()
		want := []string{"apple", "banana", "blueberry", "cherry", "date"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("InOrder() = %v, want %v", got, want)
		}
	})
}

func TestInOrderNodes(t *testing.T) {
	t.Run("Empty tree", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		got := bst.InOrderNodes()
		if got == nil {
			t.Errorf("InOrderNodes() = nil, want empty slice")
		}
		if len(got) != 0 {
			t.Errorf("InOrderNodes() has length %d, want 0", len(got))
		}
	})

	t.Run("Single node tree", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		bst.InsertRecursive(5)
		got := bst.InOrderNodes()
		if len(got) != 1 {
			t.Errorf("InOrderNodes() has length %d, want 1", len(got))
		}
		if got[0].Value != 5 {
			t.Errorf("InOrderNodes()[0].Value = %v, want 5", got[0].Value)
		}
	})

	t.Run("Balanced tree", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		values := []int{5, 3, 7, 2, 4, 6, 8}
		for _, v := range values {
			bst.InsertRecursive(v)
		}
		got := bst.InOrderNodes()
		if len(got) != len(values) {
			t.Errorf("InOrderNodes() has length %d, want %d", len(got), len(values))
		}

		// Check if the values are in the correct order
		expectedValues := []int{2, 3, 4, 5, 6, 7, 8}
		for i, node := range got {
			if node.Value != expectedValues[i] {
				t.Errorf("InOrderNodes()[%d].Value = %v, want %v", i, node.Value, expectedValues[i])
			}
		}
	})

	t.Run("Right heavy tree", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		values := []int{1, 2, 3, 4, 5}
		for _, v := range values {
			bst.InsertRecursive(v)
		}
		got := bst.InOrderNodes()
		if len(got) != len(values) {
			t.Errorf("InOrderNodes() has length %d, want %d", len(got), len(values))
		}

		// Check if the values are in the correct order
		expectedValues := []int{1, 2, 3, 4, 5}
		for i, node := range got {
			if node.Value != expectedValues[i] {
				t.Errorf("InOrderNodes()[%d].Value = %v, want %v", i, node.Value, expectedValues[i])
			}
		}
	})

	t.Run("Left heavy tree", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		values := []int{5, 4, 3, 2, 1}
		for _, v := range values {
			bst.InsertRecursive(v)
		}
		got := bst.InOrderNodes()
		if len(got) != len(values) {
			t.Errorf("InOrderNodes() has length %d, want %d", len(got), len(values))
		}

		// Check if the values are in the correct order
		expectedValues := []int{1, 2, 3, 4, 5}
		for i, node := range got {
			if node.Value != expectedValues[i] {
				t.Errorf("InOrderNodes()[%d].Value = %v, want %v", i, node.Value, expectedValues[i])
			}
		}
	})

	t.Run("Tree with duplicates", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		values := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
		for _, v := range values {
			bst.InsertRecursive(v)
		}
		got := bst.InOrderNodes()
		if len(got) != len(values) {
			t.Errorf("InOrderNodes() has length %d, want %d", len(got), len(values))
		}

		// Check if the values are in the correct order
		expectedValues := []int{1, 1, 2, 3, 4, 5, 5, 6, 9}
		for i, node := range got {
			if node.Value != expectedValues[i] {
				t.Errorf("InOrderNodes()[%d].Value = %v, want %v", i, node.Value, expectedValues[i])
			}
		}
	})

	t.Run("String type tree", func(t *testing.T) {
		bst := binarysearchtree.NewBST[string]()
		values := []string{"banana", "apple", "cherry", "date", "blueberry"}
		for _, v := range values {
			bst.InsertRecursive(v)
		}
		got := bst.InOrderNodes()
		if len(got) != len(values) {
			t.Errorf("InOrderNodes() has length %d, want %d", len(got), len(values))
		}

		// Check if the values are in the correct order
		expectedValues := []string{"apple", "banana", "blueberry", "cherry", "date"}
		for i, node := range got {
			if node.Value != expectedValues[i] {
				t.Errorf("InOrderNodes()[%d].Value = %v, want %v", i, node.Value, expectedValues[i])
			}
		}
	})

	t.Run("Verify actual node references", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		bst.InsertRecursive(5)
		bst.InsertRecursive(3)
		bst.InsertRecursive(7)

		nodes := bst.InOrderNodes()
		if len(nodes) != 3 {
			t.Fatalf("Expected 3 nodes, got %d", len(nodes))
		}

		// Check that we're getting the actual nodes, not copies
		if nodes[0] != bst.Root.Left {
			t.Errorf("First node should be the left child of root")
		}

		if nodes[1] != bst.Root {
			t.Errorf("Second node should be the root")
		}

		if nodes[2] != bst.Root.Right {
			t.Errorf("Third node should be the right child of root")
		}
	})
}

func TestFindMaxNode(t *testing.T) {
	t.Run("Empty tree", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		got := bst.FindMaxNode()
		if got != nil {
			t.Errorf("FindMaxNode() = %v, want nil", got)
		}
	})

	t.Run("Single node tree", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		bst.InsertRecursive(5)
		got := bst.FindMaxNode()
		if got == nil {
			t.Fatal("FindMaxNode() = nil, want non-nil")
		}
		if got.Value != 5 {
			t.Errorf("FindMaxNode().Value = %v, want 5", got.Value)
		}
	})

	t.Run("Balanced tree", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		values := []int{5, 3, 7, 2, 4, 6, 8}
		for _, v := range values {
			bst.InsertRecursive(v)
		}
		got := bst.FindMaxNode()
		if got == nil {
			t.Fatal("FindMaxNode() = nil, want non-nil")
		}
		if got.Value != 8 {
			t.Errorf("FindMaxNode().Value = %v, want 8", got.Value)
		}
	})

	t.Run("Right heavy tree", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		values := []int{1, 2, 3, 4, 5}
		for _, v := range values {
			bst.InsertRecursive(v)
		}
		got := bst.FindMaxNode()
		if got == nil {
			t.Fatal("FindMaxNode() = nil, want non-nil")
		}
		if got.Value != 5 {
			t.Errorf("FindMaxNode().Value = %v, want 5", got.Value)
		}
	})

	t.Run("Left heavy tree", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		values := []int{5, 4, 3, 2, 1}
		for _, v := range values {
			bst.InsertRecursive(v)
		}
		got := bst.FindMaxNode()
		if got == nil {
			t.Fatal("FindMaxNode() = nil, want non-nil")
		}
		if got.Value != 5 {
			t.Errorf("FindMaxNode().Value = %v, want 5", got.Value)
		}
	})

	t.Run("String type tree", func(t *testing.T) {
		bst := binarysearchtree.NewBST[string]()
		values := []string{"banana", "apple", "cherry", "date", "blueberry"}
		for _, v := range values {
			bst.InsertRecursive(v)
		}
		got := bst.FindMaxNode()
		if got == nil {
			t.Fatal("FindMaxNode() = nil, want non-nil")
		}
		if got.Value != "date" {
			t.Errorf("FindMaxNode().Value = %v, want 'date'", got.Value)
		}
	})

	t.Run("Verify actual node reference", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		bst.InsertRecursive(5)
		bst.InsertRecursive(3)
		bst.InsertRecursive(7)
		bst.InsertRecursive(9)

		rightMost := bst.FindMaxNode()
		if rightMost == nil {
			t.Fatal("FindMaxNode() = nil, want non-nil")
		}

		// Verify it's the actual rightmost node
		nodes := bst.InOrderNodes()
		lastNode := nodes[len(nodes)-1]
		if rightMost != lastNode {
			t.Errorf("FindMaxNode() returned node with value %v, but it should be the last node with value %v",
				rightMost.Value, lastNode.Value)
		}
	})
}

func TestFind(t *testing.T) {
	t.Run("Empty tree", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		got := bst.Find(5)
		if got != nil {
			t.Errorf("Find(5) = %v, want nil", got)
		}
	})

	t.Run("Single node tree - value exists", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		bst.InsertRecursive(5)
		got := bst.Find(5)
		if got == nil {
			t.Fatal("Find(5) = nil, want non-nil")
		}
		if got.Value != 5 {
			t.Errorf("Find(5).Value = %v, want 5", got.Value)
		}
	})

	t.Run("Single node tree - value doesn't exist", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		bst.InsertRecursive(5)
		got := bst.Find(10)
		if got != nil {
			t.Errorf("Find(10) = %v, want nil", got)
		}
	})

	t.Run("Balanced tree - value in root", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		values := []int{5, 3, 7, 2, 4, 6, 8}
		for _, v := range values {
			bst.InsertRecursive(v)
		}
		got := bst.Find(5)
		if got == nil {
			t.Fatal("Find(5) = nil, want non-nil")
		}
		if got.Value != 5 {
			t.Errorf("Find(5).Value = %v, want 5", got.Value)
		}
		if got != bst.Root {
			t.Errorf("Find(5) should return the root node")
		}
	})

	t.Run("Balanced tree - value in left subtree", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		values := []int{5, 3, 7, 2, 4, 6, 8}
		for _, v := range values {
			bst.InsertRecursive(v)
		}
		got := bst.Find(2)
		if got == nil {
			t.Fatal("Find(2) = nil, want non-nil")
		}
		if got.Value != 2 {
			t.Errorf("Find(2).Value = %v, want 2", got.Value)
		}
	})

	t.Run("Balanced tree - value in right subtree", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		values := []int{5, 3, 7, 2, 4, 6, 8}
		for _, v := range values {
			bst.InsertRecursive(v)
		}
		got := bst.Find(8)
		if got == nil {
			t.Fatal("Find(8) = nil, want non-nil")
		}
		if got.Value != 8 {
			t.Errorf("Find(8).Value = %v, want 8", got.Value)
		}
	})

	t.Run("Balanced tree - value doesn't exist", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		values := []int{5, 3, 7, 2, 4, 6, 8}
		for _, v := range values {
			bst.InsertRecursive(v)
		}
		got := bst.Find(10)
		if got != nil {
			t.Errorf("Find(10) = %v, want nil", got)
		}
	})

	t.Run("Duplicate values - finds first occurrence", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		values := []int{5, 3, 7, 5, 4} // 5 appears twice
		for _, v := range values {
			bst.InsertRecursive(v)
		}
		got := bst.Find(5)
		if got == nil {
			t.Fatal("Find(5) = nil, want non-nil")
		}
		if got.Value != 5 {
			t.Errorf("Find(5).Value = %v, want 5", got.Value)
		}
		// Should find the root node with value 5, not the duplicate
		if got != bst.Root {
			t.Errorf("Find(5) should return the root node, not the duplicate")
		}
	})

	t.Run("String type tree", func(t *testing.T) {
		bst := binarysearchtree.NewBST[string]()
		values := []string{"banana", "apple", "cherry", "date", "blueberry"}
		for _, v := range values {
			bst.InsertRecursive(v)
		}
		got := bst.Find("cherry")
		if got == nil {
			t.Fatal("Find(\"cherry\") = nil, want non-nil")
		}
		if got.Value != "cherry" {
			t.Errorf("Find(\"cherry\").Value = %v, want \"cherry\"", got.Value)
		}

		// Test non-existent value
		got = bst.Find("grape")
		if got != nil {
			t.Errorf("Find(\"grape\") = %v, want nil", got)
		}
	})
}
func TestDelete(t *testing.T) {
	t.Run("Empty tree", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		bst.Delete(5) // Should do nothing
		if bst.Root != nil {
			t.Errorf("Delete(5) on empty tree should keep root nil")
		}
	})

	t.Run("Delete leaf node", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		values := []int{10, 5, 15, 3, 7}
		for _, v := range values {
			bst.InsertRecursive(v)
		}

		bst.Delete(3) // Leaf node

		// Check if the node was deleted
		result := bst.InOrder()
		expected := []int{5, 7, 10, 15}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("After Delete(3), InOrder() = %v, want %v", result, expected)
		}

		// Verify the node is actually gone
		if bst.Find(3) != nil {
			t.Errorf("After Delete(3), Find(3) should return nil")
		}
	})

	t.Run("Delete node with one child", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		values := []int{10, 5, 15, 3}
		for _, v := range values {
			bst.InsertRecursive(v)
		}

		bst.Delete(5) // Node with one child (left child 3)

		// Check if the node was deleted and structure is maintained
		result := bst.InOrder()
		expected := []int{3, 10, 15}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("After Delete(5), InOrder() = %v, want %v", result, expected)
		}

		// Root's left child should now be 3
		if bst.Root.Left.Value != 3 {
			t.Errorf("After Delete(5), Root.Left.Value = %v, want 3", bst.Root.Left.Value)
		}
	})

	t.Run("Delete node with two children", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		values := []int{10, 5, 15, 3, 7, 12, 17}
		for _, v := range values {
			bst.InsertRecursive(v)
		}

		bst.Delete(15) // Node with two children (12 and 17)

		// Check if the node was deleted and structure is maintained
		result := bst.InOrder()
		expected := []int{3, 5, 7, 10, 12, 17}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("After Delete(15), InOrder() = %v, want %v", result, expected)
		}

		// Root's right child should now be 17
		if bst.Root.Right.Value != 17 {
			t.Errorf("After Delete(15), Root.Right.Value = %v, want 17", bst.Root.Right.Value)
		}

		// 17's left child should now be 12
		if bst.Root.Right.Left.Value != 12 {
			t.Errorf("After Delete(15), Root.Right.Left.Value = %v, want 12", bst.Root.Right.Left.Value)
		}
	})

	t.Run("Delete root node", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		values := []int{10, 5, 15, 3, 7, 12, 17}
		for _, v := range values {
			bst.InsertRecursive(v)
		}

		bst.Delete(10) // Root node

		// Check if the node was deleted and structure is maintained
		result := bst.InOrder()
		expected := []int{3, 5, 7, 12, 15, 17}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("After Delete(10), InOrder() = %v, want %v", result, expected)
		}

		// Root should now be 12
		if bst.Root.Value != 12 {
			t.Errorf("After Delete(10), Root.Value = %v, want 12", bst.Root.Value)
		}
	})

	t.Run("Delete non-existent value", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		values := []int{10, 5, 15}
		for _, v := range values {
			bst.InsertRecursive(v)
		}

		// Should not change the tree
		bst.Delete(20)

		result := bst.InOrder()
		expected := []int{5, 10, 15}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("After Delete(20), InOrder() = %v, want %v", result, expected)
		}
	})

	t.Run("Delete from single node tree", func(t *testing.T) {
		bst := binarysearchtree.NewBST[int]()
		bst.InsertRecursive(10)

		bst.Delete(10)

		if bst.Root != nil {
			t.Errorf("After Delete(10) on single node tree, Root should be nil")
		}
	})

	t.Run("Delete with string values", func(t *testing.T) {
		bst := binarysearchtree.NewBST[string]()
		values := []string{"banana", "apple", "cherry", "date"}
		for _, v := range values {
			bst.InsertRecursive(v)
		}

		bst.Delete("cherry")

		result := bst.InOrder()
		expected := []string{"apple", "banana", "date"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("After Delete(\"cherry\"), InOrder() = %v, want %v", result, expected)
		}
	})
}

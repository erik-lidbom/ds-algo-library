package trees

import (
	"testing"
)

func TestNewBinarySearchTree(t *testing.T) {
	bst := NewBinarySearchTree[int]()
	if bst == nil {
		t.Fatal("NewBinarySearchTree returned nil")
	}
	if bst.Size() != 0 {
		t.Errorf("Expected size of 0, got %d", bst.Size())
	}
	if !bst.IsEmpty() {
		t.Error("Expected empty BinarySearchTree")
	}
}

func TestBST_Size(t *testing.T) {
	bst := NewBinarySearchTree[int]()

	if bst.Size() != 0 {
		t.Errorf("Expected size of 0, got %d", bst.Size())
	}

	bst.Insert(1)

	if bst.Size() != 1 {
		t.Errorf("Expected size 1, got %d", bst.Size())
	}

	bst.Insert(2)

	if bst.Size() != 2 {
		t.Errorf("Expected size 2, got %d", bst.Size())
	}

	bst.Delete(2)

	if bst.Size() != 1 {
		t.Errorf("Expected size 1, got %d", bst.Size())
	}
}

func TestBST_IsEmpty(t *testing.T) {
	bst := NewBinarySearchTree[int]()

	if !bst.IsEmpty() {
		t.Error("New BST should be empty")
	}

	bst.Insert(1)

	if bst.IsEmpty() {
		t.Error("BST with elements should not be empty")
	}

	bst.Delete(1)

	if !bst.IsEmpty() {
		t.Error("BST should be empty after deleting all elements")
	}
}

func TestBST_Insert(t *testing.T) {
	bst := NewBinarySearchTree[int]()

	// Test inserting single element
	err := bst.Insert(5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if bst.Size() != 1 {
		t.Errorf("Expected size 1, got %d", bst.Size())
	}

	// Test inserting multiple elements
	elements := []int{3, 7, 1, 9, 2, 8}
	for _, elem := range elements {
		err = bst.Insert(elem)
		if err != nil {
			t.Errorf("Unexpected error inserting %d: %v", elem, err)
		}
	}

	// Test inserting duplicate (should not increase size)
	initialSize := bst.Size()
	err = bst.Insert(5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if bst.Size() != initialSize {
		t.Errorf("Expected size %d after inserting duplicate, got %d", initialSize, bst.Size())
	}

	// Test inserting negative numbers
	err = bst.Insert(-5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test inserting zero
	err = bst.Insert(0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestBST_Search(t *testing.T) {
	bst := NewBinarySearchTree[int]()

	// Test searching empty tree
	value, found := bst.Search(5)
	if found {
		t.Error("Should not find element in empty tree")
	}
	if value != 0 {
		t.Errorf("Expected zero value, got %d", value)
	}

	// Test searching existing element
	bst.Insert(5)
	value, found = bst.Search(5)
	if !found {
		t.Error("Should find existing element")
	}
	if value != 5 {
		t.Errorf("Expected 5, got %d", value)
	}

	// Test searching non-existing element
	value, found = bst.Search(10)
	if found {
		t.Error("Should not find non-existing element")
	}
	if value != 0 {
		t.Errorf("Expected zero value, got %d", value)
	}

	// Test searching after multiple insertions
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(1)
	bst.Insert(9)

	testCases := []struct {
		searchValue int
		shouldFind  bool
		expected    int
	}{
		{1, true, 1},
		{3, true, 3},
		{5, true, 5},
		{7, true, 7},
		{9, true, 9},
		{2, false, 0},
		{4, false, 0},
		{6, false, 0},
		{8, false, 0},
		{10, false, 0},
	}

	for _, tc := range testCases {
		value, found := bst.Search(tc.searchValue)
		if found != tc.shouldFind {
			t.Errorf("Search for %d: expected found=%v, got %v", tc.searchValue, tc.shouldFind, found)
		}
		if found && value != tc.expected {
			t.Errorf("Search for %d: expected %d, got %d", tc.searchValue, tc.expected, value)
		}
	}
}

func TestBST_Delete(t *testing.T) {
	bst := NewBinarySearchTree[int]()

	// Test deleting from empty tree
	err := bst.Delete(5)
	if err == nil {
		t.Error("Expected error when deleting from empty tree")
	}

	// Test deleting existing element
	bst.Insert(5)
	err = bst.Delete(5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if bst.Size() != 0 {
		t.Errorf("Expected size 0 after deletion, got %d", bst.Size())
	}

	// Test deleting non-existing element
	err = bst.Delete(5)
	if err == nil {
		t.Error("Expected error when deleting non-existing element")
	}

	// Test deleting leaf node
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)
	err = bst.Delete(3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if bst.Size() != 2 {
		t.Errorf("Expected size 2, got %d", bst.Size())
	}

	// Test deleting node with one child
	bst.Insert(1)
	err = bst.Delete(5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test deleting node with two children
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)
	err = bst.Delete(5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestBST_Clear(t *testing.T) {
	bst := NewBinarySearchTree[int]()

	// Test clearing empty tree
	err := bst.Clear()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if bst.Size() != 0 {
		t.Errorf("Expected size 0, got %d", bst.Size())
	}

	// Test clearing tree with elements
	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(3)
	err = bst.Clear()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if bst.Size() != 0 {
		t.Errorf("Expected size 0, got %d", bst.Size())
	}
	if !bst.IsEmpty() {
		t.Error("Tree should be empty after clearing")
	}
}

func TestBST_TraverseInOrder(t *testing.T) {
	bst := NewBinarySearchTree[int]()

	// Test traversing empty tree
	result, err := bst.TraverseInOrder(bst.root)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result.Size() != 0 {
		t.Errorf("Expected empty result, got size %d", result.Size())
	}

	// Test traversing tree with elements
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(1)
	bst.Insert(9)

	result, err = bst.TraverseInOrder(bst.root)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := []int{1, 3, 5, 7, 9}
	if result.Size() != len(expected) {
		t.Errorf("Expected size %d, got %d", len(expected), result.Size())
	}

	for i, expectedValue := range expected {
		actualValue, err := result.Get(i)
		if err != nil {
			t.Errorf("Error getting value at index %d: %v", i, err)
		}
		if actualValue != expectedValue {
			t.Errorf("Expected %d at index %d, got %d", expectedValue, i, actualValue)
		}
	}
}

func TestBST_TraversePreOrder(t *testing.T) {
	bst := NewBinarySearchTree[int]()

	// Test traversing empty tree
	result, err := bst.TraversePreOrder(bst.root)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result.Size() != 0 {
		t.Errorf("Expected empty result, got size %d", result.Size())
	}

	// Test traversing tree with elements
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(1)
	bst.Insert(9)

	result, err = bst.TraversePreOrder(bst.root)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := []int{5, 3, 1, 7, 9}
	if result.Size() != len(expected) {
		t.Errorf("Expected size %d, got %d", len(expected), result.Size())
	}

	for i, expectedValue := range expected {
		actualValue, err := result.Get(i)
		if err != nil {
			t.Errorf("Error getting value at index %d: %v", i, err)
		}
		if actualValue != expectedValue {
			t.Errorf("Expected %d at index %d, got %d", expectedValue, i, actualValue)
		}
	}
}

func TestBST_TraversePostOrder(t *testing.T) {
	bst := NewBinarySearchTree[int]()

	// Test traversing empty tree
	result, err := bst.TraversePostOrder(bst.root)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result.Size() != 0 {
		t.Errorf("Expected empty result, got size %d", result.Size())
	}

	// Test traversing tree with elements
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(1)
	bst.Insert(9)

	result, err = bst.TraversePostOrder(bst.root)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := []int{1, 3, 9, 7, 5}
	if result.Size() != len(expected) {
		t.Errorf("Expected size %d, got %d", len(expected), result.Size())
	}

	for i, expectedValue := range expected {
		actualValue, err := result.Get(i)
		if err != nil {
			t.Errorf("Error getting value at index %d: %v", i, err)
		}
		if actualValue != expectedValue {
			t.Errorf("Expected %d at index %d, got %d", expectedValue, i, actualValue)
		}
	}
}

func TestBST_ComplexOperations(t *testing.T) {
	bst := NewBinarySearchTree[int]()

	// Test complex sequence of operations
	elements := []int{50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45, 55, 65, 75, 85}
	
	// Insert all elements
	for _, elem := range elements {
		err := bst.Insert(elem)
		if err != nil {
			t.Errorf("Unexpected error inserting %d: %v", elem, err)
		}
	}

	if bst.Size() != len(elements) {
		t.Errorf("Expected size %d, got %d", len(elements), bst.Size())
	}

	// Verify all elements exist
	for _, elem := range elements {
		value, found := bst.Search(elem)
		if !found {
			t.Errorf("Element %d should exist", elem)
		}
		if value != elem {
			t.Errorf("Expected %d, got %d", elem, value)
		}
	}

	// Test deleting leaf nodes
	err := bst.Delete(10)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if bst.Size() != len(elements)-1 {
		t.Errorf("Expected size %d, got %d", len(elements)-1, bst.Size())
	}

	// Test deleting node with one child
	err = bst.Delete(20)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test deleting node with two children
	err = bst.Delete(50)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test deleting root
	err = bst.Delete(30)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test clearing tree
	err = bst.Clear()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if bst.Size() != 0 {
		t.Errorf("Expected size 0 after clearing, got %d", bst.Size())
	}
}

func TestBST_EdgeCases(t *testing.T) {
	bst := NewBinarySearchTree[int]()

	// Test with negative numbers
	bst.Insert(-5)
	bst.Insert(-10)
	bst.Insert(-3)

	value, found := bst.Search(-5)
	if !found {
		t.Error("Should find negative number")
	}
	if value != -5 {
		t.Errorf("Expected -5, got %d", value)
	}

	// Test with zero
	bst.Insert(0)
	value, found = bst.Search(0)
	if !found {
		t.Error("Should find zero")
	}
	if value != 0 {
		t.Errorf("Expected 0, got %d", value)
	}

	// Test with large numbers
	bst.Insert(1000)
	bst.Insert(999)
	bst.Insert(1001)

	value, found = bst.Search(1000)
	if !found {
		t.Error("Should find large number")
	}
	if value != 1000 {
		t.Errorf("Expected 1000, got %d", value)
	}

	// Test with duplicate insertions
	initialSize := bst.Size()
	bst.Insert(0) // Duplicate
	if bst.Size() != initialSize {
		t.Errorf("Expected size %d after duplicate insertion, got %d", initialSize, bst.Size())
	}
}

func TestBST_TraversalWithNilNode(t *testing.T) {
	bst := NewBinarySearchTree[int]()

	// Test traversals with nil node
	result, err := bst.TraverseInOrder(nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result.Size() != 0 {
		t.Errorf("Expected empty result for nil node, got size %d", result.Size())
	}

	result, err = bst.TraversePreOrder(nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result.Size() != 0 {
		t.Errorf("Expected empty result for nil node, got size %d", result.Size())
	}

	result, err = bst.TraversePostOrder(nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result.Size() != 0 {
		t.Errorf("Expected empty result for nil node, got size %d", result.Size())
	}
} 
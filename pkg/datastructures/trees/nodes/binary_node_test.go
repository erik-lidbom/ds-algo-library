package nodes

import (
	"testing"
)

func TestBinaryNode_New(t *testing.T) {
	node := &BinaryNode[int]{
		Value: 5,
		Left:  nil,
		Right: nil,
	}

	if node.Value != 5 {
		t.Errorf("Expected value 5, got %d", node.Value)
	}
	if node.Left != nil {
		t.Error("Expected left child to be nil")
	}
	if node.Right != nil {
		t.Error("Expected right child to be nil")
	}
}

func TestBinaryNode_GetLeft(t *testing.T) {
	// Test with nil node
	var nilNode *BinaryNode[int]
	left := nilNode.GetLeft()
	if left != nil {
		t.Error("Expected nil for nil node")
	}

	// Test with node that has nil left child
	node := &BinaryNode[int]{Value: 5, Left: nil, Right: nil}
	left = node.GetLeft()
	if left != nil {
		t.Error("Expected nil for node with nil left child")
	}

	// Test with node that has left child
	leftChild := &BinaryNode[int]{Value: 3, Left: nil, Right: nil}
	node.Left = leftChild
	left = node.GetLeft()
	if left == nil {
		t.Error("Expected non-nil left child")
	}
	if left.GetValue() != 3 {
		t.Errorf("Expected left child value 3, got %d", left.GetValue())
	}
}

func TestBinaryNode_GetRight(t *testing.T) {
	// Test with nil node
	var nilNode *BinaryNode[int]
	right := nilNode.GetRight()
	if right != nil {
		t.Error("Expected nil for nil node")
	}

	// Test with node that has nil right child
	node := &BinaryNode[int]{Value: 5, Left: nil, Right: nil}
	right = node.GetRight()
	if right != nil {
		t.Error("Expected nil for node with nil right child")
	}

	// Test with node that has right child
	rightChild := &BinaryNode[int]{Value: 7, Left: nil, Right: nil}
	node.Right = rightChild
	right = node.GetRight()
	if right == nil {
		t.Error("Expected non-nil right child")
	}
	if right.GetValue() != 7 {
		t.Errorf("Expected right child value 7, got %d", right.GetValue())
	}
}

func TestBinaryNode_GetValue(t *testing.T) {
	// Test with nil node
	var nilNode *BinaryNode[int]
	value := nilNode.GetValue()
	if value != 0 {
		t.Errorf("Expected zero value for nil node, got %d", value)
	}

	// Test with valid node
	node := &BinaryNode[int]{Value: 5, Left: nil, Right: nil}
	value = node.GetValue()
	if value != 5 {
		t.Errorf("Expected value 5, got %d", value)
	}

	// Test with negative value
	node = &BinaryNode[int]{Value: -5, Left: nil, Right: nil}
	value = node.GetValue()
	if value != -5 {
		t.Errorf("Expected value -5, got %d", value)
	}
}

func TestBinaryNode_IsNil(t *testing.T) {
	// Test with nil node
	var nilNode *BinaryNode[int]
	if !nilNode.IsNil() {
		t.Error("Expected nil node to return true for IsNil")
	}

	// Test with valid node
	node := &BinaryNode[int]{Value: 5, Left: nil, Right: nil}
	if node.IsNil() {
		t.Error("Expected non-nil node to return false for IsNil")
	}
}

func TestBinaryNode_ComplexStructure(t *testing.T) {
	// Create a complex tree structure
	root := &BinaryNode[int]{Value: 5, Left: nil, Right: nil}
	left := &BinaryNode[int]{Value: 3, Left: nil, Right: nil}
	right := &BinaryNode[int]{Value: 7, Left: nil, Right: nil}
	leftLeft := &BinaryNode[int]{Value: 1, Left: nil, Right: nil}
	leftRight := &BinaryNode[int]{Value: 4, Left: nil, Right: nil}
	rightLeft := &BinaryNode[int]{Value: 6, Left: nil, Right: nil}
	rightRight := &BinaryNode[int]{Value: 9, Left: nil, Right: nil}

	// Build the tree
	root.Left = left
	root.Right = right
	left.Left = leftLeft
	left.Right = leftRight
	right.Left = rightLeft
	right.Right = rightRight

	// Test root
	if root.GetValue() != 5 {
		t.Errorf("Expected root value 5, got %d", root.GetValue())
	}
	if root.IsNil() {
		t.Error("Root should not be nil")
	}

	// Test left subtree
	leftChild := root.GetLeft()
	if leftChild == nil {
		t.Error("Expected non-nil left child")
	}
	if leftChild.GetValue() != 3 {
		t.Errorf("Expected left child value 3, got %d", leftChild.GetValue())
	}

	// Test right subtree
	rightChild := root.GetRight()
	if rightChild == nil {
		t.Error("Expected non-nil right child")
	}
	if rightChild.GetValue() != 7 {
		t.Errorf("Expected right child value 7, got %d", rightChild.GetValue())
	}

	// Test deeper levels
	leftLeftChild := leftChild.GetLeft()
	if leftLeftChild == nil {
		t.Error("Expected non-nil left-left child")
	}
	if leftLeftChild.GetValue() != 1 {
		t.Errorf("Expected left-left child value 1, got %d", leftLeftChild.GetValue())
	}

	rightRightChild := rightChild.GetRight()
	if rightRightChild == nil {
		t.Error("Expected non-nil right-right child")
	}
	if rightRightChild.GetValue() != 9 {
		t.Errorf("Expected right-right child value 9, got %d", rightRightChild.GetValue())
	}
}

func TestBinaryNode_StringValues(t *testing.T) {
	// Test with string values
	node := &BinaryNode[string]{Value: "test", Left: nil, Right: nil}
	value := node.GetValue()
	if value != "test" {
		t.Errorf("Expected value 'test', got '%s'", value)
	}

	// Test with empty string
	node = &BinaryNode[string]{Value: "", Left: nil, Right: nil}
	value = node.GetValue()
	if value != "" {
		t.Errorf("Expected empty string, got '%s'", value)
	}

	// Test nil node with string
	var nilNode *BinaryNode[string]
	value = nilNode.GetValue()
	if value != "" {
		t.Errorf("Expected empty string for nil node, got '%s'", value)
	}
}

func TestBinaryNode_FloatValues(t *testing.T) {
	// Test with float values
	node := &BinaryNode[float64]{Value: 3.14, Left: nil, Right: nil}
	value := node.GetValue()
	if value != 3.14 {
		t.Errorf("Expected value 3.14, got %f", value)
	}

	// Test with negative float
	node = &BinaryNode[float64]{Value: -2.5, Left: nil, Right: nil}
	value = node.GetValue()
	if value != -2.5 {
		t.Errorf("Expected value -2.5, got %f", value)
	}

	// Test nil node with float
	var nilNode *BinaryNode[float64]
	value = nilNode.GetValue()
	if value != 0.0 {
		t.Errorf("Expected 0.0 for nil node, got %f", value)
	}
} 
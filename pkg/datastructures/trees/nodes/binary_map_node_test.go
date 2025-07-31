package nodes

import (
	"testing"
)

func TestBinaryMapNode_New(t *testing.T) {
	node := &BinaryMapNode[int, string]{
		Key:   5,
		Value: "five",
		Left:  nil,
		Right: nil,
	}

	if node.Key != 5 {
		t.Errorf("Expected key 5, got %d", node.Key)
	}
	if node.Value != "five" {
		t.Errorf("Expected value 'five', got '%s'", node.Value)
	}
	if node.Left != nil {
		t.Error("Expected left child to be nil")
	}
	if node.Right != nil {
		t.Error("Expected right child to be nil")
	}
}

func TestBinaryMapNode_GetLeft(t *testing.T) {
	// Test with nil node
	var nilNode *BinaryMapNode[int, string]
	left := nilNode.GetLeft()
	if left != nil {
		t.Error("Expected nil for nil node")
	}

	// Test with node that has nil left child
	node := &BinaryMapNode[int, string]{Key: 5, Value: "five", Left: nil, Right: nil}
	left = node.GetLeft()
	if left != nil {
		t.Error("Expected nil for node with nil left child")
	}

	// Test with node that has left child
	leftChild := &BinaryMapNode[int, string]{Key: 3, Value: "three", Left: nil, Right: nil}
	node.Left = leftChild
	left = node.GetLeft()
	if left == nil {
		t.Error("Expected non-nil left child")
	}
	if left.GetValue() != 3 {
		t.Errorf("Expected left child key 3, got %d", left.GetValue())
	}
}

func TestBinaryMapNode_GetRight(t *testing.T) {
	// Test with nil node
	var nilNode *BinaryMapNode[int, string]
	right := nilNode.GetRight()
	if right != nil {
		t.Error("Expected nil for nil node")
	}

	// Test with node that has nil right child
	node := &BinaryMapNode[int, string]{Key: 5, Value: "five", Left: nil, Right: nil}
	right = node.GetRight()
	if right != nil {
		t.Error("Expected nil for node with nil right child")
	}

	// Test with node that has right child
	rightChild := &BinaryMapNode[int, string]{Key: 7, Value: "seven", Left: nil, Right: nil}
	node.Right = rightChild
	right = node.GetRight()
	if right == nil {
		t.Error("Expected non-nil right child")
	}
	if right.GetValue() != 7 {
		t.Errorf("Expected right child key 7, got %d", right.GetValue())
	}
}

func TestBinaryMapNode_GetValue(t *testing.T) {
	// Test with nil node
	var nilNode *BinaryMapNode[int, string]
	value := nilNode.GetValue()
	if value != 0 {
		t.Errorf("Expected zero value for nil node, got %d", value)
	}

	// Test with valid node
	node := &BinaryMapNode[int, string]{Key: 5, Value: "five", Left: nil, Right: nil}
	value = node.GetValue()
	if value != 5 {
		t.Errorf("Expected key 5, got %d", value)
	}

	// Test with negative key
	node = &BinaryMapNode[int, string]{Key: -5, Value: "negative", Left: nil, Right: nil}
	value = node.GetValue()
	if value != -5 {
		t.Errorf("Expected key -5, got %d", value)
	}
}

func TestBinaryMapNode_IsNil(t *testing.T) {
	// Test with nil node
	var nilNode *BinaryMapNode[int, string]
	if !nilNode.IsNil() {
		t.Error("Expected nil node to return true for IsNil")
	}

	// Test with valid node
	node := &BinaryMapNode[int, string]{Key: 5, Value: "five", Left: nil, Right: nil}
	if node.IsNil() {
		t.Error("Expected non-nil node to return false for IsNil")
	}
}

func TestBinaryMapNode_ComplexStructure(t *testing.T) {
	// Create a complex tree structure
	root := &BinaryMapNode[int, string]{Key: 5, Value: "five", Left: nil, Right: nil}
	left := &BinaryMapNode[int, string]{Key: 3, Value: "three", Left: nil, Right: nil}
	right := &BinaryMapNode[int, string]{Key: 7, Value: "seven", Left: nil, Right: nil}
	leftLeft := &BinaryMapNode[int, string]{Key: 1, Value: "one", Left: nil, Right: nil}
	leftRight := &BinaryMapNode[int, string]{Key: 4, Value: "four", Left: nil, Right: nil}
	rightLeft := &BinaryMapNode[int, string]{Key: 6, Value: "six", Left: nil, Right: nil}
	rightRight := &BinaryMapNode[int, string]{Key: 9, Value: "nine", Left: nil, Right: nil}

	// Build the tree
	root.Left = left
	root.Right = right
	left.Left = leftLeft
	left.Right = leftRight
	right.Left = rightLeft
	right.Right = rightRight

	// Test root
	if root.GetValue() != 5 {
		t.Errorf("Expected root key 5, got %d", root.GetValue())
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
		t.Errorf("Expected left child key 3, got %d", leftChild.GetValue())
	}

	// Test right subtree
	rightChild := root.GetRight()
	if rightChild == nil {
		t.Error("Expected non-nil right child")
	}
	if rightChild.GetValue() != 7 {
		t.Errorf("Expected right child key 7, got %d", rightChild.GetValue())
	}

	// Test deeper levels
	leftLeftChild := leftChild.GetLeft()
	if leftLeftChild == nil {
		t.Error("Expected non-nil left-left child")
	}
	if leftLeftChild.GetValue() != 1 {
		t.Errorf("Expected left-left child key 1, got %d", leftLeftChild.GetValue())
	}

	rightRightChild := rightChild.GetRight()
	if rightRightChild == nil {
		t.Error("Expected non-nil right-right child")
	}
	if rightRightChild.GetValue() != 9 {
		t.Errorf("Expected right-right child key 9, got %d", rightRightChild.GetValue())
	}
}

func TestBinaryMapNode_StringKeys(t *testing.T) {
	// Test with string keys
	node := &BinaryMapNode[string, int]{Key: "test", Value: 42, Left: nil, Right: nil}
	key := node.GetValue()
	if key != "test" {
		t.Errorf("Expected key 'test', got '%s'", key)
	}

	// Test with empty string key
	node = &BinaryMapNode[string, int]{Key: "", Value: 0, Left: nil, Right: nil}
	key = node.GetValue()
	if key != "" {
		t.Errorf("Expected empty string key, got '%s'", key)
	}

	// Test nil node with string key
	var nilNode *BinaryMapNode[string, int]
	key = nilNode.GetValue()
	if key != "" {
		t.Errorf("Expected empty string for nil node, got '%s'", key)
	}
}

func TestBinaryMapNode_FloatKeys(t *testing.T) {
	// Test with float keys
	node := &BinaryMapNode[float64, string]{Key: 3.14, Value: "pi", Left: nil, Right: nil}
	key := node.GetValue()
	if key != 3.14 {
		t.Errorf("Expected key 3.14, got %f", key)
	}

	// Test with negative float key
	node = &BinaryMapNode[float64, string]{Key: -2.5, Value: "negative", Left: nil, Right: nil}
	key = node.GetValue()
	if key != -2.5 {
		t.Errorf("Expected key -2.5, got %f", key)
	}

	// Test nil node with float key
	var nilNode *BinaryMapNode[float64, string]
	key = nilNode.GetValue()
	if key != 0.0 {
		t.Errorf("Expected 0.0 for nil node, got %f", key)
	}
}

func TestBinaryMapNode_ComplexValues(t *testing.T) {
	// Test with complex value types
	type Person struct {
		Name string
		Age  int
	}

	node := &BinaryMapNode[int, Person]{
		Key:   1,
		Value: Person{Name: "Alice", Age: 30},
		Left:  nil,
		Right: nil,
	}

	key := node.GetValue()
	if key != 1 {
		t.Errorf("Expected key 1, got %d", key)
	}
	if node.Value.Name != "Alice" {
		t.Errorf("Expected name 'Alice', got '%s'", node.Value.Name)
	}
	if node.Value.Age != 30 {
		t.Errorf("Expected age 30, got %d", node.Value.Age)
	}
}

func TestBinaryMapNode_InterfaceValues(t *testing.T) {
	// Test with interface values
	node := &BinaryMapNode[int, interface{}]{
		Key:   42,
		Value: "string value",
		Left:  nil,
		Right: nil,
	}

	key := node.GetValue()
	if key != 42 {
		t.Errorf("Expected key 42, got %d", key)
	}

	// Test with different value types
	node.Value = 123
	if node.Value != 123 {
		t.Errorf("Expected value 123, got %v", node.Value)
	}

	node.Value = []int{1, 2, 3}
	if len(node.Value.([]int)) != 3 {
		t.Errorf("Expected slice of length 3, got %d", len(node.Value.([]int)))
	}
} 
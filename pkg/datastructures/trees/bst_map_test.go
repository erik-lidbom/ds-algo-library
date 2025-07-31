package trees

import (
	"testing"
)

func TestNewBSTMap(t *testing.T) {
	bstm := NewBinarySearchTreeMap[int, string]()
	if bstm == nil {
		t.Fatal("NewBinarySearchTreeMap returned nil")
	}
	if bstm.Size() != 0 {
		t.Errorf("Expected size of 0, got %d", bstm.Size())
	}
	if !bstm.IsEmpty() {
		t.Error("Expected empty BinarySearchTreeMap")
	}
}

func TestBSTMap_Size(t *testing.T) {
	bstm := NewBinarySearchTreeMap[int, string]()

	if bstm.Size() != 0 {
		t.Errorf("Expected size of 0, got %d", bstm.Size())
	}

	bstm.Put(1, "test")

	if bstm.Size() != 1 {
		t.Errorf("Expected size 1, got %d", bstm.Size())
	}

	bstm.Put(2, "test")

	if bstm.Size() != 2 {
		t.Errorf("Expected size 2, got %d", bstm.Size())
	}

	bstm.Remove(2)

	if bstm.Size() != 1 {
		t.Errorf("Expected size 1, got %d", bstm.Size())
	}
}

func TestBSTMap_Put(t *testing.T) {
	cases := []struct {
		name         string
		entries      [][2]any
		expectedSize int
	}{
		{"put single entry", [][2]any{{1, "one"}}, 1},
		{"put multiple entries", [][2]any{{1, "one"}, {2, "two"}, {3, "three"}}, 3},
		{"put zero key", [][2]any{{0, "zero"}}, 1},
		{"put negative key", [][2]any{{-1, "negative"}}, 1},
		{"put empty string value", [][2]any{{1, ""}}, 1},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			bstm := NewBinarySearchTreeMap[int, string]()

			for _, entry := range c.entries {
				key := entry[0].(int)
				value := entry[1].(string)
				bstm.Put(key, value)
			}

			if bstm.Size() != c.expectedSize {
				t.Errorf("Expected size %d, got %d", c.expectedSize, bstm.Size())
			}
		})
	}
}

func TestBSTMap_Get(t *testing.T) {
	bstm := NewBinarySearchTreeMap[int, string]()

	// Test getting from empty map
	_, err := bstm.Get(1)
	if err == nil {
		t.Error("Expected error when getting from empty map")
	}

	// Test getting existing key
	bstm.Put(1, "one")
	value, err := bstm.Get(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if value != "one" {
		t.Errorf("Expected 'one', got '%s'", value)
	}

	// Test getting non-existing key
	_, err = bstm.Get(2)
	if err == nil {
		t.Error("Expected error when getting non-existing key")
	}

	// Test updating existing key
	bstm.Put(1, "updated")
	value, err = bstm.Get(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != "updated" {
		t.Errorf("Expected 'updated', got '%s'", value)
	}
}

func TestBSTMap_ContainsKey(t *testing.T) {
	bstm := NewBinarySearchTreeMap[int, string]()

	// Test empty map
	if bstm.ContainsKey(1) {
		t.Error("Empty map should not contain any keys")
	}

	// Test existing key
	bstm.Put(1, "one")
	if !bstm.ContainsKey(1) {
		t.Error("Map should contain key 1")
	}

	// Test non-existing key
	if bstm.ContainsKey(2) {
		t.Error("Map should not contain key 2")
	}

	// Test after removal
	bstm.Remove(1)
	if bstm.ContainsKey(1) {
		t.Error("Map should not contain removed key")
	}
}

func TestBSTMap_Remove(t *testing.T) {
	bstm := NewBinarySearchTreeMap[int, string]()

	// Test removing from empty map
	_, err := bstm.Remove(1)
	if err == nil {
		t.Error("Expected error when removing from empty map")
	}

	// Test removing existing key
	bstm.Put(1, "one")
	value, err := bstm.Remove(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != "one" {
		t.Errorf("Expected 'one', got '%s'", value)
	}
	if bstm.Size() != 0 {
		t.Errorf("Expected size 0 after removal, got %d", bstm.Size())
	}

	// Test removing non-existing key
	_, err = bstm.Remove(1)
	if err == nil {
		t.Error("Expected error when removing non-existing key")
	}
}

func TestBSTMap_Search(t *testing.T) {
	bstm := NewBinarySearchTreeMap[int, string]()

	// Test searching empty map
	value, found := bstm.Search(1)
	if found {
		t.Error("Should not find key in empty map")
	}
	if value != "" {
		t.Errorf("Expected empty string, got '%s'", value)
	}

	// Test searching existing key
	bstm.Put(1, "one")
	value, found = bstm.Search(1)
	if !found {
		t.Error("Should find existing key")
	}
	if value != "one" {
		t.Errorf("Expected 'one', got '%s'", value)
	}

	// Test searching non-existing key
	value, found = bstm.Search(2)
	if found {
		t.Error("Should not find non-existing key")
	}
	if value != "" {
		t.Errorf("Expected empty string, got '%s'", value)
	}
}

func TestBSTMap_Delete(t *testing.T) {
	bstm := NewBinarySearchTreeMap[int, string]()

	// Test deleting from empty map
	err := bstm.Delete(1)
	if err == nil {
		t.Error("Expected error when deleting from empty map")
	}

	// Test deleting existing key
	bstm.Put(1, "one")
	err = bstm.Delete(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if bstm.Size() != 0 {
		t.Errorf("Expected size 0 after deletion, got %d", bstm.Size())
	}

	// Test deleting non-existing key
	err = bstm.Delete(1)
	if err == nil {
		t.Error("Expected error when deleting non-existing key")
	}
}

func TestBSTMap_KeysBetween(t *testing.T) {
	bstm := NewBinarySearchTreeMap[int, string]()

	// Test empty map
	keys := bstm.KeysBetween(1, 5)
	if keys.Size() != 0 {
		t.Errorf("Expected 0 keys, got %d", keys.Size())
	}

	// Test with data
	bstm.Put(1, "one")
	bstm.Put(2, "two")
	bstm.Put(3, "three")
	bstm.Put(4, "four")
	bstm.Put(5, "five")

	// Test range within bounds
	keys = bstm.KeysBetween(2, 4)
	if keys.Size() != 3 {
		t.Errorf("Expected 3 keys, got %d", keys.Size())
	}

	// Test range outside bounds
	keys = bstm.KeysBetween(6, 10)
	if keys.Size() != 0 {
		t.Errorf("Expected 0 keys, got %d", keys.Size())
	}

	// Test invalid range
	keys = bstm.KeysBetween(5, 1)
	if keys.Size() != 0 {
		t.Errorf("Expected 0 keys for invalid range, got %d", keys.Size())
	}
}

func TestBSTMap_Insert(t *testing.T) {
	bstm := NewBinarySearchTreeMap[int, string]()

	// Test inserting new key
	err := bstm.Insert(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if bstm.Size() != 1 {
		t.Errorf("Expected size 1, got %d", bstm.Size())
	}

	// Test inserting existing key
	err = bstm.Insert(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if bstm.Size() != 1 {
		t.Errorf("Expected size 1, got %d", bstm.Size())
	}
}

func TestBSTMap_ComplexOperations(t *testing.T) {
	bstm := NewBinarySearchTreeMap[int, string]()

	// Test complex sequence of operations
	bstm.Put(5, "five")
	bstm.Put(3, "three")
	bstm.Put(7, "seven")
	bstm.Put(1, "one")
	bstm.Put(9, "nine")

	if bstm.Size() != 5 {
		t.Errorf("Expected size 5, got %d", bstm.Size())
	}

	// Test all keys exist
	keys := []int{1, 3, 5, 7, 9}
	for _, key := range keys {
		if !bstm.ContainsKey(key) {
			t.Errorf("Key %d should exist", key)
		}
	}

	// Test removal of leaf node
	value, err := bstm.Remove(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != "one" {
		t.Errorf("Expected 'one', got '%s'", value)
	}
	if bstm.Size() != 4 {
		t.Errorf("Expected size 4, got %d", bstm.Size())
	}

	// Test removal of node with one child
	value, err = bstm.Remove(3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != "three" {
		t.Errorf("Expected 'three', got '%s'", value)
	}

	// Test removal of node with two children
	value, err = bstm.Remove(5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != "five" {
		t.Errorf("Expected 'five', got '%s'", value)
	}
}

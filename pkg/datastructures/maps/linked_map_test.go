package maps

import (
	"testing"
)

func TestNewLinkedMap(t *testing.T) {
	lm := NewLinkedMap[int, string]()
	if lm == nil {
		t.Fatal("NewLinkedMap returned nil")
	}
	if lm.Size() != 0 {
		t.Errorf("Expected size 0, got %d", lm.Size())
	}
	if !lm.IsEmpty() {
		t.Error("Expected empty map")
	}
	if lm.head != nil {
		t.Error("Expected head to be nil")
	}
}

func TestLinkedMap_Size(t *testing.T) {
	lm := NewLinkedMap[int, string]()
	if lm.Size() != 0 {
		t.Errorf("Expected size 0, got %d", lm.Size())
	}
	
	lm.Put(1, "one")
	if lm.Size() != 1 {
		t.Errorf("Expected size 1, got %d", lm.Size())
	}
	
	lm.Put(2, "two")
	if lm.Size() != 2 {
		t.Errorf("Expected size 2, got %d", lm.Size())
	}
}

func TestLinkedMap_IsEmpty(t *testing.T) {
	lm := NewLinkedMap[int, string]()
	if !lm.IsEmpty() {
		t.Error("Expected empty map")
	}
	
	lm.Put(1, "one")
	if lm.IsEmpty() {
		t.Error("Expected non-empty map")
	}
}

func TestLinkedMap_Put(t *testing.T) {
	cases := []struct {
		name string
		entries [][2]any
		expectSize int
	}{
		{"put single entry", [][2]any{{1, "one"}}, 1},
		{"put multiple entries", [][2]any{{1, "one"}, {2, "two"}, {3, "three"}}, 3},
		{"put zero key", [][2]any{{0, "zero"}}, 1},
		{"put negative key", [][2]any{{-1, "negative"}}, 1},
		{"put empty string value", [][2]any{{1, ""}}, 1},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			lm := NewLinkedMap[int, string]()
			
			for _, entry := range c.entries {
				key := entry[0].(int)
				value := entry[1].(string)
				lm.Put(key, value)
			}
			
			if lm.Size() != c.expectSize {
				t.Errorf("Expected size %d, got %d", c.expectSize, lm.Size())
			}
		})
	}
}

func TestLinkedMap_Put_Update(t *testing.T) {
	lm := NewLinkedMap[int, string]()
	
	lm.Put(1, "one")
	if lm.Size() != 1 {
		t.Errorf("Expected size 1, got %d", lm.Size())
	}
	
	lm.Put(1, "updated")
	if lm.Size() != 1 {
		t.Errorf("Expected size 1 after update, got %d", lm.Size())
	}
	
	value, err := lm.Get(1)
	if err != nil {
		t.Errorf("Get failed: %v", err)
	}
	if value != "updated" {
		t.Errorf("Expected value 'updated', got '%s'", value)
	}
}

func TestLinkedMap_Get(t *testing.T) {
	cases := []struct {
		name string
		entries [][2]any
		searchKey int
		expectValue string
		expectError bool
	}{
		{"get existing key", [][2]any{{1, "one"}, {2, "two"}}, 1, "one", false},
		{"get non-existent key", [][2]any{{1, "one"}, {2, "two"}}, 3, "", true},
		{"get from empty map", [][2]any{}, 1, "", true},
		{"get zero key", [][2]any{{0, "zero"}}, 0, "zero", false},
		{"get negative key", [][2]any{{-1, "negative"}}, -1, "negative", false},
		{"get empty string value", [][2]any{{1, ""}}, 1, "", false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			lm := NewLinkedMap[int, string]()
			
			for _, entry := range c.entries {
				key := entry[0].(int)
				value := entry[1].(string)
				lm.Put(key, value)
			}
			
			value, err := lm.Get(c.searchKey)
			
			if c.expectError {
				if err == nil {
					t.Error("Expected error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}
				if value != c.expectValue {
					t.Errorf("Expected value '%s', got '%s'", c.expectValue, value)
				}
			}
		})
	}
}

func TestLinkedMap_Remove(t *testing.T) {
	cases := []struct {
		name string
		entries [][2]any
		removeKey int
		expectSize int
		expectError bool
		expectRemoved string
	}{
		{"remove existing key", [][2]any{{1, "one"}, {2, "two"}}, 1, 1, false, "one"},
		{"remove non-existent key", [][2]any{{1, "one"}, {2, "two"}}, 3, 2, true, ""},
		{"remove from empty map", [][2]any{}, 1, 0, true, ""},
		{"remove single entry", [][2]any{{1, "one"}}, 1, 0, false, "one"},
		{"remove first entry", [][2]any{{1, "one"}, {2, "two"}, {3, "three"}}, 1, 2, false, "one"},
		{"remove last entry", [][2]any{{1, "one"}, {2, "two"}, {3, "three"}}, 3, 2, false, "three"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			lm := NewLinkedMap[int, string]()
			
			for _, entry := range c.entries {
				key := entry[0].(int)
				value := entry[1].(string)
				lm.Put(key, value)
			}
			
			removed, err := lm.Remove(c.removeKey)
			
			if c.expectError {
				if err == nil {
					t.Error("Expected error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}
				if removed != c.expectRemoved {
					t.Errorf("Expected removed value '%s', got '%s'", c.expectRemoved, removed)
				}
			}
			
			if lm.Size() != c.expectSize {
				t.Errorf("Expected size %d, got %d", c.expectSize, lm.Size())
			}
		})
	}
}

func TestLinkedMap_ContainsKey(t *testing.T) {
	cases := []struct {
		name string
		entries [][2]any
		searchKey int
		expectFound bool
	}{
		{"contains existing key", [][2]any{{1, "one"}, {2, "two"}}, 1, true},
		{"contains second key", [][2]any{{1, "one"}, {2, "two"}}, 2, true},
		{"does not contain key", [][2]any{{1, "one"}, {2, "two"}}, 3, false},
		{"empty map", [][2]any{}, 1, false},
		{"single key found", [][2]any{{1, "one"}}, 1, true},
		{"single key not found", [][2]any{{1, "one"}}, 2, false},
		{"zero key", [][2]any{{0, "zero"}}, 0, true},
		{"negative key", [][2]any{{-1, "negative"}}, -1, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			lm := NewLinkedMap[int, string]()
			
			for _, entry := range c.entries {
				key := entry[0].(int)
				value := entry[1].(string)
				lm.Put(key, value)
			}
			
			found := lm.ContainsKey(c.searchKey)
			if found != c.expectFound {
				t.Errorf("Expected found=%v, got %v", c.expectFound, found)
			}
		})
	}
}

func TestLinkedMap_EdgeCases(t *testing.T) {
	t.Run("put and remove multiple times", func(t *testing.T) {
		lm := NewLinkedMap[int, string]()
		
		// Put entries
		lm.Put(1, "one")
		lm.Put(2, "two")
		lm.Put(3, "three")
		
		if lm.Size() != 3 {
			t.Errorf("Expected size 3, got %d", lm.Size())
		}
		
		// Remove middle entry
		removed, err := lm.Remove(2)
		if err != nil {
			t.Errorf("Remove failed: %v", err)
		}
		if removed != "two" {
			t.Errorf("Expected removed 'two', got '%s'", removed)
		}
		
		if lm.Size() != 2 {
			t.Errorf("Expected size 2, got %d", lm.Size())
		}
		
		// Verify remaining entries
		if !lm.ContainsKey(1) || !lm.ContainsKey(3) {
			t.Error("Expected keys 1 and 3 to remain")
		}
		if lm.ContainsKey(2) {
			t.Error("Expected key 2 to be removed")
		}
	})
	
	t.Run("put same key after remove", func(t *testing.T) {
		lm := NewLinkedMap[int, string]()
		lm.Put(1, "one")
		lm.Remove(1)
		lm.Put(1, "new") // Should work after removal
		
		if lm.Size() != 1 {
			t.Errorf("Expected size 1, got %d", lm.Size())
		}
		if !lm.ContainsKey(1) {
			t.Error("Expected key 1 to be present")
		}
		
		value, _ := lm.Get(1)
		if value != "new" {
			t.Errorf("Expected value 'new', got '%s'", value)
		}
	})
	
	t.Run("update existing key", func(t *testing.T) {
		lm := NewLinkedMap[int, string]()
		lm.Put(1, "old")
		lm.Put(1, "new")
		
		if lm.Size() != 1 {
			t.Errorf("Expected size 1, got %d", lm.Size())
		}
		
		value, _ := lm.Get(1)
		if value != "new" {
			t.Errorf("Expected value 'new', got '%s'", value)
		}
	})
}

func BenchmarkLinkedMap_Put(b *testing.B) {
	for n := 0; n < b.N; n++ {
		lm := NewLinkedMap[int, string]()
		for i := 0; i < 1000; i++ {
			lm.Put(i, "value")
		}
	}
}

func BenchmarkLinkedMap_Get(b *testing.B) {
	lm := NewLinkedMap[int, string]()
	for i := 0; i < 1000; i++ {
		lm.Put(i, "value")
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		lm.Get(n % 1000)
	}
}

func BenchmarkLinkedMap_Remove(b *testing.B) {
	for n := 0; n < b.N; n++ {
		lm := NewLinkedMap[int, string]()
		for i := 0; i < 1000; i++ {
			lm.Put(i, "value")
		}
		for i := 0; i < 1000; i++ {
			lm.Remove(i)
		}
	}
} 
package maps

import (
	"testing"
)

func TestNewSortedListMap(t *testing.T) {
	slm := NewSortedListMap[int, string]()
	if slm == nil {
		t.Fatal("NewSortedListMap returned nil")
	}
	if slm.Size() != 0 {
		t.Errorf("Expected size 0, got %d", slm.Size())
	}
	if !slm.IsEmpty() {
		t.Error("Expected empty map")
	}
	if slm.keys == nil {
		t.Error("Expected keys array to be initialized")
	}
	if slm.values == nil {
		t.Error("Expected values array to be initialized")
	}
}

func TestSortedListMap_Size(t *testing.T) {
	slm := NewSortedListMap[int, string]()
	if slm.Size() != 0 {
		t.Errorf("Expected size 0, got %d", slm.Size())
	}
	
	slm.Put(1, "one")
	if slm.Size() != 1 {
		t.Errorf("Expected size 1, got %d", slm.Size())
	}
	
	slm.Put(2, "two")
	if slm.Size() != 2 {
		t.Errorf("Expected size 2, got %d", slm.Size())
	}
}

func TestSortedListMap_IsEmpty(t *testing.T) {
	slm := NewSortedListMap[int, string]()
	if !slm.IsEmpty() {
		t.Error("Expected empty map")
	}
	
	slm.Put(1, "one")
	if slm.IsEmpty() {
		t.Error("Expected non-empty map")
	}
}

func TestSortedListMap_Put(t *testing.T) {
	cases := []struct {
		name string
		entries [][2]any
		expectSize int
		expectKeys []int
		expectValues []string
	}{
		{"put single entry", [][2]any{{1, "one"}}, 1, []int{1}, []string{"one"}},
		{"put multiple entries", [][2]any{{3, "three"}, {1, "one"}, {2, "two"}}, 3, []int{1, 2, 3}, []string{"one", "two", "three"}},
		{"put zero key", [][2]any{{0, "zero"}}, 1, []int{0}, []string{"zero"}},
		{"put negative key", [][2]any{{-1, "negative"}}, 1, []int{-1}, []string{"negative"}},
		{"put empty string value", [][2]any{{1, ""}}, 1, []int{1}, []string{""}},
		{"put in reverse order", [][2]any{{5, "five"}, {4, "four"}, {3, "three"}, {2, "two"}, {1, "one"}}, 5, []int{1, 2, 3, 4, 5}, []string{"one", "two", "three", "four", "five"}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			slm := NewSortedListMap[int, string]()
			
			for _, entry := range c.entries {
				key := entry[0].(int)
				value := entry[1].(string)
				slm.Put(key, value)
			}
			
			if slm.Size() != c.expectSize {
				t.Errorf("Expected size %d, got %d", c.expectSize, slm.Size())
			}
			
			// Verify keys are sorted
			for i := 0; i < slm.Size(); i++ {
				key, _ := slm.keys.Get(i)
				if key != c.expectKeys[i] {
					t.Errorf("At index %d: got key %d, want %d", i, key, c.expectKeys[i])
				}
			}
			
			// Verify values correspond to sorted keys
			for i := 0; i < slm.Size(); i++ {
				value, _ := slm.values.Get(i)
				if value != c.expectValues[i] {
					t.Errorf("At index %d: got value '%s', want '%s'", i, value, c.expectValues[i])
				}
			}
		})
	}
}

func TestSortedListMap_Put_Update(t *testing.T) {
	slm := NewSortedListMap[int, string]()
	
	// Put initial value
	slm.Put(1, "one")
	if slm.Size() != 1 {
		t.Errorf("Expected size 1, got %d", slm.Size())
	}
	
	// Update existing key
	slm.Put(1, "updated")
	if slm.Size() != 1 {
		t.Errorf("Expected size 1 after update, got %d", slm.Size())
	}
	
	// Verify updated value
	value, err := slm.Get(1)
	if err != nil {
		t.Errorf("Get failed: %v", err)
	}
	if value != "updated" {
		t.Errorf("Expected value 'updated', got '%s'", value)
	}
}

func TestSortedListMap_Get(t *testing.T) {
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
		{"get middle key", [][2]any{{1, "one"}, {2, "two"}, {3, "three"}}, 2, "two", false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			slm := NewSortedListMap[int, string]()
			
			for _, entry := range c.entries {
				key := entry[0].(int)
				value := entry[1].(string)
				slm.Put(key, value)
			}
			
			value, err := slm.Get(c.searchKey)
			
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

func TestSortedListMap_Remove(t *testing.T) {
	cases := []struct {
		name string
		entries [][2]any
		removeKey int
		expectSize int
		expectError bool
		expectRemoved string
		expectKeys []int
		expectValues []string
	}{
		{"remove existing key", [][2]any{{1, "one"}, {2, "two"}}, 1, 1, false, "one", []int{2}, []string{"two"}},
		{"remove non-existent key", [][2]any{{1, "one"}, {2, "two"}}, 3, 2, true, "", []int{1, 2}, []string{"one", "two"}},
		{"remove from empty map", [][2]any{}, 1, 0, true, "", []int{}, []string{}},
		{"remove single entry", [][2]any{{1, "one"}}, 1, 0, false, "one", []int{}, []string{}},
		{"remove first entry", [][2]any{{1, "one"}, {2, "two"}, {3, "three"}}, 1, 2, false, "one", []int{2, 3}, []string{"two", "three"}},
		{"remove last entry", [][2]any{{1, "one"}, {2, "two"}, {3, "three"}}, 3, 2, false, "three", []int{1, 2}, []string{"one", "two"}},
		{"remove middle entry", [][2]any{{1, "one"}, {2, "two"}, {3, "three"}}, 2, 2, false, "two", []int{1, 3}, []string{"one", "three"}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			slm := NewSortedListMap[int, string]()
			
			for _, entry := range c.entries {
				key := entry[0].(int)
				value := entry[1].(string)
				slm.Put(key, value)
			}
			
			removed, err := slm.Remove(c.removeKey)
			
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
			
			if slm.Size() != c.expectSize {
				t.Errorf("Expected size %d, got %d", c.expectSize, slm.Size())
			}
			
			// Verify remaining keys are still sorted
			for i := 0; i < slm.Size(); i++ {
				key, _ := slm.keys.Get(i)
				if key != c.expectKeys[i] {
					t.Errorf("At index %d: got key %d, want %d", i, key, c.expectKeys[i])
				}
			}
			
			// Verify remaining values correspond to sorted keys
			for i := 0; i < slm.Size(); i++ {
				value, _ := slm.values.Get(i)
				if value != c.expectValues[i] {
					t.Errorf("At index %d: got value '%s', want '%s'", i, value, c.expectValues[i])
				}
			}
		})
	}
}

func TestSortedListMap_ContainsKey(t *testing.T) {
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
		{"large set", [][2]any{{1, "one"}, {2, "two"}, {3, "three"}, {4, "four"}, {5, "five"}}, 3, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			slm := NewSortedListMap[int, string]()
			
			for _, entry := range c.entries {
				key := entry[0].(int)
				value := entry[1].(string)
				slm.Put(key, value)
			}
			
			found := slm.ContainsKey(c.searchKey)
			if found != c.expectFound {
				t.Errorf("Expected found=%v, got %v", c.expectFound, found)
			}
		})
	}
}

func TestSortedListMap_EdgeCases(t *testing.T) {
	t.Run("put and remove multiple times", func(t *testing.T) {
		slm := NewSortedListMap[int, string]()
		
		// Put entries
		slm.Put(3, "three")
		slm.Put(1, "one")
		slm.Put(2, "two")
		
		if slm.Size() != 3 {
			t.Errorf("Expected size 3, got %d", slm.Size())
		}
		
		// Remove middle entry
		removed, err := slm.Remove(2)
		if err != nil {
			t.Errorf("Remove failed: %v", err)
		}
		if removed != "two" {
			t.Errorf("Expected removed 'two', got '%s'", removed)
		}
		
		if slm.Size() != 2 {
			t.Errorf("Expected size 2, got %d", slm.Size())
		}
		
		// Verify remaining entries
		if !slm.ContainsKey(1) || !slm.ContainsKey(3) {
			t.Error("Expected keys 1 and 3 to remain")
		}
		if slm.ContainsKey(2) {
			t.Error("Expected key 2 to be removed")
		}
	})
	
	t.Run("put same key after remove", func(t *testing.T) {
		slm := NewSortedListMap[int, string]()
		slm.Put(1, "one")
		slm.Remove(1)
		slm.Put(1, "new") // Should work after removal
		
		if slm.Size() != 1 {
			t.Errorf("Expected size 1, got %d", slm.Size())
		}
		if !slm.ContainsKey(1) {
			t.Error("Expected key 1 to be present")
		}
		
		value, _ := slm.Get(1)
		if value != "new" {
			t.Errorf("Expected value 'new', got '%s'", value)
		}
	})
	
	t.Run("maintain sorted order", func(t *testing.T) {
		slm := NewSortedListMap[int, string]()
		entries := [][2]any{
			{5, "five"}, {2, "two"}, {8, "eight"}, {1, "one"}, 
			{9, "nine"}, {3, "three"}, {7, "seven"}, {4, "four"}, {6, "six"},
		}
		
		for _, entry := range entries {
			key := entry[0].(int)
			value := entry[1].(string)
			slm.Put(key, value)
		}
		
		expectedKeys := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		expectedValues := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
		
		for i := 0; i < slm.Size(); i++ {
			key, _ := slm.keys.Get(i)
			if key != expectedKeys[i] {
				t.Errorf("At index %d: got key %d, want %d", i, key, expectedKeys[i])
			}
			
			value, _ := slm.values.Get(i)
			if value != expectedValues[i] {
				t.Errorf("At index %d: got value '%s', want '%s'", i, value, expectedValues[i])
			}
		}
	})
}

func BenchmarkSortedListMap_Put(b *testing.B) {
	for n := 0; n < b.N; n++ {
		slm := NewSortedListMap[int, string]()
		for i := 0; i < 1000; i++ {
			slm.Put(i, "value")
		}
	}
}

func BenchmarkSortedListMap_Get(b *testing.B) {
	slm := NewSortedListMap[int, string]()
	for i := 0; i < 1000; i++ {
		slm.Put(i, "value")
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		slm.Get(n % 1000)
	}
}

func BenchmarkSortedListMap_Remove(b *testing.B) {
	for n := 0; n < b.N; n++ {
		slm := NewSortedListMap[int, string]()
		for i := 0; i < 1000; i++ {
			slm.Put(i, "value")
		}
		for i := 0; i < 1000; i++ {
			slm.Remove(i)
		}
	}
} 
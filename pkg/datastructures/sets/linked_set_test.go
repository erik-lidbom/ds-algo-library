package sets

import (
	"testing"
)

func TestNewLinkedSet(t *testing.T) {
	ls := NewLinkedSet[int]()
	if ls == nil {
		t.Fatal("NewLinkedSet returned nil")
	}
	if ls.Size() != 0 {
		t.Errorf("Expected size 0, got %d", ls.Size())
	}
	if !ls.IsEmpty() {
		t.Error("Expected empty set")
	}
	if ls.head != nil {
		t.Error("Expected head to be nil")
	}
}

func TestLinkedSet_Size(t *testing.T) {
	ls := NewLinkedSet[int]()
	if ls.Size() != 0 {
		t.Errorf("Expected size 0, got %d", ls.Size())
	}
	
	ls.Add(1)
	if ls.Size() != 1 {
		t.Errorf("Expected size 1, got %d", ls.Size())
	}
	
	ls.Add(2)
	if ls.Size() != 2 {
		t.Errorf("Expected size 2, got %d", ls.Size())
	}
}

func TestLinkedSet_IsEmpty(t *testing.T) {
	ls := NewLinkedSet[int]()
	if !ls.IsEmpty() {
		t.Error("Expected empty set")
	}
	
	ls.Add(1)
	if ls.IsEmpty() {
		t.Error("Expected non-empty set")
	}
}

func TestLinkedSet_Add(t *testing.T) {
	cases := []struct {
		name string
		elements []int
		expectSize int
		expectError bool
	}{
		{"add single element", []int{1}, 1, false},
		{"add multiple elements", []int{1, 2, 3}, 3, false},
		{"add duplicate element", []int{1, 1}, 1, true},
		{"add zero value", []int{0}, 1, false},
		{"add negative values", []int{-1, -2}, 2, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ls := NewLinkedSet[int]()
			var lastErr error
			
			for _, elem := range c.elements {
				lastErr = ls.Add(elem)
			}
			
			if c.expectError {
				if lastErr == nil {
					t.Error("Expected error but got nil")
				}
			} else {
				if lastErr != nil {
					t.Errorf("Expected no error but got: %v", lastErr)
				}
			}
			
			if ls.Size() != c.expectSize {
				t.Errorf("Expected size %d, got %d", c.expectSize, ls.Size())
			}
		})
	}
}

func TestLinkedSet_Remove(t *testing.T) {
	cases := []struct {
		name string
		initialElements []int
		removeElement int
		expectSize int
		expectError bool
		expectRemoved int
	}{
		{"remove existing element", []int{1, 2, 3}, 2, 2, false, 2},
		{"remove first element", []int{1, 2, 3}, 1, 2, false, 1},
		{"remove last element", []int{1, 2, 3}, 3, 2, false, 3},
		{"remove non-existent element", []int{1, 2, 3}, 4, 3, true, 0},
		{"remove from empty set", []int{}, 1, 0, true, 0},
		{"remove single element", []int{1}, 1, 0, false, 1},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ls := NewLinkedSet[int]()
			for _, elem := range c.initialElements {
				ls.Add(elem)
			}
			
			removed, err := ls.Remove(c.removeElement)
			
			if c.expectError {
				if err == nil {
					t.Error("Expected error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}
				if removed != c.expectRemoved {
					t.Errorf("Expected removed value %d, got %d", c.expectRemoved, removed)
				}
			}
			
			if ls.Size() != c.expectSize {
				t.Errorf("Expected size %d, got %d", c.expectSize, ls.Size())
			}
		})
	}
}

func TestLinkedSet_Contains(t *testing.T) {
	cases := []struct {
		name string
		elements []int
		searchElement int
		expectFound bool
	}{
		{"contains existing element", []int{1, 2, 3}, 2, true},
		{"contains first element", []int{1, 2, 3}, 1, true},
		{"contains last element", []int{1, 2, 3}, 3, true},
		{"does not contain element", []int{1, 2, 3}, 4, false},
		{"empty set", []int{}, 1, false},
		{"single element found", []int{1}, 1, true},
		{"single element not found", []int{1}, 2, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ls := NewLinkedSet[int]()
			for _, elem := range c.elements {
				ls.Add(elem)
			}
			
			found := ls.Contains(c.searchElement)
			if found != c.expectFound {
				t.Errorf("Expected found=%v, got %v", c.expectFound, found)
			}
		})
	}
}

func TestLinkedSet_EdgeCases(t *testing.T) {
	t.Run("add and remove multiple times", func(t *testing.T) {
		ls := NewLinkedSet[int]()
		
		// Add elements
		ls.Add(1)
		ls.Add(2)
		ls.Add(3)
		
		if ls.Size() != 3 {
			t.Errorf("Expected size 3, got %d", ls.Size())
		}
		
		// Remove middle element
		removed, err := ls.Remove(2)
		if err != nil {
			t.Errorf("Remove failed: %v", err)
		}
		if removed != 2 {
			t.Errorf("Expected removed 2, got %d", removed)
		}
		
		if ls.Size() != 2 {
			t.Errorf("Expected size 2, got %d", ls.Size())
		}
		
		// Verify remaining elements
		if !ls.Contains(1) || !ls.Contains(3) {
			t.Error("Expected elements 1 and 3 to remain")
		}
		if ls.Contains(2) {
			t.Error("Expected element 2 to be removed")
		}
	})
	
	t.Run("add duplicate after remove", func(t *testing.T) {
		ls := NewLinkedSet[int]()
		ls.Add(1)
		ls.Remove(1)
		ls.Add(1) // Should work after removal
		
		if ls.Size() != 1 {
			t.Errorf("Expected size 1, got %d", ls.Size())
		}
		if !ls.Contains(1) {
			t.Error("Expected element 1 to be present")
		}
	})
}

func BenchmarkLinkedSet_Add(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ls := NewLinkedSet[int]()
		for i := 0; i < 1000; i++ {
			ls.Add(i)
		}
	}
}

func BenchmarkLinkedSet_Contains(b *testing.B) {
	ls := NewLinkedSet[int]()
	for i := 0; i < 1000; i++ {
		ls.Add(i)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ls.Contains(n % 1000)
	}
}

func BenchmarkLinkedSet_Remove(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ls := NewLinkedSet[int]()
		for i := 0; i < 1000; i++ {
			ls.Add(i)
		}
		for i := 0; i < 1000; i++ {
			ls.Remove(i)
		}
	}
} 
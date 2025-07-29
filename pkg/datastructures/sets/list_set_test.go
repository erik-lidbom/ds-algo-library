package sets

import (
	"testing"
)

func TestNewListSet(t *testing.T) {
	ls := NewListSet[int]()
	if ls == nil {
		t.Fatal("NewListSet returned nil")
	}
	if ls.Size() != 0 {
		t.Errorf("Expected size 0, got %d", ls.Size())
	}
	if !ls.IsEmpty() {
		t.Error("Expected empty set")
	}
	if ls.arr == nil {
		t.Error("Expected array to be initialized")
	}
}

func TestListSet_Size(t *testing.T) {
	ls := NewListSet[int]()
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

func TestListSet_IsEmpty(t *testing.T) {
	ls := NewListSet[int]()
	if !ls.IsEmpty() {
		t.Error("Expected empty set")
	}

	ls.Add(1)
	if ls.IsEmpty() {
		t.Error("Expected non-empty set")
	}
}

func TestListSet_Add(t *testing.T) {
	cases := []struct {
		name        string
		elements    []int
		expectSize  int
		expectError bool
		expectOrder []int
	}{
		{"add single element", []int{1}, 1, false, []int{1}},
		{"add multiple elements", []int{3, 1, 2}, 3, false, []int{1, 2, 3}},
		{"add duplicate element", []int{1, 1}, 1, true, []int{1}},
		{"add zero value", []int{0}, 1, false, []int{0}},
		{"add negative values", []int{-1, -2}, 2, false, []int{-2, -1}},
		{"add in reverse order", []int{5, 4, 3, 2, 1}, 5, false, []int{1, 2, 3, 4, 5}},
		{"add in random order", []int{3, 1, 4, 1, 5, 9, 2, 6}, 7, true, []int{1, 2, 3, 4, 5, 6, 9}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ls := NewListSet[int]()
			var hasError bool

			for _, elem := range c.elements {
				err := ls.Add(elem)
				if err != nil {
					hasError = true
				}
			}

			if c.expectError {
				if !hasError {
					t.Error("Expected error but got none")
				}
			} else {
				if hasError {
					t.Error("Expected no error but got one")
				}
			}

			if ls.Size() != c.expectSize {
				t.Errorf("Expected size %d, got %d", c.expectSize, ls.Size())
			}

			// Verify order
			for i := 0; i < ls.Size(); i++ {
				val, _ := ls.arr.Get(i)
				if val != c.expectOrder[i] {
					t.Errorf("At index %d: got %d, want %d", i, val, c.expectOrder[i])
				}
			}
		})
	}
}

func TestListSet_Remove(t *testing.T) {
	cases := []struct {
		name            string
		initialElements []int
		removeElement   int
		expectSize      int
		expectError     bool
		expectRemoved   int
		expectOrder     []int
	}{
		{"remove existing element", []int{1, 2, 3}, 2, 2, false, 2, []int{1, 3}},
		{"remove first element", []int{1, 2, 3}, 1, 2, false, 1, []int{2, 3}},
		{"remove last element", []int{1, 2, 3}, 3, 2, false, 3, []int{1, 2}},
		{"remove non-existent element", []int{1, 2, 3}, 4, 3, true, 0, []int{1, 2, 3}},
		{"remove from empty set", []int{}, 1, 0, true, 0, []int{}},
		{"remove single element", []int{1}, 1, 0, false, 1, []int{}},
		{"remove middle element", []int{1, 2, 3, 4, 5}, 3, 4, false, 3, []int{1, 2, 4, 5}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ls := NewListSet[int]()
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

			// Verify order after removal
			for i := 0; i < ls.Size(); i++ {
				val, _ := ls.arr.Get(i)
				if val != c.expectOrder[i] {
					t.Errorf("At index %d: got %d, want %d", i, val, c.expectOrder[i])
				}
			}
		})
	}
}

func TestListSet_Contains(t *testing.T) {
	cases := []struct {
		name          string
		elements      []int
		searchElement int
		expectFound   bool
	}{
		{"contains existing element", []int{1, 2, 3}, 2, true},
		{"contains first element", []int{1, 2, 3}, 1, true},
		{"contains last element", []int{1, 2, 3}, 3, true},
		{"does not contain element", []int{1, 2, 3}, 4, false},
		{"empty set", []int{}, 1, false},
		{"single element found", []int{1}, 1, true},
		{"single element not found", []int{1}, 2, false},
		{"large set", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ls := NewListSet[int]()
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

func TestListSet_EdgeCases(t *testing.T) {
	t.Run("add and remove multiple times", func(t *testing.T) {
		ls := NewListSet[int]()

		// Add elements
		ls.Add(3)
		ls.Add(1)
		ls.Add(2)

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
		ls := NewListSet[int]()
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

	t.Run("maintain sorted order", func(t *testing.T) {
		ls := NewListSet[int]()
		elements := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}

		for _, elem := range elements {
			ls.Add(elem)
		}

		expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		for i := 0; i < ls.Size(); i++ {
			val, _ := ls.arr.Get(i)
			if val != expected[i] {
				t.Errorf("At index %d: got %d, want %d", i, val, expected[i])
			}
		}
	})
}

func BenchmarkListSet_Add(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ls := NewListSet[int]()
		for i := 0; i < 1000; i++ {
			ls.Add(i)
		}
	}
}

func BenchmarkListSet_Contains(b *testing.B) {
	ls := NewListSet[int]()
	for i := 0; i < 1000; i++ {
		ls.Add(i)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ls.Contains(n % 1000)
	}
}

func BenchmarkListSet_Remove(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ls := NewListSet[int]()
		for i := 0; i < 1000; i++ {
			ls.Add(i)
		}
		for i := 0; i < 1000; i++ {
			ls.Remove(i)
		}
	}
}

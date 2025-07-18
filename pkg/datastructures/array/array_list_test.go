package array

import (
	"testing"
)

func TestNewArrayList(t *testing.T) {
	al := NewArrayList[int]()
	if al.size != 0 {
		t.Errorf("Expected size to be 0, but got %d", al.size)
	}
	if !al.IsEmpty() {
		t.Errorf("Expected new list to be empty")
	}
	if len(al.arr) != 10 {
		t.Errorf("Expected initial capacity to be 10, but got %d", len(al.arr))
	}
}

func TestArrayList_Add(t *testing.T) {
	al := NewArrayList[int]()

	// Add at the end
	err := al.Add(0, 100)
	if err != nil {
		t.Errorf("Did not expect an error when adding to an empty list, but got %v", err)
	}
	if al.Size() != 1 {
		t.Errorf("Expected size to be 1, but got %d", al.Size())
	}
	val, _ := al.Get(0)
	if val != 100 {
		t.Errorf("Expected to get 100, but got %v", val)
	}

	// Add at a specific index
	err = al.Add(0, 99)
	if err != nil {
		t.Errorf("Did not expect an error when adding at the beginning, but got %v", err)
	}
	if al.Size() != 2 {
		t.Errorf("Expected size to be 2, but got %d", al.Size())
	}
	val, _ = al.Get(0)
	if val != 99 {
		t.Errorf("Expected to get 99 at index 0, but got %v", val)
	}
	val, _ = al.Get(1)
	if val != 100 {
		t.Errorf("Expected to get 100 at index 1, but got %v", val)
	}

	// Add in the middle
	err = al.Add(1, 98)
	if err != nil {
		t.Errorf("Did not expect an error when adding in the middle, but got %v", err)
	}
	if al.Size() != 3 {
		t.Errorf("Expected size to be 3, but got %d", al.Size())
	}
	val, _ = al.Get(1)
	if val != 98 {
		t.Errorf("Expected to get 98 at index 1, but got %v", val)
	}

	// Add out of bounds
	err = al.Add(-1, 0)
	if err == nil {
		t.Errorf("Expected an error for negative index, but got nil")
	}
	err = al.Add(al.Size()+1, 0)
	if err == nil {
		t.Errorf("Expected an error for index > size, but got nil")
	}
}

func TestArrayList_Get(t *testing.T) {
	al := NewArrayList[int]()
	al.Add(0, 10)
	al.Add(1, 20)

	// Get valid index
	val, err := al.Get(1)
	if err != nil {
		t.Errorf("Did not expect an error, but got %v", err)
	}
	if val != 20 {
		t.Errorf("Expected to get 20, but got %v", val)
	}

	// Get out of bounds
	_, err = al.Get(-1)
	if err == nil {
		t.Errorf("Expected an error for negative index, but got nil")
	}
	_, err = al.Get(al.Size())
	if err == nil {
		t.Errorf("Expected an error for index >= size, but got nil")
	}
}

func TestArrayList_Set(t *testing.T) {
	al := NewArrayList[int]()
	al.Add(0, 10)
	al.Add(1, 20)

	// Set valid index
	err := al.Set(1, 25)
	if err != nil {
		t.Errorf("Did not expect an error, but got %v", err)
	}
	val, _ := al.Get(1)
	if val != 25 {
		t.Errorf("Expected to get 25, but got %v", val)
	}

	// Set out of bounds
	err = al.Set(-1, 0)
	if err == nil {
		t.Errorf("Expected an error for negative index, but got nil")
	}
	err = al.Set(al.Size(), 0)
	if err == nil {
		t.Errorf("Expected an error for index >= size, but got nil")
	}
}

func TestArrayList_Remove(t *testing.T) {
	al := NewArrayList[int]()
	al.Add(0, 10)
	al.Add(1, 20)
	al.Add(2, 30)

	// Remove from middle
	removed, err := al.Remove(1)
	if err != nil {
		t.Errorf("Did not expect an error, but got %v", err)
	}
	if removed != 20 {
		t.Errorf("Expected to remove 20, but got %v", removed)
	}
	if al.Size() != 2 {
		t.Errorf("Expected size to be 2, but got %d", al.Size())
	}
	val, _ := al.Get(1)
	if val != 30 {
		t.Errorf("Expected element at index 1 to be 30, but got %v", val)
	}

	// Remove from beginning
	removed, err = al.Remove(0)
	if err != nil {
		t.Errorf("Did not expect an error, but got %v", err)
	}
	if removed != 10 {
		t.Errorf("Expected to remove 10, but got %v", removed)
	}
	if al.Size() != 1 {
		t.Errorf("Expected size to be 1, but got %d", al.Size())
	}

	// Remove from end
	al.Add(1, 40)
	removed, err = al.Remove(al.Size() - 1)
	if err != nil {
		t.Errorf("Did not expect an error, but got %v", err)
	}
	if removed != 40 {
		t.Errorf("Expected to remove 40, but got %v", removed)
	}

	// Remove out of bounds
	_, err = al.Remove(-1)
	if err == nil {
		t.Errorf("Expected an error for negative index, but got nil")
	}
	_, err = al.Remove(al.Size())
	if err == nil {
		t.Errorf("Expected an error for index >= size, but got nil")
	}
}

func TestArrayList_ResizeAndShrink(t *testing.T) {
	al := NewArrayList[int]()

	// 1. Add 21 elements to trigger two resizes.
	// 0 -> 10 elements (cap 10)
	// 11th element -> resize to 20
	// 21st element -> resize to 40
	for i := 0; i < 21; i++ {
		al.Add(i, i)
	}

	if al.Size() != 21 {
		t.Errorf("Expected size to be 21, but got %d", al.Size())
	}
	if len(al.arr) != 40 {
		t.Errorf("Expected capacity to be 40 after resizing, but got %d", len(al.arr))
	}

	// 2. Remove elements until shrink is triggered.
	// Shrink condition: size * 3 <= capacity (40)
	// This is true when size <= 13.
	// We need to remove 21 - 13 = 8 elements.
	for i := 0; i < 8; i++ {
		al.Remove(0) // Removing from the start
	}

	// After 8 removals, size is 13. The shrink condition (13 * 3 <= 40) is met.
	// The capacity should now be halved from 40 to 20.
	if al.Size() != 13 {
		t.Errorf("Expected size to be 13, but got %d", al.Size())
	}
	if len(al.arr) != 20 {
		t.Errorf("Expected capacity to be 20 after shrinking, but got %d", len(al.arr))
	}
}

func TestArrayList_String(t *testing.T) {
	al := NewArrayList[string]()
	expectedEmpty := "ArrayList: [] (Size: 0)"
	if al.String() != expectedEmpty {
		t.Errorf("Expected '%s', but got '%s'", expectedEmpty, al.String())
	}

	al.Add(0, "10")
	al.Add(1, "hello")
	expectedPopulated := "ArrayList: [10, hello] (Size: 2)"
	if al.String() != expectedPopulated {
		t.Errorf("Expected '%s', but got '%s'", expectedPopulated, al.String())
	}
}

func BenchmarkArrayList_Add(b *testing.B) {
	for n := 0; n < b.N; n++ {
		al := NewArrayList[int]()
		for i := 0; i < 1000; i++ {
			al.Add(al.Size(), i)
		}
	}
}

func BenchmarkArrayList_Get(b *testing.B) {
	al := NewArrayList[int]()
	for i := 0; i < 1000; i++ {
		al.Add(al.Size(), i)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = al.Get(n % 1000)
	}
}

func BenchmarkArrayList_Set(b *testing.B) {
	al := NewArrayList[int]()
	for i := 0; i < 1000; i++ {
		al.Add(al.Size(), i)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		al.Set(n % 1000, n)
	}
}

func BenchmarkArrayList_Remove(b *testing.B) {
	for n := 0; n < b.N; n++ {
		al := NewArrayList[int]()
		for i := 0; i < 1000; i++ {
			al.Add(al.Size(), i)
		}
		for i := 0; i < 1000; i++ {
			al.Remove(0)
		}
	}
}
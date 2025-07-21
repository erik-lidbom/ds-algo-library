package heap

import (
	"ds-algorithms/pkg/datastructures/array"
	"testing"
)

func TestNewMinHeap(t *testing.T) {
	mh := NewMinHeap[int]()
	if mh.Size() != 0 {
		t.Errorf("Expected size to be 0, but got %d", mh.Size())
	}
	if !mh.IsEmpty() {
		t.Errorf("Expected new heap to be empty")
	}
}

func TestMinHeap_Add(t *testing.T) {
	mh := NewMinHeap[int]()

	mh.Add(10)
	if mh.Size() != 1 {
		t.Errorf("Expected size to be 1, but got %d", mh.Size())
	}
	if mh.IsEmpty() {
		t.Errorf("Expected heap to not be empty after adding element")
	}
	
	min, err := mh.GetMin()
	if err != nil {
		t.Errorf("Did not expect an error when getting min, but got %v", err)
	}
	if min != 10 {
		t.Errorf("Expected min to be 10, but got %v", min)
	}

	mh.Add(5)
	mh.Add(15)
	mh.Add(3)
	mh.Add(12)
	
	if mh.Size() != 5 {
		t.Errorf("Expected size to be 5, but got %d", mh.Size())
	}
	
	min, err = mh.GetMin()
	if err != nil {
		t.Errorf("Did not expect an error when getting min, but got %v", err)
	}
	if min != 3 {
		t.Errorf("Expected min to be 3, but got %v", min)
	}
}

func TestMinHeap_GetMin(t *testing.T) {
	mh := NewMinHeap[int]()

	_, err := mh.GetMin()
	if err == nil {
		t.Errorf("Expected an error when getting min from empty heap, but got nil")
	}

	mh.Add(20)
	mh.Add(10)
	mh.Add(30)
	
	min, err := mh.GetMin()
	if err != nil {
		t.Errorf("Did not expect an error when getting min, but got %v", err)
	}
	if min != 10 {
		t.Errorf("Expected min to be 10, but got %v", min)
	}
	
	if mh.Size() != 3 {
		t.Errorf("Expected size to remain 3 after GetMin, but got %d", mh.Size())
	}
}

func TestMinHeap_RemoveMin(t *testing.T) {
	mh := NewMinHeap[int]()

	_, err := mh.RemoveMin()
	if err == nil {
		t.Errorf("Expected an error when removing from empty heap, but got nil")
	}

	elements := []int{15, 10, 20, 8, 25, 5, 7}
	for _, elem := range elements {
		mh.Add(elem)
	}

	expectedOrder := []int{5, 7, 8, 10, 15, 20, 25}
	for i, expected := range expectedOrder {
		removed, err := mh.RemoveMin()
		if err != nil {
			t.Errorf("Did not expect an error when removing min, but got %v", err)
		}
		if removed != expected {
			t.Errorf("Expected to remove %d, but got %v", expected, removed)
		}
		if mh.Size() != len(elements)-i-1 {
			t.Errorf("Expected size to be %d, but got %d", len(elements)-i-1, mh.Size())
		}
	}

	if !mh.IsEmpty() {
		t.Errorf("Expected heap to be empty after removing all elements")
	}
}

func TestMinHeap_BuildHeap(t *testing.T) {
	al := array.NewArrayList[int]()
	elements := []int{15, 10, 20, 8, 25, 5, 7}
	for i, elem := range elements {
		al.Add(i, elem)
	}

	mh := NewMinHeap[int]()
	err := mh.BuildHeap(al)
	if err != nil {
		t.Errorf("Did not expect an error when building heap, but got %v", err)
	}

	if mh.Size() != len(elements) {
		t.Errorf("Expected size to be %d, but got %d", len(elements), mh.Size())
	}

	expectedOrder := []int{5, 7, 8, 10, 15, 20, 25}
	for _, expected := range expectedOrder {
		removed, err := mh.RemoveMin()
		if err != nil {
			t.Errorf("Did not expect an error when removing min, but got %v", err)
		}
		if removed != expected {
			t.Errorf("Expected to remove %d, but got %v", expected, removed)
		}
	}
}

func TestMinHeap_HeapPropertyMaintained(t *testing.T) {
	mh := NewMinHeap[int]()
	
	elements := []int{50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45}
	for _, elem := range elements {
		mh.Add(elem)
		
		min, err := mh.GetMin()
		if err != nil {
			t.Errorf("Did not expect an error when getting min, but got %v", err)
		}
		
		for i := 1; i < mh.Size(); i++ {
			val, err := mh.heap.Get(i)
			if err != nil {
				t.Errorf("Error getting element at index %d: %v", i, err)
			}
			if val < min {
				t.Errorf("Heap property violated: found element %v smaller than min %v", val, min)
			}
		}
	}
}

func TestMinHeap_WithStrings(t *testing.T) {
	mh := NewMinHeap[string]()
	
	words := []string{"zebra", "apple", "banana", "cherry"}
	for _, word := range words {
		mh.Add(word)
	}
	
	min, err := mh.GetMin()
	if err != nil {
		t.Errorf("Did not expect an error when getting min, but got %v", err)
	}
	if min != "apple" {
		t.Errorf("Expected min to be 'apple', but got %v", min)
	}
	
	expected := []string{"apple", "banana", "cherry", "zebra"}
	for _, exp := range expected {
		removed, err := mh.RemoveMin()
		if err != nil {
			t.Errorf("Did not expect an error when removing min, but got %v", err)
		}
		if removed != exp {
			t.Errorf("Expected to remove %s, but got %v", exp, removed)
		}
	}
}

func TestMinHeap_SingleElement(t *testing.T) {
	mh := NewMinHeap[int]()
	
	mh.Add(42)
	if mh.Size() != 1 {
		t.Errorf("Expected size to be 1, but got %d", mh.Size())
	}
	
	min, err := mh.GetMin()
	if err != nil {
		t.Errorf("Did not expect an error when getting min, but got %v", err)
	}
	if min != 42 {
		t.Errorf("Expected min to be 42, but got %v", min)
	}
	
	removed, err := mh.RemoveMin()
	if err != nil {
		t.Errorf("Did not expect an error when removing min, but got %v", err)
	}
	if removed != 42 {
		t.Errorf("Expected to remove 42, but got %v", removed)
	}
	
	if !mh.IsEmpty() {
		t.Errorf("Expected heap to be empty after removing single element")
	}
}

func BenchmarkMinHeap_Add(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mh := NewMinHeap[int]()
		for i := 0; i < 1000; i++ {
			mh.Add(i)
		}
	}
}

func BenchmarkMinHeap_RemoveMin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mh := NewMinHeap[int]()
		for i := 0; i < 1000; i++ {
			mh.Add(i)
		}
		b.StartTimer()
		for i := 0; i < 1000; i++ {
			mh.RemoveMin()
		}
		b.StopTimer()
	}
}

func BenchmarkMinHeap_BuildHeap(b *testing.B) {
	al := array.NewArrayList[int]()
	for i := 0; i < 1000; i++ {
		al.Add(i, 1000-i)
	}
	
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		mh := NewMinHeap[int]()
		mh.BuildHeap(al)
	}
} 
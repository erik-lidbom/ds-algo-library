package heap

import (
	"ds-algorithms/pkg/datastructures/array"
	"testing"
)

func TestNewMaxHeap(t *testing.T) {
	mh := NewMaxHeap[int]()
	if mh.Size() != 0 {
		t.Errorf("Expected size to be 0, but got %d", mh.Size())
	}
	if !mh.IsEmpty() {
		t.Errorf("Expected new heap to be empty")
	}
}

func TestMaxHeap_Add(t *testing.T) {
	mh := NewMaxHeap[int]()

	mh.Add(10)
	if mh.Size() != 1 {
		t.Errorf("Expected size to be 1, but got %d", mh.Size())
	}
	if mh.IsEmpty() {
		t.Errorf("Expected heap to not be empty after adding element")
	}
	
	max, err := mh.GetMax()
	if err != nil {
		t.Errorf("Did not expect an error when getting max, but got %v", err)
	}
	if max != 10 {
		t.Errorf("Expected max to be 10, but got %v", max)
	}

	mh.Add(15)
	mh.Add(5)
	mh.Add(20)
	mh.Add(12)
	
	if mh.Size() != 5 {
		t.Errorf("Expected size to be 5, but got %d", mh.Size())
	}
	
	max, err = mh.GetMax()
	if err != nil {
		t.Errorf("Did not expect an error when getting max, but got %v", err)
	}
	if max != 20 {
		t.Errorf("Expected max to be 20, but got %v", max)
	}
}

func TestMaxHeap_GetMax(t *testing.T) {
	mh := NewMaxHeap[int]()

	_, err := mh.GetMax()
	if err == nil {
		t.Errorf("Expected an error when getting max from empty heap, but got nil")
	}

	mh.Add(20)
	mh.Add(30)
	mh.Add(10)
	
	max, err := mh.GetMax()
	if err != nil {
		t.Errorf("Did not expect an error when getting max, but got %v", err)
	}
	if max != 30 {
		t.Errorf("Expected max to be 30, but got %v", max)
	}
	
	if mh.Size() != 3 {
		t.Errorf("Expected size to remain 3 after GetMax, but got %d", mh.Size())
	}
}

func TestMaxHeap_RemoveMax(t *testing.T) {
	mh := NewMaxHeap[int]()

	_, err := mh.RemoveMax()
	if err == nil {
		t.Errorf("Expected an error when removing from empty heap, but got nil")
	}

	elements := []int{15, 10, 20, 8, 25, 5, 7}
	for _, elem := range elements {
		mh.Add(elem)
	}

	expectedOrder := []int{25, 20, 15, 10, 8, 7, 5}

	for i, expected := range expectedOrder {
	
		removed, err := mh.RemoveMax()




		if err != nil {
			t.Errorf("Did not expect an error when removing max, but got %v", err)
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

func TestMaxHeap_BuildHeap(t *testing.T) {
	al := array.NewArrayList[int]()
	elements := []int{15, 10, 20, 8, 25, 5, 7}
	for i, elem := range elements {
		al.Add(i, elem)
	}

	mh := NewMaxHeap[int]()
	err := mh.BuildHeap(al)
	if err != nil {
		t.Errorf("Did not expect an error when building heap, but got %v", err)
	}

	if mh.Size() != len(elements) {
		t.Errorf("Expected size to be %d, but got %d", len(elements), mh.Size())
	}

	expectedOrder := []int{25, 20, 15, 10, 8, 7, 5}
	for _, expected := range expectedOrder {
		removed, err := mh.RemoveMax()
		if err != nil {
			t.Errorf("Did not expect an error when removing max, but got %v", err)
		}
		if removed != expected {
			t.Errorf("Expected to remove %d, but got %v", expected, removed)
		}
	}
}

func TestMaxHeap_HeapPropertyMaintained(t *testing.T) {
	mh := NewMaxHeap[int]()
	
	elements := []int{50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45}
	for _, elem := range elements {
		mh.Add(elem)
		
		max, err := mh.GetMax()
		if err != nil {
			t.Errorf("Did not expect an error when getting max, but got %v", err)
		}
		
		for i := 1; i < mh.Size(); i++ {
			val, err := mh.heap.Get(i)
			if err != nil {
				t.Errorf("Error getting element at index %d: %v", i, err)
			}
			if val > max {
				t.Errorf("Heap property violated: found element %v larger than max %v", val, max)
			}
		}
	}
}

func TestMaxHeap_WithStrings(t *testing.T) {
	mh := NewMaxHeap[string]()
	
	words := []string{"apple", "zebra", "banana", "cherry"}
	for _, word := range words {
		mh.Add(word)
	}
	
	max, err := mh.GetMax()
	if err != nil {
		t.Errorf("Did not expect an error when getting max, but got %v", err)
	}
	if max != "zebra" {
		t.Errorf("Expected max to be 'zebra', but got %v", max)
	}
	
	expected := []string{"zebra", "cherry", "banana", "apple"}
	for _, exp := range expected {
		removed, err := mh.RemoveMax()
		if err != nil {
			t.Errorf("Did not expect an error when removing max, but got %v", err)
		}
		if removed != exp {
			t.Errorf("Expected to remove %s, but got %v", exp, removed)
		}
	}
}

func TestMaxHeap_SingleElement(t *testing.T) {
	mh := NewMaxHeap[int]()
	
	mh.Add(42)
	if mh.Size() != 1 {
		t.Errorf("Expected size to be 1, but got %d", mh.Size())
	}
	
	max, err := mh.GetMax()
	if err != nil {
		t.Errorf("Did not expect an error when getting max, but got %v", err)
	}
	if max != 42 {
		t.Errorf("Expected max to be 42, but got %v", max)
	}
	
	removed, err := mh.RemoveMax()
	if err != nil {
		t.Errorf("Did not expect an error when removing max, but got %v", err)
	}
	if removed != 42 {
		t.Errorf("Expected to remove 42, but got %v", removed)
	}
	
	if !mh.IsEmpty() {
		t.Errorf("Expected heap to be empty after removing single element")
	}
}

func TestMaxHeap_DuplicateElements(t *testing.T) {
	mh := NewMaxHeap[int]()
	
	elements := []int{10, 20, 10, 30, 20, 30}
	for _, elem := range elements {
		mh.Add(elem)
	}
	
	if mh.Size() != len(elements) {
		t.Errorf("Expected size to be %d, but got %d", len(elements), mh.Size())
	}
	
	expectedOrder := []int{30, 30, 20, 20, 10, 10}
	for _, expected := range expectedOrder {
		removed, err := mh.RemoveMax()
		if err != nil {
			t.Errorf("Did not expect an error when removing max, but got %v", err)
		}
		if removed != expected {
			t.Errorf("Expected to remove %d, but got %v", expected, removed)
		}
	}
}

func TestMaxHeap_LargeDataset(t *testing.T) {
	mh := NewMaxHeap[int]()
	
	for i := 1; i <= 100; i++ {
		mh.Add(i)
	}
	
	if mh.Size() != 100 {
		t.Errorf("Expected size to be 100, but got %d", mh.Size())
	}
	
	max, err := mh.GetMax()
	if err != nil {
		t.Errorf("Did not expect an error when getting max, but got %v", err)
	}
	if max != 100 {
		t.Errorf("Expected max to be 100, but got %v", max)
	}
	
	for i := 100; i > 90; i-- {
		removed, err := mh.RemoveMax()
		if err != nil {
			t.Errorf("Did not expect an error when removing max, but got %v", err)
		}
		if removed != i {
			t.Errorf("Expected to remove %d, but got %v", i, removed)
		}
	}
}

func BenchmarkMaxHeap_Add(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mh := NewMaxHeap[int]()
		for i := 0; i < 1000; i++ {
			mh.Add(i)
		}
	}
}

func BenchmarkMaxHeap_RemoveMax(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mh := NewMaxHeap[int]()
		for i := 0; i < 1000; i++ {
			mh.Add(i)
		}
		b.StartTimer()
		for i := 0; i < 1000; i++ {
			mh.RemoveMax()
		}
		b.StopTimer()
	}
}

func BenchmarkMaxHeap_BuildHeap(b *testing.B) {
	al := array.NewArrayList[int]()
	for i := 0; i < 1000; i++ {
		al.Add(i, i)
	}
	
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		mh := NewMaxHeap[int]()
		mh.BuildHeap(al)
	}
} 
package search

import (
	"testing"

	"ds-algorithms/pkg/datastructures/array"
	"ds-algorithms/pkg/datastructures/searchable"
)

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		name        string
		arr         []int
		elem        int
		expectIdx   int
		expectFound bool
	}{
		{"empty", []int{}, 1, 0, false},
		{"single found", []int{5}, 5, 0, true},
		{"single not found", []int{5}, 1, 0, false},
		{"found", []int{1, 2, 3, 4, 5}, 3, 2, true},
		{"not found", []int{1, 2, 3, 4, 5}, 9, 0, false},
	}
	for _, c := range cases {
		currArr := searchable.SearchableSlice[int](c.arr)
		idx, found := BinarySearch(currArr, c.elem)
		if found != c.expectFound || (found && idx != c.expectIdx) {
			t.Errorf("%s: got (idx=%d, found=%v), want (idx=%d, found=%v)", c.name, idx, found, c.expectIdx, c.expectFound)
		}
	}
}

func TestFindInsertionPoint(t *testing.T) {
	cases := []struct {
		name        string
		arr         []int
		elem        int
		expectIdx   int
		expectFound bool
	}{
		{"empty", []int{}, 1, 0, false},
		{"single found", []int{5}, 5, 0, true},
		{"single not found below", []int{5}, 1, 0, false},
		{"single not found above", []int{5}, 9, 1, false},
		{"found first", []int{1, 2, 3, 4, 5}, 1, 0, true},
		{"found middle", []int{1, 2, 3, 4, 5}, 3, 2, true},
		{"found last", []int{1, 2, 3, 4, 5}, 5, 4, true},
		{"insert at beginning", []int{2, 3, 4, 5}, 1, 0, false},
		{"insert at middle", []int{1, 2, 4, 5}, 3, 2, false},
		{"insert at end", []int{1, 2, 3, 4}, 5, 4, false},
		{"duplicate elements", []int{1, 2, 2, 3, 4}, 2, 1, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			arr := array.NewArrayList[int]()
			for i, val := range c.arr {
				arr.Add(i, val)
			}
			found, idx := FindInsertionPoint(arr, c.elem)
			if found != c.expectFound || idx != c.expectIdx {
				t.Errorf("name %s, got (found=%v, idx=%d), want (found=%v, idx=%d)", c.name, found, idx, c.expectFound, c.expectIdx)
			}
		})
	}
}

func TestFindUpperBound(t *testing.T) {
	cases := []struct {
		name        string
		arr         []int
		elem        int
		expectIdx   int
		expectFound bool
	}{
		{"empty", []int{}, 1, 0, false},
		{"single found", []int{5}, 5, 1, false},	
		{"single not found below", []int{5}, 1, 0, false},
		{"single not found above", []int{5}, 9, 1, false},
		{"found first", []int{1, 2, 3, 4, 5}, 1, 1, false},
		{"found middle", []int{1, 2, 3, 4, 5}, 3, 3, false},
		{"found last", []int{1, 2, 3, 4, 5}, 5, 5, false},
		{"insert at beginning", []int{2, 3, 4, 5}, 1, 0, false},
		{"insert at middle", []int{1, 2, 4, 5}, 3, 2, false},
		{"insert at end", []int{1, 2, 3, 4}, 5, 4, false},
		{"duplicate elements - first occurrence", []int{1, 2, 2, 3, 4}, 2, 3, false},
		{"duplicate elements - last occurrence", []int{1, 2, 2, 3, 4}, 2, 3, false},
		{"multiple duplicates", []int{1, 2, 2, 2, 3, 4}, 2, 4, false},
		{"all elements less than target", []int{1, 2, 3, 4}, 5, 4, false},
		{"all elements greater than target", []int{5, 6, 7, 8}, 3, 0, false},
		{"target between duplicates", []int{1, 2, 2, 2, 4, 5}, 3, 4, false},
		{"target equals max element", []int{1, 2, 3, 4, 5}, 5, 5, false},
		{"target less than min element", []int{5, 6, 7, 8}, 3, 0, false},
		{"consecutive numbers", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 5, false},
		{"consecutive numbers - not found", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 5, false},
		{"negative numbers", []int{-5, -3, -1, 0, 2, 4}, -2, 2, false},
		{"negative numbers - found", []int{-5, -3, -1, 0, 2, 4}, -3, 2, false},
		{"zero element", []int{0, 1, 2, 3}, 0, 1, false},
		{"zero target", []int{1, 2, 3, 4}, 0, 0, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			arr := array.NewArrayList[int]()
			for i, val := range c.arr {
				arr.Add(i, val)
			}
			found, idx := FindUpperBound(arr, c.elem)
			if found != c.expectFound || idx != c.expectIdx {
				t.Errorf("name %s, got (found=%v, idx=%d), want (found=%v, idx=%d)", c.name, found, idx, c.expectFound, c.expectIdx)
			}
		})
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	arr := make([]int, 1000)
	searchableArr := searchable.SearchableSlice[int](arr)
	for i := range arr {
		arr[i] = i
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		BinarySearch(searchableArr, 999)
	}
}

func BenchmarkFindInsertionPoint(b *testing.B) {
	arr := array.NewArrayList[int]()
	for i := 0; i < 1000; i++ {
		arr.Add(i, i)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		FindInsertionPoint(arr, 999)
	}
}

func BenchmarkFindUpperBound(b *testing.B) {
	arr := array.NewArrayList[int]()
	for i := 0; i < 1000; i++ {
		arr.Add(i, i)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		FindUpperBound(arr, 999)
	}
}

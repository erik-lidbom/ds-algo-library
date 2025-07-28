package search

import (
	"ds-algorithms/pkg/datastructures/array"
	"ds-algorithms/pkg/datastructures/searchable"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		name string
		arr []int
		elem int
		expectIdx int
		expectFound bool
	}{
		{"empty", []int{}, 1, 0, false},
		{"single found", []int{5}, 5, 0, true},
		{"single not found", []int{5}, 1, 0, false},
		{"found", []int{1,2,3,4,5}, 3, 2, true},
		{"not found", []int{1,2,3,4,5}, 9, 0, false},
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
		name string
		arr []int
		elem int
		expectIdx int
		expectFound bool
	}{
		{"empty", []int{}, 1, 0, false},
		{"single found", []int{5}, 5, 0, true},
		{"single not found below", []int{5}, 1, 0, false},
		{"single not found above", []int{5}, 9, 1, false},
		{"found first", []int{1,2,3,4,5}, 1, 0, true},
		{"found middle", []int{1,2,3,4,5}, 3, 2, true},
		{"found last", []int{1,2,3,4,5}, 5, 4, true},
		{"insert at beginning", []int{2,3,4,5}, 1, 0, false},
		{"insert at middle", []int{1,2,4,5}, 3, 2, false},
		{"insert at end", []int{1,2,3,4}, 5, 4, false},
		{"duplicate elements", []int{1,2,2,3,4}, 2, 1, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			arr := array.NewArrayList[int]()
			for i, val := range c.arr {
				arr.Add(i, val)
			}
			found, idx := FindInsertionPoint(arr, c.elem)
			if found != c.expectFound || idx != c.expectIdx {
				t.Errorf("name %s, got (found=%v, idx=%d), want (found=%v, idx=%d)",c.name, found, idx, c.expectFound, c.expectIdx)
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
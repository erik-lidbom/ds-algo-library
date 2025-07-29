package sorting

import (
	"reflect"
	"testing"

	"ds-algorithms/pkg/datastructures/searchable"
)

func TestBubbleSort(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect []int
	}{
		{"empty", []int{}, []int{}},
		{"single", []int{1}, []int{1}},
		{"sorted", []int{1, 2, 3}, []int{1, 2, 3}},
		{"reverse", []int{3, 2, 1}, []int{1, 2, 3}},
		{"random", []int{5, 1, 4, 2, 8}, []int{1, 2, 4, 5, 8}},
	}
	for _, c := range cases {
		arr := make([]int, len(c.input))
		copy(arr, c.input)
		currArr := searchable.SearchableSlice[int](arr)
		BubbleSort(currArr)
		if !reflect.DeepEqual(arr, c.expect) {
			t.Errorf("%s: got %v, want %v", c.name, arr, c.expect)
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arr := make([]int, 1000)
		for i := range arr {
			arr[i] = 1000 - i
		}
		currArr := searchable.SearchableSlice[int](arr)
		BubbleSort(currArr)
	}
}

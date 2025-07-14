package search

import "testing"

func TestLinearSearch(t *testing.T) {
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
		idx, found := LinearSearch(c.arr, c.elem)
		if found != c.expectFound || (found && idx != c.expectIdx) {
			t.Errorf("%s: got (idx=%d, found=%v), want (idx=%d, found=%v)", c.name, idx, found, c.expectIdx, c.expectFound)
		}
	}
}

func BenchmarkLinearSearch(b *testing.B) {
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = i
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		LinearSearch(arr, 999)
	}
} 
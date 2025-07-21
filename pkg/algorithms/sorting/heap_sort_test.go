package sorting

import (
	"ds-algorithms/pkg/datastructures/array"
	"testing"
)

func newArrayListFromInts(ints ...int) *array.ArrayList[int] {
	al := array.NewArrayList[int]()
	for i, v := range ints {
		al.Add(i, v)
	}
	return al
}

func arrayListEqual(a, b *array.ArrayList[int]) bool {
	if a.Size() != b.Size() {
		return false
	}
	for i := 0; i < a.Size(); i++ {
		x, _ := a.Get(i)
		y, _ := b.Get(i)
		if x != y {
			return false
		}
	}
	return true
}

func TestNaiveHeapsort(t *testing.T) {
	cases := []struct {
		name   string
		input  *array.ArrayList[int]
		expect *array.ArrayList[int]
	}{
		{"empty", newArrayListFromInts(), newArrayListFromInts()},
		{"single", newArrayListFromInts(1), newArrayListFromInts(1)},
		{"sorted", newArrayListFromInts(1,2,3), newArrayListFromInts(1,2,3)},
		{"reverse", newArrayListFromInts(3,2,1), newArrayListFromInts(1,2,3)},
		{"random", newArrayListFromInts(5,1,4,2,8), newArrayListFromInts(1,2,4,5,8)},
	}
	for _, c := range cases {
		al := newArrayListFromInts() // create a fresh list
		for i := 0; i < c.input.Size(); i++ {
			v, _ := c.input.Get(i)
			al.Add(i, v)
		}
		err := NaiveHeapsort(al)
		if err != nil {
			t.Errorf("%s: unexpected error: %v", c.name, err)
		}
		if !arrayListEqual(al, c.expect) {
			t.Errorf("%s: got %v, want %v", c.name, al, c.expect)
		}
	}
}

func TestHeapSort(t *testing.T) {
	cases := []struct {
		name   string
		input  *array.ArrayList[int]
		expect *array.ArrayList[int]
	}{
		{"empty", newArrayListFromInts(), newArrayListFromInts()},
		{"single", newArrayListFromInts(1), newArrayListFromInts(1)},
		{"sorted", newArrayListFromInts(1,2,3), newArrayListFromInts(1,2,3)},
		{"reverse", newArrayListFromInts(3,2,1), newArrayListFromInts(1,2,3)},
		{"random", newArrayListFromInts(5,1,4,2,8), newArrayListFromInts(1,2,4,5,8)},
	}
	for _, c := range cases {
		al := newArrayListFromInts() // create a fresh list
		for i := 0; i < c.input.Size(); i++ {
			v, _ := c.input.Get(i)
			al.Add(i, v)
		}
		err := HeapSort(al)
		if err != nil {
			t.Errorf("%s: unexpected error: %v", c.name, err)
		}
		if !arrayListEqual(al, c.expect) {
			t.Errorf("%s: got %v, want %v", c.name, al, c.expect)
		}
	}
}

func BenchmarkHeapSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		al := newArrayListFromInts()
		for i := 0; i < 1000; i++ {
			al.Add(i, 1000-i)
		}
		HeapSort(al)
	}
} 
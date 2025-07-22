package sets

import (
	"cmp"
	"ds-algorithms/pkg/datastructures/array"
	"errors"
	"fmt"
)

// TODO --> Add binary search for contain method. This requires that the ListSet always is sorted so logic has to be added for Add and Removal.

type ListSet[T cmp.Ordered] struct {
	arr *array.ArrayList[T]
	size int
}

func NewListSet[T cmp.Ordered]() *ListSet[T] {
	return &ListSet[T]{arr: array.NewArrayList[T](), size: 0}
}

func (ls *ListSet[T]) Size() int {
	return ls.size
}

func (ls *ListSet[T]) IsEmpty() bool {
	return ls.size == 0
}

func (ls *ListSet[T]) Add(elem T) error {
	containVal := ls.Contains(elem)

	if containVal {
		return fmt.Errorf("cannot add element to set: value %v already exists\n", containVal)
	}

	ls.arr.Add(ls.size, elem)
	ls.size++
	return nil
}

func (ls *ListSet[T]) Remove(elem T) (T, error) {
	var zero T
	for i := 0; i < ls.size - 1; i++ {
		currVal, err := ls.arr.Get(i)
		if err != nil {
			return zero, errors.New("cannot retrieve value from list")
		}

		if elem == currVal {
			ls.arr.Remove(i)
			return elem, nil
		}
	}
	return zero, fmt.Errorf("element %v not found in the set", elem)
}

func (ls *ListSet[T]) Contains(elem T) bool {

	for i := 0; i < ls.size - 1; i++ {
		currVal, err := ls.arr.Get(i)
		if err != nil {
			return false
		}

		if elem == currVal {
			return true
		}
	}
	return false
}
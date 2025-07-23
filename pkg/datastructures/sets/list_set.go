package sets

import (
	"cmp"
	"ds-algorithms/pkg/datastructures/array"
	"ds-algorithms/pkg/algorithms/searching"
	"fmt"
)

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
	duplicatedVal, index := search.FindInsertionPoint(ls.arr, elem)

	if duplicatedVal {
		return fmt.Errorf("cannot add element to set: value %v already exists\n", duplicatedVal)
	}

	ls.arr.Add(index, elem)
	ls.size++
	return nil
}

func (ls *ListSet[T]) Remove(elem T) (T, error) {
	var zero T

	containsElem, index := search.BinarySearchArrayList(ls.arr, elem)
	
	if containsElem {
		ls.arr.Remove(index)
		ls.size--
		return elem, nil
	}

	return zero, fmt.Errorf("element %v not found in the set", elem)
}


func (ls *ListSet[T]) Contains(elem T) bool {
	res, _ := search.BinarySearchArrayList(ls.arr, elem)
	return res
}

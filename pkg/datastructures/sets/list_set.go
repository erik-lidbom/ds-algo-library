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

	index, containsElem := search.BinarySearch(ls.arr, elem)
	
	if containsElem {
		ls.arr.Remove(index)
		ls.size--
		return elem, nil
	}

	return zero, fmt.Errorf("element %v not found in the set", elem)
}


func (ls *ListSet[T]) Contains(elem T) bool {
	_, res := search.BinarySearch(ls.arr, elem)
	return res
}

func (ls *ListSet[T]) First() (T, error) {
	var zero T
	if ls.IsEmpty() {
		return zero, fmt.Errorf("set is empty")
	}

	first, err := ls.arr.Get(0)
	if err != nil {
		return zero, fmt.Errorf("failed to retrieve first element: %w", err)
	}

	return first, nil
}

func (ls *ListSet[T]) Last() (T, error) {
	var zero T
	if ls.IsEmpty() {
		return zero, fmt.Errorf("set is empty")
	}

	last, err := ls.arr.Get(ls.Size() - 1)
	if err != nil {
		return zero, fmt.Errorf("failed to retrieve first element: %w", err)
	}

	return last, nil
}

func (ls *ListSet[T]) Floor(x T) (T, error) {
	var zero T
	if ls.IsEmpty() {
		return zero, fmt.Errorf("set is empty")
	}

	found, index := search.FindInsertionPoint(ls.arr, x)

	if found {
		floor, err := ls.arr.Get(index)
        if err != nil {
            return zero, fmt.Errorf("internal error: failed to retrieve floor element at index %d: %w", index, err)
        }
        return floor, nil
	}

	if index == 0 {
		return zero, fmt.Errorf("no element <= %v found in set", x)
	}

	floor, err := ls.arr.Get(index - 1)
    if err != nil {
        return zero, fmt.Errorf("internal error: failed to retrieve floor element at index %d-1: %w", index, err)
    }
	return floor, nil
}

func (ls *ListSet[T]) Ceiling(x T) (T, error) {
	var zero T
	if ls.IsEmpty() {
		return zero, fmt.Errorf("set is empty")
	}

	_, index := search.FindInsertionPoint(ls.arr, x)

	if index == ls.Size() {
		return zero, fmt.Errorf("no element >= %v found in set", x)
	}

	ceiling, err := ls.arr.Get(index)
	 if err != nil {
        return zero, fmt.Errorf("internal error: failed to retrieve ceiling element at index %d: %w", index, err)
    }

	return ceiling, nil
}
func (ls *ListSet[T]) Lower(x T) (T, error) {
	var zero T
	if ls.IsEmpty() {
		return zero, fmt.Errorf("set is empty")
	}

	_, index := search.FindInsertionPoint(ls.arr, x)

	if index == 0 {
		return zero, fmt.Errorf("no element < %v found in set", x)
	}

	lower, err := ls.arr.Get(index - 1)
	if err != nil {
		return zero, fmt.Errorf("internal error: failed to retrieve lower element at index %d: %w", index-1, err)
	}

	return lower, nil
}

func (ls *ListSet[T]) Higher(x T) (T, error) {
	var zero T
	if ls.IsEmpty() {
		return zero, fmt.Errorf("set is empty")
	}

	_, index := search.FindUpperBound(ls.arr, x)

	if index == ls.Size() {
		return zero, fmt.Errorf("no element > %v found in set", x)
	}

	higher, err := ls.arr.Get(index)
	if err != nil {
		return zero, fmt.Errorf("internal error: failed to retrieve higher element at index %d: %w", index-1, err)
	}

	return higher, nil
}

func (ls *ListSet[T]) Between(x, y T) *array.ArrayList[T]{

	arr := array.NewArrayList[T]()

	if ls.IsEmpty() {
		return arr
	}

	_, startIndex := search.FindInsertionPoint(ls.arr, x)
	_, endIndex := search.FindUpperBound(ls.arr, y)

	if startIndex >= ls.Size() || startIndex >= endIndex {
		return  arr
	}

	for i := startIndex; i < endIndex; i++ {
		currVal, _ := ls.arr.Get(i)
		arr.Add(i, currVal)
	}

	return arr
}
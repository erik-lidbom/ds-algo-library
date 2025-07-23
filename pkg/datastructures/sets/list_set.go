package sets

import (
	"cmp"
	"ds-algorithms/pkg/datastructures/array"
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
	duplicatedVal, index := ls.findInsertionPoint(elem)

	if duplicatedVal {
		return fmt.Errorf("cannot add element to set: value %v already exists\n", duplicatedVal)
	}

	ls.arr.Add(index, elem)
	ls.size++
	return nil
}

func (ls *ListSet[T]) Remove(elem T) (T, error) {
	var zero T

	duplicatedVal, index := ls.binary_search(elem)
	
	if duplicatedVal {
		ls.arr.Remove(index)
		ls.size--
		return elem, nil
	}

	return zero, fmt.Errorf("element %v not found in the set", elem)
}


func (ls *ListSet[T]) Contains(elem T) bool {
	res, _ := ls.binary_search(elem)
	return res
}

func (ls *ListSet[T]) binary_search(elem T) (bool, int) {
	
	leftIndex := 0
	rightIndex := ls.size - 1

	for leftIndex <= rightIndex {
		midIndex := leftIndex + (rightIndex - leftIndex) / 2
		currVal, _ := ls.arr.Get(midIndex)
		if elem == currVal {
		return true, midIndex
		} else if elem < currVal {
			rightIndex = midIndex - 1
		} else { 
			leftIndex = midIndex + 1
		}
	}
	return false, -1
}

func (ls *ListSet[T]) findInsertionPoint(elem T) (bool, int) {
    leftIndex := 0
    rightIndex := ls.size - 1 

    for leftIndex < rightIndex { 
        mid := leftIndex + (rightIndex - leftIndex) / 2
        currVal, _ := ls.arr.Get(mid)

        if currVal < elem {
            leftIndex = mid + 1
        } else {
            rightIndex = mid   
        }
    }
  
    if leftIndex < ls.size - 1 {
        lowerBoundVal, _ := ls.arr.Get(leftIndex)
        if lowerBoundVal == elem {
            return true, leftIndex
        }
    }

    return false, leftIndex
}

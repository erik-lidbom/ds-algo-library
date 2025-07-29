package heap

import (
	"cmp"
	"ds-algorithms/pkg/datastructures/array"
	"ds-algorithms/pkg/datastructures/common"
	"ds-algorithms/pkg/datastructures/searchable"
	"errors"
	"fmt"
)

type MaxHeap[T cmp.Ordered] struct {
	heap common.List[T]
	size int
}

func NewMaxHeap[T cmp.Ordered]() *MaxHeap[T] {
	return &MaxHeap[T]{
		heap: array.NewArrayList[T](),
		size: 0,
	}
}

func (mh *MaxHeap[T]) BuildHeap(data common.List[T]) error {
	mh.heap = data
	mh.size = data.Size()
	mid := getParent(mh.size - 1)

	for i := mid; i >= 0; i-- {
		err := mh.siftDown(i)

		if err != nil {
			return fmt.Errorf("build heap error: failed to siftDown at index %d: %w", i, err)
		}
	}

	return nil
}

func (mh *MaxHeap[T]) Size() int {
	return mh.size
}

func (mh *MaxHeap[T]) IsEmpty() bool {
	return mh.size == 0
}

func (mh *MaxHeap[T]) GetMax() (T, error) {
	var zero T
	if mh.size <= 0 {
		return zero, errors.New("heap is empty, cannot retrieve the maximum value")
	}
	return mh.heap.Get(0)
}

func (mh *MaxHeap[T]) Add(elem T) {
	// Since the ArrayList already handles resizing, we do not need to take that into consideration

	size := mh.heap.Size()
	mh.heap.Add(size, elem)
	mh.siftUp(mh.size)
	mh.size++	
}


func (mh *MaxHeap[T]) RemoveMax() (T, error) {
	var zero T

	if mh.size == 0 {
		return zero, errors.New("heap is empty, cannot remove the maximum value")
	}

	removedVal, err := mh.heap.Get(0)
	if err != nil {
		return zero, fmt.Errorf("failed to retrieve element for index %d\nerror: %w", 0, err)
	}

	
	
	swap_err := searchable.Swap(mh.heap, 0, mh.size - 1)
	if swap_err != nil {
		return zero, fmt.Errorf("failed to swap root with last element at index %d: %w", mh.size - 1, swap_err)
	}


	_, removeErr := mh.heap.Remove(mh.size - 1)
	if removeErr != nil {
		return zero, fmt.Errorf("failed to remove last element: %w", removeErr)
	}
	mh.size--

	if mh.size > 0 {
		mh.siftDown(0)
	}

	return removedVal, nil
}


func (mh *MaxHeap[T]) siftUp(pos int) error {
	for pos > 0 {
		parent := getParent(pos)

		newVal, err := mh.heap.Get(pos)
		if err != nil {
			return fmt.Errorf("failed to retrieve element for index %d\nerror: %w", pos, err)
		}
		parentVal, err := mh.heap.Get(parent)
		if err != nil {
			return fmt.Errorf("failed to retrieve element for index %d\nerror: %w", parent, err)
		}

		if newVal <= parentVal {
			return nil
		}
		searchable.Swap(mh.heap,pos, parent)

		pos = parent
	}
	return nil
}

func (mh *MaxHeap[T]) siftDown(pos int) error {

	for !isLeaf(pos, mh.size) {
		leftChild := getLeftChildIndex(pos)
		rightChild := getRightChildIndex(pos)

		currVal, err := mh.heap.Get(pos)
		if err != nil {
			return fmt.Errorf("failed to retrieve element for index %d\nerror: %w", pos, err)
		}

		leftVal, err := mh.heap.Get(leftChild)
		if err != nil {
			return fmt.Errorf("failed to retrieve element for index %d\nerror: %w", leftChild, err)
		}

		maxChild := leftChild
		maxVal := leftVal

		if rightChild < mh.size {
			rightVal, err := mh.heap.Get(rightChild)
			if err != nil {
				return fmt.Errorf("failed to retrieve element for index %d\nerror: %w", rightChild, err)
			}

			if rightVal > leftVal {
				maxChild = rightChild
				maxVal = rightVal
			}
		}

		if maxVal <= currVal {
			return nil
		}

		searchable.Swap(mh.heap, pos, maxChild)
		pos = maxChild
	}
	return nil
}

/*
This sorting function will invalidate the heap by setting size to -1. 
All heap operations will fail after calling Sort().
To reuse the heap again you have to call the BuildHeap function
*/

func (mh *MaxHeap[T]) Sort() (common.List[T], error) {
	if mh.size == 0 {
		return mh.heap, nil
	}

	originalSize := mh.Size()

	for i := originalSize - 1; i > 0; i-- {
		swap_err := searchable.Swap(mh.heap, 0, i)
		if swap_err != nil {
			return nil, fmt.Errorf("failed to swap root with last element at index %d: %w", mh.size - 1, swap_err)
		}
		mh.size--
		if mh.size > 0 {
			mh.siftDown(0)
		}
	}

	mh.size = -1

	return mh.heap, nil
}

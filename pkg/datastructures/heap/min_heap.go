package heap

import (
	"cmp"
	"ds-algorithms/pkg/datastructures/array"
	"errors"
	"fmt"
)

type MinHeap[T cmp.Ordered] struct {
	heap *array.ArrayList[T]
	size int
}

func NewMinHeap[T cmp.Ordered]() *MinHeap[T] {
	return &MinHeap[T]{
		heap: array.NewArrayList[T](),
		size: 0,
	}
}

func (mh *MinHeap[T]) BuildHeap(arr *array.ArrayList[T]) error {
	mh.heap = arr
	mh.size = arr.Size()
	mid := getParent(mh.size - 1)

	for i := mid; i >= 0; i-- {
		err := mh.siftDown(i)

		if err != nil {
            return fmt.Errorf("build heap error: failed to siftDown at index %d: %w", i, err)
		}
	}
	return nil
}

func (mh *MinHeap[T]) Size() int {
	return mh.size
}

func (mh *MinHeap[T]) IsEmpty() bool {
	return mh.size == 0
}

func (mh *MinHeap[T]) GetMin() (T, error) {
	var zero T
	if mh.size <= 0 {
		return zero, errors.New("heap is empty, cannot retrieve the minimum value")
	}

	return mh.heap.Get(0)
}

func (mh *MinHeap[T]) Add(elem T) {
	// Since the ArrayList already handles resizing, we do not need to take that into consideration
	size := mh.heap.Size()
	mh.heap.Add(size, elem)
	mh.siftUp(mh.size)
	mh.size++
}

func (mh *MinHeap[T]) RemoveMin() (T, error) {
	var zero T

	if mh.size == 0 {
		return zero, errors.New("heap is empty, cannot remove the maximum value")
	}


	removedVal, err := mh.heap.Get(0)
	if err != nil {
		return zero, fmt.Errorf("failed to retrieve element for index %d\nerror: %w", 0, err)
	}

	swap_err := array.Swap(mh.heap, 0, mh.size - 1)
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

func (mh *MinHeap[T]) swap(i, j int) error {
	iValue, err := mh.heap.Get(i)
	if err != nil {
		return fmt.Errorf("failed to retrieve element for index %d\nerror: %w", i, err)
	}

	jValue, err := mh.heap.Get(j)
	if err != nil {
		return fmt.Errorf("failed to retrieve element for index %d\nerror: %w", j, err)
	}

	err = mh.heap.Set(j, iValue)
	if err != nil {
		return fmt.Errorf("failed to swap element: %w", err)
	}

	err = mh.heap.Set(i, jValue) 
	if err != nil {
		// Since the first swap worked as expected, we need to do a rollback.
		rollbackErr := mh.heap.Set(j, jValue)
		if rollbackErr != nil {
			return fmt.Errorf("critical swap error: failed to set element at index %d (original error: %w), AND rollback for index %d failed: %w", j, err, i, rollbackErr)
		}
		return fmt.Errorf("failed to swap element: %w", err) 
	}

	return nil
}

func (mh *MinHeap[T]) siftUp(pos int) error {
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

		if newVal >= parentVal {
			return nil
		}
		mh.swap(pos, parent)
		pos = parent
	}
	return nil
}

func (mh *MinHeap[T]) siftDown(pos int) error {

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

		minChild := leftChild
		minVal := leftVal

		if rightChild < mh.size {
			rightVal, err := mh.heap.Get(rightChild)
			if err != nil {
				return fmt.Errorf("failed to retrieve element for index %d\nerror: %w", rightChild, err)
			}

			if rightVal < leftVal {
				minChild = rightChild
				minVal = rightVal
			}
		}

		if minVal >= currVal {
			return nil
		}

		mh.swap(pos, minChild)
		pos = minChild
	}
	return nil
}
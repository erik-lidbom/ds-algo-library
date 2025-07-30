package sorting

import (
	"cmp"
	"fmt"

	"ds-algorithms/pkg/datastructures/common"
	"ds-algorithms/pkg/datastructures/heap"
)

// todo: add comments
func NaiveHeapsort[T cmp.Ordered](arr common.
	List[T],
) error {
	heap := heap.NewMinHeap[T]()
	for i := 0; i < arr.Size(); i++ {
		elem, err := arr.Get(i)
		if err != nil {
			return fmt.Errorf("failed to retrieve element for index %d\nerror: %w", i, err)
		}

		heap.Add(elem)
	}

	for i := 0; i < arr.Size(); i++ {
		sortedVal, err := heap.RemoveMin()
		if err != nil {
			return fmt.Errorf("failed to remove element for index %d\nerror: %w", i, err)
		}
		arr.Set(i, sortedVal)
	}
	return nil
}

func HeapSort[T cmp.Ordered](arr common.List[T]) error {
	maxHeap := heap.NewMaxHeap[T]()
	maxHeap.BuildHeap(arr)

	_, err := maxHeap.Sort()
	if err != nil {
		return fmt.Errorf("failed to sort the array\nerror: %w", err)
	}
	return nil
}

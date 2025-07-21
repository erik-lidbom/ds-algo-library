package sorting

import (
	"cmp"
	"ds-algorithms/pkg/datastructures/array"
	"ds-algorithms/pkg/datastructures/heap"
	"fmt"
)

func NaiveHeapsort[T cmp.Ordered](arr *array.ArrayList[T]) error {

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

func HeapSort[T cmp.Ordered](arr *array.ArrayList[T]) error {
	max_heap := heap.NewMaxHeap[T]()
	max_heap.BuildHeap(arr)

	
	_, err := max_heap.Sort()
	if err != nil {
		return fmt.Errorf("failed to sort the array\nerror: %w", err)
	}
	return nil
}
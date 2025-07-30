package sorting

import (
	"cmp"

	"ds-algorithms/pkg/datastructures/common"
)

// SelectionSort function sorts a list on a Searchable compatible data structure.
//
// Description:
// Selection Sort is a sorting algorithm that works by
// repeatedly finding the minimum element
// from the unsorted part of the list and putting it at the beginning of
// the unsorted segment.
//
// The algorithm maintains two sub-arrays in a given array:
// - The sorted sub-array.
// - The unsorted sub-array.
// In each iteration of Selection Sort, the minimum element from the unsorted
// sub-array is picked and moved to the sorted sub-array.
//
// Time Complexity: O(N^2)
// Space Complexity: O(1) - since we do the operations in-place.
//
// Parameters:
//
//	arr  common.Searchable[T]: Unordered list
//
// Example:
//
//	myList := common.NewArrayList([]int{3, 1, 5, 4, 2})
//	sorting.SelectionSort(myList)
//
// Output:
//
//	Before sorting: [3 1 5 4 2]
//	After sorting:  [1 2 3 4 5]
func SelectionSort[T cmp.Ordered](arr common.Searchable[T]) {
	size := arr.Size()

	for i := 0; i < size-1; i++ {
		minIndex := i

		for j := i + 1; j < size; j++ {
			currVal, _ := arr.Get(j)
			minVal, _ := arr.Get(minIndex)
			if currVal < minVal {
				minIndex = j
			}
		}
		currVal, _ := arr.Get(i)
		minVal, _ := arr.Get(minIndex)
		arr.Set(i, minVal)
		arr.Set(minIndex, currVal)
	}
}

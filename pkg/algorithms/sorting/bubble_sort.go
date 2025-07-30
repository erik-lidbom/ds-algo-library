package sorting

import (
	"cmp"

	"ds-algorithms/pkg/datastructures/common"
)

// BubbleSort function sorts a list on Searchable a compatible data structure
// Description:
// Bubble sort retrieves a list and compare adjacent elements, if the elements are not in the correct order, a swap is performed.
// Every iteration moves the largest unsorted element to its correct position, which mean for every subsequent iteration shrinks the comparison area by 1.
//
// Time Complexity: O(N^2) - since we perform nested iterations of N*N comparisons.
// Space Complexity: O(1) - since we do the operations in-place.
//
// Parameters:
//
//	arr common.Searchable[T]: Unordered list
//
// Example:
//
//	myList := common.NewArrayList([]int{3, 1, 5, 4, 2})
//
//	sorting.BubbleSort(myList)
//
// Output:
//
//	Before sorting: [3 1 5 4 2]
//	After sorting: [1 2 3 4 5]
func BubbleSort[T cmp.Ordered](arr common.Searchable[T]) {
	size := arr.Size()

	for i := 0; i < size-1; i++ {
		for j := 0; j < size-i-1; j++ {
			valX, _ := arr.Get(j)
			valY, _ := arr.Get(j + 1)

			if valX > valY {
				arr.Set(j, valY)
				arr.Set(j+1, valX)
			}
		}
	}
}

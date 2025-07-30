package sorting

import (
	"cmp"

	"ds-algorithms/pkg/datastructures/common"
)

// InsertionSort function sorts a list on Searchable a compatible data structure
//
// Description:
//
// Insertion sort retrieves a list and sorts from left to right. For each element, the current elemet compares with the elements to the left, if the element is smaller than preceding it performs a swap and continue comparing until the correct position is found.
//
// For every iteration the comparison area increments by one.
//
// Time Complexity: O(N^2) - since we in worst case perform nested iterations of N*N comparisons.
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
//	sorting.InsertionSort(myList)
//
// Output:
//
//	Before sorting: [3 1 5 4 2]
//	After sorting: [1 2 3 4 5]
func InsertionSort[T cmp.Ordered](
	arr common.Searchable[T],
) {
	size := arr.Size()

	for i := 1; i < size; i++ {
		j := i

		valX, _ := arr.Get(j)

		for j > 0 {
			valY, _ := arr.Get(j - 1)
			if valX >= valY {
				break
			}
			arr.Set(j, valY)
			j--
		}
		arr.Set(j, valX)
	}
}

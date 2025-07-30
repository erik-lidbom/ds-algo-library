package sorting

import (
	"cmp"

	"ds-algorithms/pkg/datastructures/common"
)

// QuickSort sorts a list in-place using the Hoare partition scheme.
//
// Description:
//
// QuickSort is a divide-and-conquer sorting algorithm. It works by selecting a 'pivot' element
// from the list and partitioning the other elements into two sub-arrays, according to whether they are
// less than or greater than the pivot. The sub-arrays are then sorted recursively. This implementation
// uses the Hoare partition scheme.
//
// The algorithm recursively sorts the sublists to the left and right of the partition index.
// The process continues until the sublists are of length zero or one, at which point the list is sorted.
//
// Time Complexity: O(N log N) on average, O(N^2) worst case
// Space Complexity: O(log N) due to recursion stack
//
// Parameters:
//
//	arr  common.Searchable[T]: Unordered list
//	left int: The starting index of the sublist
//	right int: The ending index of the sublist
//
// Example:
//
//	myList := common.NewArrayList([]int{3, 1, 5, 4, 2})
//	sorting.QuickSort(myList, 0, myList.Size()-1)
//
// Output:
//
//	Before sorting: [3 1 5 4 2]
//	After sorting:  [1 2 3 4 5]
func QuickSort[T cmp.Ordered](arr common.Searchable[T], left int, right int) {
	if left >= right { // Base case for recursion
		return
	}

	// Optimization technique, for smaller list an insertion sort is more efficient than QuickSort
	if arr.Size() <= 100 {
		InsertionSort(arr)
		return
	}

	pivot := findPivot(left, right)
	pivot = partition(arr, left, right, pivot)
	QuickSort(arr, left, pivot)
	QuickSort(arr, pivot+1, right)
}

// findPivot returns the pivot and currently the middle value will always be selected as a pivot
// Parameters:
//
//	left int: The starting index of the sublist
//	right int: The ending index of the sublist
func findPivot(left int, right int) int {
	return (left + right) / 2
}

// partition rearranges elements around a pivot using Hoare's Partition Scheme.
//
// Parameters:
//
//	arr common.Searchable[T]: The list to be partitioned.
//	left int: The starting index of the sublist
//	right int: The ending index of the sublist
//	pivot int: The index of the chosen pivot element.
//
// Returns:
//
//	int: An index separating the two partitions; elements to its left are
//	  less than or equal to the pivot, and elements to its right are greater
//	  than or equal.
func partition[T cmp.Ordered](arr common.Searchable[T], left int, right int, pivot int) int {
	value, _ := arr.Get(pivot)
	pivotVal := value

	for {
		for left <= right {
			leftVal, _ := arr.Get(left)
			if leftVal >= pivotVal {
				break
			}
			left++
		}

		for left <= right {
			rightVal, _ := arr.Get(right)
			if rightVal <= pivotVal {
				break
			}
			right--
		}

		if left >= right {
			return right
		}

		leftVal, _ := arr.Get(left)
		rightVal, _ := arr.Get(right)
		arr.Set(left, rightVal)
		arr.Set(right, leftVal)
		left++
		right--
	}
}

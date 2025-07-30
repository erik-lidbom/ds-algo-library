package sorting

import (
	"cmp"

	"ds-algorithms/pkg/datastructures/common"
)

// MergeSort function sorts a list on Searchable a compatible data structure
//
// Description:
//
// Merge Sort is a divide-and-conquer algorithm. It works by recursively dividing
// the unsorted list into n sublists, each containing one element (a list of one
// element is considered sorted). Then, it repeatedly merges sublists to produce
// new sorted sublists until there is only one sorted list remaining.
//
// The core idea is that merging two already sorted sublists into a single sorted
// list is more efficient than sorting the whole list at once.
//
// Time Complexity: O(N log N)
// Space Complexity: O(N) - since an auxiliary array of size N if needed for the merge process.
//
// Parameters:
//
//	arr common.Searchable[T]: Unordered list
//	left int: The starting index of the sublist
//	right int: The ending index of the sublist
//
// Example:
//
//	myList := common.NewArrayList([]int{3, 1, 5, 4, 2})
//
//	sorting.MergeSort(myList, 0, myList.Size() - 1)
//
// Output:
//
//	Before sorting: [3 1 5 4 2]
//	After sorting: [1 2 3 4 5]
func MergeSort[T cmp.Ordered](arr common.Searchable[T], left int, right int) {
	if left >= right { // Base case for recursion
		return
	}

	// Recursively splits the list into sublists and lastly performs a merge
	mid := (left + right) / 2    // Calculates the midpoint
	MergeSort(arr, left, mid)    // Sorts the left half
	MergeSort(arr, mid+1, right) // Sorts the right half
	merge(arr, left, mid, right) // Merges both halves
}

// merge function combines two sorted lists into one single list
// Parameters:
//
//	arr common.Searchable[T]: The list containing two sorted sublists to be merged
//	left int: Start index for left sublist
//	mid int: Ending index of left sublist
//	right int: Ending index of right sublist
func merge[T cmp.Ordered](arr common.Searchable[T], left int, mid int, right int) {
	temp_arr := make([]T, arr.Size()) // Temporary slice to add sorted elemets

	leftP := left     // References to the start of left sublist
	rightP := mid + 1 // References to the start of right sublist

	for i := left; i <= right; i++ {
		leftVal, _ := arr.Get(leftP)
		rightVal, _ := arr.Get(rightP)
		if rightP > right { // All elements of the right sublist have been processed, the remaining elements of left sublist are added to the temporary slice
			temp_arr[i] = leftVal
			leftP++
		} else if leftP > mid { // All elements of the left sublist have been processed, the remaining elements of the right sublist are added to the temporary slice
			temp_arr[i] = rightVal
			rightP++
		} else if leftVal <= rightVal { // If the element in the left sublist is smaller or equal to the element in the right sublist, add the left element to temporary slice
			temp_arr[i] = leftVal
			leftP++
		} else { // If the element in the right sublist is smaller to the element in left sublist, add the right element to temporary slice
			temp_arr[i] = rightVal
			rightP++
		}
	}

	// Copy all elements from sorted slice to the original list
	for i := left; i <= right; i++ {
		arr.Set(i, temp_arr[i])
	}
}

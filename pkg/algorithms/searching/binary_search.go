package search

import (
	"cmp"

	"ds-algorithms/pkg/datastructures/common"
)

// BinarySearch function performs binary search on a Searchable compatible data structure
//
// Description:
// The algorithm finds the requested element by repeatedly splitting the array slice by half. Instead of linearly searching from the first
// element to the last element, we compare the element with the middle element. If the element is less than the comparing element, we narrow down the
// search field by half and only search on the left array. This process repeatedly occurs until we find the element we are looking for or
// if the element does not exist.
//
// Time Complexity: O(log N) - since the search space is halved in each iteration
// Space Complexity: O(1) - since we do the operations in-place
//
// Parameters:
//
//	data common.Searchable[T]: The ordered list that the algorithm performs the search on.
//	elem T: The element the algorithm looks for.
//
// Returns:
//
//	int: The index of the elem, or -1 if it does not exist in the slice.
//	bool: True if the algorithm found the elem, false otherwise.
//
// Example:
//
//	myList := common.NewArrayList([]int{1, 2, 3, 4, 5})
//
//	index, found := BinarySearch(myList, 2)
//	if found {
//		fmt.Printf("Element found at index: %d\n", index)
//	} else {
//		fmt.Println("Element not found.")
//	}
func BinarySearch[T cmp.Ordered](data common.Searchable[T], elem T) (int, bool) {
	leftIndex := 0
	rightIndex := data.Size() - 1

	for leftIndex <= rightIndex {
		middleIndex := leftIndex + (rightIndex-leftIndex)/2 // Calculates the midpoint of the current search field
		middleValue, err := data.Get(middleIndex)
		if err != nil {
			return -1, false
		}

		if elem == middleValue {
			return middleIndex, true // elem found at mid_index[i]
		} else if elem < middleValue { // elem is less than mid_index, search at left half of slice
			rightIndex = middleIndex - 1
		} else if elem > middleValue { // elem is greater than mid_index, search at right half of the slice
			leftIndex = middleIndex + 1
		}
	}
	return -1, false // elem do not exist in slice
}

// FindInsertionPoint performs a binary search on a Searchable compatible data structure to find the insertion point
// for an element.
//
// Algorithm: Binary Search (Lower Bound / Insertion Point)
// The algorithm determines the index where the elem should be inserted.
// It returns the index of the first element that is not less than elem.
// If elem is already present, this will be the index of its first occurrence.
// If elem is greater than all elements, the returned index will be the size of the data property.
//
// Time Complexity: O(log N) - The search space is halved in each iteration.
// Space Complexity: O(1) - since we do everything in-place
//
// Parameters:
//
//	data common.Searchable[T]: The ordered list that the algorithm performs the search on.
//	elem T: The element to find the insertion point for.
//
// Returns:
//
//	bool: true if elem was found at the returned index, false otherwise.
//	int: The index where elem should be inserted.
//
// Example:
//
//	myList := common.NewArrayList([]int{10, 20, 30, 40, 50})
//
//	// Element 30 is found at index 2
//	found, index := FindInsertionPoint(myList, 30)
//	if found {
//		fmt.Printf("Element 30 found at index: %d\n", index) // Output: Element 30 found at index: 2
//	}
//
//	// Element 25 is not found, but its insertion point is index 2
//	found, index = FindInsertionPoint(myList, 25)
//	if !found {
//		fmt.Printf("Element 25 not found, insertion point: %d\n", index) // Output: Element 25 not found, insertion point: 2
//	}
func FindInsertionPoint[T cmp.Ordered](data common.Searchable[T], elem T) (bool, int) {
	size := data.Size()
	if size == 0 {
		return false, 0 // if data is empty, the insertion point will be at index 0
	}

	leftIndex := 0
	rightIndex := size

	// Perform binary search to find the insertion point
	for leftIndex < rightIndex {
		mid := leftIndex + (rightIndex-leftIndex)/2
		currVal, _ := data.Get(mid)

		if currVal < elem {
			leftIndex = mid + 1
		} else {
			rightIndex = mid
		}
	}

	// Checks if the element at the current insertion point is equal to the value we want to insert
	if leftIndex < size {
		lowerBoundVal, _ := data.Get(leftIndex)
		if lowerBoundVal == elem {
			return true, leftIndex
		}
	}

	return false, leftIndex // Element were not found in current list
}

// FindUpperBound performs a binary search on a Searchable compatible data structure to find the upper bound
// for an element in a sorted data structure.
//
// Algorithm: Binary Search (Upper Bound)
// This algorithm determines the index of the first element that is
// strictly greater than elem. If elem is present, this will be the index
// immediately after its last occurrence. If all elements are less than or equal to
// elem, the returned index will be the size of the data property.
//
// Time Complexity: O(log N) - The search space is halved in each iteration.
// Space Complexity: O(1) - since we do everything in-place
//
// Parameters:
//
//	data common.Searchable[T]: The ordered list that the algorithm performs the search on.
//	elem T: The element to find the upper bound.
//
// Returns:
//
//	bool: true if elem was found anywhere in the data structure, false otherwise.
//	int: The index of the first element strictly greater than elem (the upper bound).
//
// Example:
//
//	myList := common.NewArrayList([]int{10, 20, 30, 30, 40, 50})
//	found, index := FindUpperBound(myList, 30)
//	fmt.Printf("Upper bound for 30: found=%t, index=%d\n", found, index) // Output: Upper bound for 30: found=true, index=4
func FindUpperBound[T cmp.Ordered](data common.Searchable[T], elem T) (bool, int) {
	size := data.Size()
	if size == 0 {
		return false, 0 // if data is empty, the upper bound will be at index 0
	}

	leftIndex := 0
	rightIndex := size

	// Perform binary search to find the upperbound
	for leftIndex < rightIndex {
		mid := leftIndex + (rightIndex-leftIndex)/2
		currVal, _ := data.Get(mid)

		if currVal <= elem {
			leftIndex = mid + 1
		} else {
			rightIndex = mid
		}
	}

	// Check if element is present in the list
	if leftIndex < size {
		lowerBoundVal, _ := data.Get(leftIndex)
		if lowerBoundVal == elem {
			return true, leftIndex
		}
	}

	return false, leftIndex
}

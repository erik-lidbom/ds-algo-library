package search

import (
	"cmp"
	"ds-algorithms/pkg/datastructures/common"
)

// BinarySearch function performs binary search on a Searchable compatible data structure
// Description:
// The algorithm finds the requested element by repeatedly splitting the array slice by half. Instead of linearly search from first //element to last element, we compare the elem with the middle element. If the elem is less than the comparing element, we narrow down the search field by half and only search on the left array. This process repeatedly occurs until we find the element we are looking for or if the element do not exist.
//
// Time Complexity: O(log N) - since the search space halved in each iteratiom
// Space Complexity: O(1) - since we do the operations in-place
//
// Parameters:
// data - the ordered list that the algorithm performs the search on. The list must satisfy the Searchable interface, so to use a normal slice you must cast the slice to be a searchable slice. The ArrayList struct already supports this interface.
// elem - the elem the algorithm looks for
//
// Returns:
// int - index of the elem, -1 if it do not exist in the slice
// bool - if the algorithm found the elem or not

func BinarySearch[T cmp.Ordered](data common.Searchable[T], elem T) (int, bool){

	leftIndex := 0
	rightIndex := data.Size() - 1

	for leftIndex <= rightIndex {
		middleIndex := leftIndex + (rightIndex - leftIndex) / 2 // Calculates the midpoint of the current search field
		middleValue, err := data.Get(middleIndex)
		if err != nil {
			return -1, false
		}

		if(elem == middleValue){
			return middleIndex, true // elem found at mid_index[i]
		} else if(elem < middleValue){ // elem is less than mid_index, search at left half of slice
			rightIndex = middleIndex - 1 
		} else if(elem > middleValue){ // elem is greater than mid_index, search at right half of the slice
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
// 	data - the ordered list that the algorithm performs the search on. 
// 	The list must satisfy the Searchable interface, so to use a normal slice you must cast the slice to be a searchable slice. 
//  The ArrayList struct already supports this interface.
// 	elem - the element to find the insertion point for
//
// Returns:
//   bool - true if elem was found at the returned index, false otherwise.
//   int  - The index where elem should be inserted

func FindInsertionPoint[T cmp.Ordered](data common.Searchable[T], elem T) (bool, int) {
	size := data.Size()
	if size == 0 {
		return false, 0 // if data is empty, the insertion point will be at index 0
	}
	
	leftIndex := 0
	rightIndex := size

	// Perform binary search to find the insertion point
	for leftIndex < rightIndex { 
		mid := leftIndex + (rightIndex - leftIndex) / 2
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
// 	data - the ordered list that the algorithm performs the search on. 
// 	The list must satisfy the Searchable interface, so to use a normal slice you must cast the slice to be a searchable slice. 
//  The ArrayList struct already supports this interface.
// 	elem - the element to find the upper bound
//
// Returns:
//  bool - true if elem was found anywhere in the data structure, false otherwise.
//  int  - The index of the first element strictly greater than elem (the upper bound).

// TODO --> WRITE TEST

func FindUpperBound[T cmp.Ordered](data common.Searchable[T], elem T) (bool, int) {
		size := data.Size()
	if size == 0 {
		return false, 0 // if data is empty, the upper bound will be at index 0
	}
	
	leftIndex := 0
	rightIndex := size
	
	// Perform binary search to find the upperbound
	for leftIndex < rightIndex { 
		mid := leftIndex + (rightIndex - leftIndex) / 2
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
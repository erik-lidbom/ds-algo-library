package search

import 
	(	
	"cmp"
	"ds-algorithms/pkg/datastructures/common"
	)

/*
Algorithm to search for elements in an unordered list
Complexity: O(n)
*/

// LinearSearch function performs linear search on Searchable compatible data structure
// Description:
// The algorithm finds the requested element by iterating from the first index until the last index of the list
// If the element is found, it returns the index
// Time Complexity: O(N) - since we iterate through N elements
// Space Complexity: O(1) - since we do the iterations in-place
//
// Parameters:
// data - the ordered list that the algorithm performs the search on. The list must satisfy the Searchable interface, so to use a normal slice you must cast the slice to be a searchable slice. The ArrayList struct already supports this interface.
// elem - the element the algorithm looks for
//
// Returns:
// int - index of the elem, -1 if it do not exist in the slice
// bool - if the algorithm found the elem or not

func LinearSearch[T cmp.Ordered] (data common.Searchable[T], element T) (int,bool) {
	
	for i := range data.Size(){
		value, err := data.Get(i)
		if err != nil {
			return -1, false
		}
		if(value == element){
			return i, true
		}
	}
	return 0, false
}
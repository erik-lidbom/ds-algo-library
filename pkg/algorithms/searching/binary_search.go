package search

import (
	"cmp"
	"ds-algorithms/pkg/datastructures/array"
)

/*
Searching algorithm of an ordered list
Complexity: O(log n)
*/

func BinarySearch(arr []int, element int) (int, bool){

	left_index := 0
	right_index := len(arr) - 1

	for left_index <= right_index {
		mid_index := left_index + (right_index - left_index) / 2

		if(element == arr[mid_index]){
			return mid_index, true
		} else if(element < arr[mid_index]){
			right_index = mid_index - 1
		} else if(element > arr[mid_index]){
			left_index = mid_index + 1
		}
	}
	return 0, false
}

// TODO --> ADD TESTS

func BinarySearchArrayList[T cmp.Ordered](arr *array.ArrayList[T], elem T) (bool, int) {
	leftIndex := 0
	rightIndex := arr.Size() - 1

	for leftIndex <= rightIndex {
		midIndex := leftIndex + (rightIndex - leftIndex) / 2
		currVal, _ := arr.Get(midIndex)
		if elem == currVal {
		return true, midIndex
		} else if elem < currVal {
			rightIndex = midIndex - 1
		} else { 
			leftIndex = midIndex + 1
		}
	}
	return false, -1
}

func FindInsertionPoint[T cmp.Ordered](arr *array.ArrayList[T], elem T) (bool, int) {

	size := arr.Size() - 1
    leftIndex := 0
    rightIndex := size

    for leftIndex < rightIndex { 
        mid := leftIndex + (rightIndex - leftIndex) / 2
        currVal, _ := arr.Get(mid)

        if currVal < elem {
            leftIndex = mid + 1
        } else {
            rightIndex = mid   
        }
    }
  
    if leftIndex < size {
        lowerBoundVal, _ := arr.Get(leftIndex)
        if lowerBoundVal == elem {
            return true, leftIndex
        }
    }

    return false, leftIndex
}

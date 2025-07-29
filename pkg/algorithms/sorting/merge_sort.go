package sorting

import (
	"cmp"
	"ds-algorithms/pkg/datastructures/common"
)

func MergeSort[T cmp.Ordered] (arr common.Searchable[T], left int, right int) {
	if left >= right {
		return
	}

	mid := (left + right) / 2
	MergeSort(arr, left, mid)
	MergeSort(arr, mid + 1, right)
	merge(arr, left, mid, right)
}

func merge[T cmp.Ordered] (arr common.Searchable[T], left int, mid int, right int) {
	temp_arr := make([]T, arr.Size())

	leftP := left
	rightP := mid + 1

	for i := left; i <= right; i++{
		leftVal, _ := arr.Get(leftP)
		rightVal, _ := arr.Get(rightP)
		if(rightP > right) {
			temp_arr[i] = leftVal
			leftP++
		} else if(leftP > mid) {
			temp_arr[i] = rightVal
			rightP++
		} else if(leftVal <= rightVal){
			temp_arr[i] = leftVal
			leftP++
		} else {
			temp_arr[i] = rightVal
			rightP++
		}
	}

	for i := left; i <= right; i++ {
		arr.Set(i, temp_arr[i])
	}
}

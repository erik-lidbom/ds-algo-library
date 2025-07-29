package sorting

import (
	"cmp"

	"ds-algorithms/pkg/datastructures/common"
)

func QuickSort[T cmp.Ordered](arr common.Searchable[T], left int, right int) {
	if left >= right {
		return
	}

	if arr.Size() <= 100 {
		InsertionSort(arr)
		return
	}

	pivot := findPivot(left, right)
	pivot = partition(arr, left, right, pivot)
	QuickSort(arr, left, pivot)
	QuickSort(arr, pivot+1, right)
}

func findPivot(left int, right int) int {
	return (left + right) / 2
}

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

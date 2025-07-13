package sorting



func QuickSort(arr []int, left int, right int) {
	if left >= right {
		return
	}

	if len(arr) <= 100 {
		InsertionSort(arr)
		return
	}

	pivot := findPivot(left, right)
	pivot = partition(arr, left, right, pivot)
	QuickSort(arr, left, pivot)
	QuickSort(arr, pivot + 1, right)
}

func findPivot(left int, right int) int {
	return (left + right) / 2
}

func partition(arr []int, left int, right int, pivot int) int {
	pivotVal := arr[pivot]

	for {
		for left <= right && arr[left] < pivotVal {
			left++
		}
		for left <= right && arr[right] > pivotVal {
			right--
		} 

		if left >= right {
			return right
		}

		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
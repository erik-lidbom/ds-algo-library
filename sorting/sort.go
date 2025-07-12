package sort

import "fmt"

func BubbleSort (arr[]int) {
	size := len(arr)

	for i := 0; i < size - 1; i++ {
		for j := 0; j < size - i - 1; j++{
			if(arr[j] > arr[j + 1]){
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println("Newly sorted array")
	
	for i:=0; i < size - 1; i++{
		fmt.Printf("%d, ", arr[i])
	}
}
func SelectionSort (arr[]int) {
	size := len(arr)

	for i := 0; i < size - 1; i++ {
		min_index := i
		
		for j := i + 1; j < size; j++ {
			if(arr[j] < arr[min_index]){
				min_index = j
			}
		}

		arr[i], arr[min_index] = arr[min_index], arr[i]
	}

	fmt.Println("Newly sorted array")
	
	for i:=0; i < size - 1; i++{
		fmt.Printf("%d, ", arr[i])
	}
}

func InsertionSort (arr[]int) {

	size := len(arr)
	for i := 0; i < size - 1; i++ {
		j := i

		for j > 0 && arr[j] < arr[j - 1] {
			arr[j], arr[j - 1] = arr[j - 1], arr[j]
			j--
		}
	} 

	fmt.Println("Newly sorted array")
	
	for i := 0; i < size - 1; i++{
		fmt.Printf("%d, ", arr[i])
	}
}

func MergeSort (arr []int, left int, right int) {
	if left >= right {
		return
	}

	mid := (left + right) / 2
	MergeSort(arr, left, mid)
	MergeSort(arr, mid + 1, right)
	merge(arr, left, mid, right)
}

func merge (arr []int, left int, mid int, right int) {
	temp_arr := make([]int, len(arr))

	left_p := left
	right_p := mid + 1

	for i := left; i <= right; i++{
		if(right_p > right) {
			temp_arr[i] = arr[left_p]
			left_p++
		} else if(left_p > mid) {
			temp_arr[i] = arr[right_p]
			right_p++
		} else if(arr[left_p] <= arr[right_p]){
			temp_arr[i] = arr[left_p]
			left_p++
		} else {
			temp_arr[i] = arr[right_p]
			right_p++
		}
	}

	for i := left; i <= right; i++ {
		arr[i] = temp_arr[i]
	}
}


func QuickSort(arr []int, left int, right int) {
	if left >= right {
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
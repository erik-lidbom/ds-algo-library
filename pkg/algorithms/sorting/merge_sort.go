package sorting

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

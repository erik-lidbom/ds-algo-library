package sorting

func InsertionSort (arr[]int) {

	size := len(arr)

	for i := 1; i < size; i++ {
		j := i

		for j > 0 && arr[j] < arr[j - 1] {
			arr[j], arr[j - 1] = arr[j - 1], arr[j]
			j--
		}
	} 
}